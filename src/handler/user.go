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

func ValidateUser(c *gin.Context) {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	password := json_map["passWord"].(string)
	email := json_map["email"].(string)

	if email == "" {
		c.String(http.StatusBadRequest, "Validate User Error: email not find")
		return
	}

	if password == "" {
		c.String(http.StatusBadRequest, "Create User Error: password not find")
		return
	}

	userId, err := service.GetInstanceUser().ValidateUser(context.Background(), email, password)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, &userId)
}

func CreateUser(c *gin.Context) {

	var user model.UserRequest
	err := json.NewDecoder(c.Request.Body).Decode(&user)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if user.Name == "" {
		c.String(http.StatusBadRequest, "Create User Error: name not find")
		return
	}

	if user.Email == "" {
		c.String(400, "Create User Error: email not find")
		return
	}

	if user.PassWord == "" {
		c.String(400, "Create User Error: password not find")
		return
	}

	if user.IdAcess == 0 {
		c.String(400, "Create User Error: idAcess not find")
		return
	}

	err = service.GetInstanceUser().CreateUserKafka(context.Background(), user)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "" )
}

func EditUser(c *gin.Context) {

	var user model.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if user.Name == "" {
		c.String(http.StatusBadRequest, "Edit User Error: name not find")
		return
	}

	if user.Email == "" {
		c.String(400, "Edit User Error: email not find")
		return
	}

	if user.PassWord == "" {
		c.String(400, "Edit User Error: password not find")
		return
	}


	if user.UserId == "" {
		c.String(400, "Edit User Error: password not find")
		return
	}

	err = service.GetInstanceUser().EditUserKafka(context.Background(), user)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func DeleteUser(c *gin.Context) {

	var user model.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if user.UserId == "" {
		c.String(http.StatusBadRequest, "Delete User Error: userId not find")
		return
	}

	err = service.GetInstanceUser().DeleteUserKafka(context.Background(), user.UserId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func GetInformationByUserId(c *gin.Context) {

	var user model.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if user.UserId == "" {
		c.String(http.StatusBadRequest, "Delete User Error: userId not find")
		return
	}

	result, err := service.GetInstanceUser().GetUser(context.Background(), user.UserId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}

func GetUserByName(c *gin.Context) {

	var user model.UserRequest
	err := json.NewDecoder(c.Request.Body).Decode(&user)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if user.Name == "" {
		c.String(http.StatusBadRequest, "Get User By Name: name not find")
		return
	}

	result, err := service.GetInstanceUser().GetUserByName(context.Background(), user.Name)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", "ola", "")

	c.JSON(http.StatusOK, result)
}


func GetUsersByAcess(c *gin.Context) {

	var user model.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if user.IdAcess == 0 {
		c.String(http.StatusBadRequest, "Get User By Acess: IdAcess not find")
		return
	}

	result, err := service.GetInstanceUser().GetUsersByAcess(context.Background(), user.IdAcess)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", "", "")

	c.JSON(http.StatusOK,result) 
}

func GetUsers(c *gin.Context) {

	result, err := service.GetInstanceUser().GetUsers(context.Background())
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK,result) 
}

func SaveUser(c *gin.Context){

	err := service.GetInstanceUser().SaveUser(context.Background())
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok") //return list
}