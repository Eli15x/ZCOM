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

func GetUnidade(c *gin.Context) {

	var unidade model.Unidade
	err := json.NewDecoder(c.Request.Body).Decode(&unidade)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if unidade.Sigla == "" {
		c.String(http.StatusBadRequest, "Get Product Error: Sigla not informed")
		return
	}


	result, err := service.GetInstanceUnidade().GetUnidade(context.Background(), unidade.Sigla)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}

func GetUnidades(c *gin.Context) {

	result, err := service.GetInstanceUnidade().GetUnidades(context.Background())
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, result) //return list
}