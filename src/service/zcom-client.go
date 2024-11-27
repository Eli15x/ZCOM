package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"sync"
	//kafka "ZCOM/src/client/kafka"
)

var (
	instanceServiceZcomClient ServiceZcomClient
	onceServiceZcomClient     sync.Once
)

type ServiceZcomClient interface {
	Login(ctx context.Context) (string, error)
}

type zcomClient struct{}

func GetInstanceZcomClient() ServiceZcomClient {
	onceServiceZcomClient.Do(func() {
		instanceServiceZcomClient = &zcomClient{}
	})
	return instanceServiceZcomClient
}

// Função para login e obtenção do token JWT
func (z *zcomClient) Login(ctx context.Context) (string, error) {

	credentials := map[string]string{
		"Email":    os.Getenv("LOGINZCOMCLIENT_EMAIL"),
		"Password": os.Getenv("LOGINZCOMCLIENT_PASSWORD"),
	}
	credentialsJSON, _ := json.Marshal(credentials)

	resp, err := http.Post(os.Getenv("LoginUrl"), "application/json", bytes.NewBuffer(credentialsJSON))
	if err != nil {
		return "", fmt.Errorf("erro ao fazer requisição de login: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("erro de login: status %d - %s", resp.StatusCode, body)
	}

	var loginResponse struct {
		UserID string `bson:"user_id,omitempty" json:"user_id,omitempty"`
		JWT    string `bson:"JWT,omitempty" json:"JWT,omitempty"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
		return "", fmt.Errorf("erro ao decodificar resposta de login: %v", err)
	}

	return loginResponse.JWT, nil
}
