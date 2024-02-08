package handlers

import (
	"context"
	"encoding/json"

	"net/http"
	"ZCOM/src/service"
	"ZCOM/src/model"
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

	if product.CODIGO_CEST == "" {
		c.String(http.StatusBadRequest, "Create Product Error: barCode not informed")
		return
	}

	if product.NAME == "" {
		c.String(400, "Create Product Error: Name not informed")
		return
	}

	if product.CODIGO_NCM == "" {
		c.String(400, "Create Product Error: NCM not informed")
		return
	}

	if product.CFOP == "" {
		c.String(400, "Create Product Error: CFOP not informed")
		return
	}

	if product.GTIN == "" {
		c.String(400, "Create Product Error: GTIN not informed")
		return
	}

	if product.IAT == "" {
		c.String(400, "Create Product Error: IAT not informed")
		return
	}

	if product.ID_PRODUTO_GRUPO == 0 {
		c.String(400, "Create Product Error: ID_PRODUTO_GRUPO not informed")
		return
	}

	if product.ID_PRODUTO_MARCA == 0 {
		c.String(400, "Create Product Error: ID_PRODUTO_MARCA not informed")
		return
	}

	if product.ID_PRODUTO_UNIDADE == 0 {
		c.String(400, "Create Product Error: ID_PRODUTO_UNIDADE not informed")
		return
	}

	if product.ID_TRIBUT_GRUPO_TRIBUTARIO == 0 {
		c.String(400, "Create Product Error: ID_TRIBUT_GRUPO_TRIBUTARIO not informed")
		return
	}

	if product.IPPT == "" {
		c.String(400, "Create Product Error: IPPT not informed")
		return
	}

	if product.PESO == 0 {
		c.String(400, "Create Product Error: PESO not informed")
		return
	}

	if product.QUANTIDADE_EMBALAGEM == 0 {
		c.String(400, "Create Product Error: QUANTIDADE_EMBALAGEM not informed")
		return
	}

	if product.QUANTIDADE_ESTOQUE == 0 {
		c.String(400, "Create Product Error: QUANTIDADE_ESTOQUE not informed")
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


	/*if product.IndRegra == "" {
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
	}*/


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

	
	if product.CODIGO_CEST == "" {
		c.String(http.StatusBadRequest, "Edit Product Error: barCode not informed")
		return
	}

	if product.NAME == "" {
		c.String(400, "Edit Product Error: Name not informed")
		return
	}

	if product.CODIGO_NCM == "" {
		c.String(400, "Edit Product Error: NCM not informed")
		return
	}

	if product.CFOP == "" {
		c.String(400, "Edit Product Error: CFOP not informed")
		return
	}

	if product.GTIN == "" {
		c.String(400, "Edit Product Error: GTIN not informed")
		return
	}

	if product.IAT == "" {
		c.String(400, "Edit Product Error: IAT not informed")
		return
	}

	if product.ID_PRODUTO_GRUPO == 0 {
		c.String(400, "Edit Product Error: ID_PRODUTO_GRUPO not informed")
		return
	}

	if product.ID_PRODUTO_MARCA == 0 {
		c.String(400, "Edit Product Error: ID_PRODUTO_MARCA not informed")
		return
	}

	if product.ID_PRODUTO_UNIDADE == 0 {
		c.String(400, "Edit Product Error: ID_PRODUTO_UNIDADE not informed")
		return
	}

	if product.ID_TRIBUT_GRUPO_TRIBUTARIO == 0 {
		c.String(400, "Edit Product Error: ID_TRIBUT_GRUPO_TRIBUTARIO not informed")
		return
	}

	if product.IPPT == "" {
		c.String(400, "Edit Product Error: IPPT not informed")
		return
	}

	if product.PESO == 0 {
		c.String(400, "Edit Product Error: PESO not informed")
		return
	}

	if product.QUANTIDADE_EMBALAGEM == 0 {
		c.String(400, "Edit Product Error: QUANTIDADE_EMBALAGEM not informed")
		return
	}

	if product.QUANTIDADE_ESTOQUE == 0 {
		c.String(400, "Edit Product Error: QUANTIDADE_ESTOQUE not informed")
		return
	}


	/*if product.Desconto == 0 {
		c.String(400, "Edit Product Error: Desconto not informed")
		return
	}


	if product.OutrosDesconto == 0 {
		c.String(400, "Edit Product Error: OutrosDesconto not informed")
		return
	}*/


	/*if product.IndRegra == "" {
		c.String(400, "Edit Product Error: IndRegra not informed")
		return
	}


	if product.UCom == 0 {
		c.String(400, "Edit Product Error: UCom not informed")
		return
	}

	if product.QCom == 0 {
		c.String(400, "Edit Product Error: QCom not informed")
		return
	}

	if product.VUnCom == 0 {
		c.String(400, "Edit Product Error: VUnCom not informed")
		return
	}*/

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

	if product.CODIGO_CEST == "" {
		c.String(http.StatusBadRequest, "Delete Product Error: CODIGO_CEST not informed")
		return
	}

	err = service.GetInstanceProduct().DeleteProduct(context.Background(), product.CODIGO_CEST)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func GetCodeCest(c *gin.Context) {

	var product model.Product
	err := json.NewDecoder(c.Request.Body).Decode(&product)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if product.CODIGO_CEST == "" {
		c.String(http.StatusBadRequest, "Get Product Error: CODIGO_CEST not informed")
		return
	}

	result, err := service.GetInstanceProduct().GetProduct(context.Background(), product.CODIGO_CEST)
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

	if product.NAME == "" {
		c.String(http.StatusBadRequest, "Get Product Error: NAME not informed")
		return
	}


	result, err := service.GetInstanceProduct().GetProductByName(context.Background(), product.NAME)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

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