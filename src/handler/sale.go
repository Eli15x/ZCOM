package handlers

import (
	"context"
	"encoding/json"

	"ZCOM/src/model"
	"ZCOM/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSaleXml(c *gin.Context) {

	var saleXML model.SaleXML
	err := json.NewDecoder(c.Request.Body).Decode(&saleXML)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if saleXML.Path == "" {
		c.String(400, "Create SaleXml Error: Path not informed")
		return
	}

	if saleXML.Name == "" {
		c.String(400, "Create SaleXml Error: Name not informed")
		return
	}

	err = service.GetInstanceSale().SaveSaleXMLKafka(context.Background(), saleXML)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func SendSale(c *gin.Context) {

	err := service.GetInstanceSale().SendSaleZipToHostgator(context.Background())
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}
