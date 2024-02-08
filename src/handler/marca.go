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


func GetMarca(c *gin.Context) {

	var marca model.Marca
	err := json.NewDecoder(c.Request.Body).Decode(&marca)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if marca.Nome == "" {
		c.String(http.StatusBadRequest, "Get marca Error: Nome not informed")
		return
	}

	result, err := service.GetInstanceMarca().GetMarca(context.Background(), marca.Nome)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}

func GetMarcas(c *gin.Context) {

	result, err := service.GetInstanceMarca().GetMarcas(context.Background())
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, result) //return list
}