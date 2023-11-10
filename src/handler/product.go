package handlers

import (
	"context"
	"encoding/json"

	"net/http"
	"github.com/Eli15x/ZCOM/src/service"
	"github.com/Eli15x/ZCOM/src/model"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func CreateProduct(c *gin.Context) {

	var product model.Product
	err := json.NewDecoder(c.Request.Body).Decode(&product)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if product.BarCodeNumber == "" {
		c.String(http.StatusBadRequest, "Create Product Error: barCode not informed")
		return
	}

	if product.Name == "" {
		c.String(400, "Create Product Error: Name not informed")
		return
	}

	if product.NCM == "" {
		c.String(400, "Create Product Error: NCM not informed")
		return
	}

	if product.CFOP == "" {
		c.String(400, "Create Product Error: CFOP not informed")
		return
	}

	/*if product.Desconto == 0 {
		c.String(400, "Create Product Error: Desconto not informed")
		return
	}


	if product.OutrosDesconto == 0 {
		c.String(400, "Create Product Error: OutrosDesconto not informed")
		return
	}*/


	if product.IndRegra == "" {
		c.String(400, "Create Product Error: IndRegra not informed")
		return
	}


	if product.UCom == 0 {
		c.String(400, "Create Product Error: UCom not informed")
		return
	}

	if product.QCom == 0 {
		c.String(400, "Create Product Error: QCom not informed")
		return
	}

	if product.VUnCom == 0 {
		c.String(400, "Create Product Error: VUnCom not informed")
		return
	}


	err = service.GetInstanceProduct().CreateProduct(context.Background(), product)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "" )
}

func EditProduct(c *gin.Context) {

	var product model.Product
	err := json.NewDecoder(c.Request.Body).Decode(&product)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if product.BarCodeNumber == "" {
		c.String(http.StatusBadRequest, "Create Product Error: barCode not informed")
		return
	}

	if product.Name == "" {
		c.String(400, "Create Product Error: Name not informed")
		return
	}

	if product.NCM == "" {
		c.String(400, "Create Product Error: NCM not informed")
		return
	}

	if product.CFOP == "" {
		c.String(400, "Create Product Error: CFOP not informed")
		return
	}

	/*if product.Desconto == 0 {
		c.String(400, "Create Product Error: Desconto not informed")
		return
	}


	if product.OutrosDesconto == 0 {
		c.String(400, "Create Product Error: OutrosDesconto not informed")
		return
	}*/


	if product.IndRegra == "" {
		c.String(400, "Create Product Error: IndRegra not informed")
		return
	}


	if product.UCom == 0 {
		c.String(400, "Create Product Error: UCom not informed")
		return
	}

	if product.QCom == 0 {
		c.String(400, "Create Product Error: QCom not informed")
		return
	}

	if product.VUnCom == 0 {
		c.String(400, "Create Product Error: VUnCom not informed")
		return
	}

	err = service.GetInstanceProduct().EditProduct(context.Background(), product)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func DeleteProduct(c *gin.Context) {

	var product model.Product
	err := json.NewDecoder(c.Request.Body).Decode(&product)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if product.BarCodeNumber == "" {
		c.String(http.StatusBadRequest, "Create Product Error: barCode not informed")
		return
	}

	err = service.GetInstanceProduct().DeleteProduct(context.Background(), product.BarCodeNumber)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func GetByBarCode(c *gin.Context) {

	var product model.Product
	err := json.NewDecoder(c.Request.Body).Decode(&product)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if product.BarCodeNumber == "" {
		c.String(http.StatusBadRequest, "Create Product Error: barCode not informed")
		return
	}

	result, err := service.GetInstanceProduct().GetProduct(context.Background(), product.BarCodeNumber)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}

func GetProductByName(c *gin.Context) {

	var product model.Product
	err := json.NewDecoder(c.Request.Body).Decode(&product)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if product.Name == "" {
		c.String(http.StatusBadRequest, "Create Product Error: barCode not informed")
		return
	}


	result, err := service.GetInstanceProduct().GetProductByName(context.Background(), product.Name)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", "ola", "")

	c.JSON(http.StatusOK, result)
}

func GetProducts(c *gin.Context) {

	result, err := service.GetInstanceProduct().GetProducts(context.Background())
	if err != nil {
		c.String(400, err.Error())
		return
	}


	c.JSON(http.StatusOK, result) //return list
}

func SaveProduct(c *gin.Context){

	err := service.GetInstanceProduct().SaveProduct(context.Background())
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok") //return list
}