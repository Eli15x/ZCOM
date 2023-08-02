package handlers

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func ValidateUser(c *gin.Context) {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	password := json_map["password"].(string)
	email := json_map["email"].(string)

	if email == "" {
		c.String(http.StatusBadRequest, "Validate User Error: email not find")
		return
	}

	if password == "" {
		c.String(http.StatusBadRequest, "Create User Error: password not find")
		return
	}

	/*resultUser, err := service.GetInstanceUser().ValidateUser(context.Background(), email, password)
	if err != nil {
		c.String(400, err.Error())
		return
	}*/

	c.JSON(http.StatusOK, &resultUser)
}

func CreateUser(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	name := json_map["name"].(string)
	email := json_map["email"].(string) 
	password := json_map["password"].(string)
	acess := json_map["acess"].(int)
	userId := json_map["userId"].(string)

	if name == "" {
		c.String(http.StatusBadRequest, "Create User Error: name not find")
		return
	}

	if email == "" {
		c.String(400, "Create User Error: email not find")
		return
	}

	if password == "" {
		c.String(400, "Create User Error: password not find")
		return
	}

	/*userId, err := service.GetInstanceUser().CreateNewUser(context.Background(), name, email, password, telefone)
	if err != nil {
		c.String(400, err.Error())
		return
	}*/

	c.String(http.StatusOK, userId)
}

func EditUser(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	name := json_map["name"].(string)
	email := json_map["email"].(string) 
	password := json_map["password"].(string)
	acess := json_map["acess"].(int)
	userId := json_map["userId"].(string)

	if name == "" {
		c.String(http.StatusBadRequest, "Edit User Error: name not find")
		return
	}

	if email == "" {
		c.String(400, "Edit User Error: email not find")
		return
	}

	if password == "" {
		c.String(400, "Edit User Error: password not find")
		return
	}

	if telefone == "" {
		c.String(400, "Edit User Error: password not find")
		return
	}

	if userId == "" {
		c.String(400, "Edit User Error: password not find")
		return
	}

	/*err = service.GetInstanceUser().EditUser(context.Background(), userId, name, email, password, telefone)
	if err != nil {
		c.String(400, err.Error())
		return
	}*/

	c.String(http.StatusOK, "")
}

func DeleteUser(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	userId := json_map["userId"].(string)

	if userId == "" {
		c.String(http.StatusBadRequest, "Delete User Error: userId not find")
		return
	}

	/*err = service.GetInstanceUser().DeleteUser(context.Background(), userId)
	if err != nil {
		c.String(400, err.Error())
		return
	}*/

	c.String(http.StatusOK, "")
}

func GetInformationByUserId(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	userId := json_map["userId"].(string) //está dando erro quando tenta pegar o "email" e ele não existe.

	/*result, err := service.GetInstanceUser().GetInformationUser(context.Background(), userId)
	if err != nil {
		c.String(400, err.Error())
		return
	}*/

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}


func GetUserByName(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	userId := json_map["userId"].(string) //está dando erro quando tenta pegar o "email" e ele não existe.

	/*result, err := service.GetInstanceUser().GetInformationUser(context.Background(), userId)
	if err != nil {
		c.String(400, err.Error())
		return
	}*/

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result) //return list
}

func GetUserByFunction(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	userId := json_map["userId"].(string) //está dando erro quando tenta pegar o "email" e ele não existe.

	/*result, err := service.GetInstanceUser().GetInformationUser(context.Background(), userId)
	if err != nil {
		c.String(400, err.Error())
		return
	}*/

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result) //return list
}

func GetUserByAcess(c *gin.Context) {

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	userId := json_map["userId"].(string) //está dando erro quando tenta pegar o "email" e ele não existe.

	/*result, err := service.GetInstanceUser().GetInformationUser(context.Background(), userId)
	if err != nil {
		c.String(400, err.Error())
		return
	}*/

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result) //return list
}