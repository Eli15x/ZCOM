package handlers

import (
	"context"
	"encoding/json"

	"ZCOM/src/model"
	"ZCOM/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func GetGrupo(c *gin.Context) {

	var grupo model.Grupo
	err := json.NewDecoder(c.Request.Body).Decode(&grupo)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if grupo.Grupo == "" {
		c.String(http.StatusBadRequest, "Get Grupo Error: Grupo not informed")
		return
	}

	result, err := service.GetInstanceGrupo().GetGrupo(context.Background(), grupo.Grupo)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", "ola", "")

	c.JSON(http.StatusOK, result)
}

func GetGrupos(c *gin.Context) {

	result, err := service.GetInstanceGrupo().GetGrupos(context.Background())
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, result) //return list
}
