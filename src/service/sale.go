package service

import (
	kafka "ZCOM/src/client/kafka"
	"ZCOM/src/utils"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"ZCOM/src/model"
	"encoding/json"
	"os"
)

var (
	instanceServiceSale ServiceSale
	onceServiceSale     sync.Once
)

type ServiceSale interface {
	SaveSaleXML(ctx context.Context, saleXML model.SaleXML) error
	SaveSaleXMLKafka(ctx context.Context, saleXML model.SaleXML) error
	SendSaleZipToHostgator(ctx context.Context) error
}

type sale struct{}

func GetInstanceSale() ServiceSale {
	onceServiceSale.Do(func() {
		instanceServiceSale = &sale{}
	})
	return instanceServiceSale
}

func (s *sale) SaveSaleXMLKafka(ctx context.Context, saleXML model.SaleXML) error {
	saleXMLExist, _ := s.GetSaleXML(ctx, saleXML.Name, saleXML.Path)
	if saleXMLExist.Path != "" {
		return errors.New("Save Sale Kafka: this Sale XML exists")
	}

	saleXMLJson, err := json.Marshal(saleXML)

	err = kafka.GetInstanceKafka().SendMessage(saleXMLJson, "createSaleXML")
	if err != nil {
		return err
	}

	return nil
}

func (s *sale) GetSaleXML(ctx context.Context, name string, path string) (model.SaleXML, error) {
	var saleXML model.SaleXML
	namefile := name + ".xml"
	data, err := os.ReadFile(path + "/" + namefile)
	if err != nil {
		return saleXML, err
	}
	json.Unmarshal([]byte(data), &saleXML)

	return saleXML, nil
}

func (s *sale) SaveSaleXML(ctx context.Context, saleXML model.SaleXML) error {

	if _, err := os.Stat(saleXML.Path); os.IsNotExist(err) {
		err := os.Mkdir(saleXML.Path, 0755) //create a directory and give it required permissions
		if err != nil {
			return err
		}
	}

	SaleJson, err := json.Marshal(saleXML)
	if err != nil {
		return err
	}

	namefile := saleXML.Name + ".xml"

	if err = os.WriteFile(saleXML.Path+"/"+namefile, SaleJson, 0666); err != nil {
		return err
	}

	return nil
}

func (s *sale) SendSaleZipToHostgator(ctx context.Context) error {

	//no servidor de la quando receber esses sales eu perciso ler cada um e pegar as informacoes para atualizar no sistema.
	zipData, err := utils.CreateZipFromDirectory(os.Getenv("UrlSaleXml"))

	jwtToken, err := GetInstanceZcomClient().Login(ctx)

	req, err := http.NewRequest("POST", os.Getenv("SendSaleUrl"), bytes.NewBuffer(zipData))
	if err != nil {
		return fmt.Errorf("erro ao criar requisição: %v", err)
	}
	req.Header.Set("Content-Type", "application/zip")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erro ao enviar requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("erro no servidor: status %d - %s", resp.StatusCode, body)
	}

	return nil

}
