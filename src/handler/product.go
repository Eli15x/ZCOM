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

func ValidateProduct(c *gin.Context) {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request.Body).Decode(&json_map)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	barCodeNumber := json_map["barCodeNumber"].(string)
	name := json_map["name"].(string)

	if barCodeNumber == "" {
		c.String(http.StatusBadRequest, "Validate Product Error: productId not find")
		return
	}

	if name == "" {
		c.String(http.StatusBadRequest, "Validate Product Error: Name not find")
		return
	}

	userId, err := service.GetInstanceUser().ValidateUser(context.Background(), barCodeNumber, name)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(http.StatusOK, &userId)
}

func CreateProduct(c *gin.Context) {

	var user model.UserRequest
	err := json.NewDecoder(c.Request.Body).Decode(&user)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if user.Name == "" {
		c.String(http.StatusBadRequest, "Create Product Error: name not find")
		return
	}

	if user.Email == "" {
		c.String(400, "Create Product Error: email not find")
		return
	}

	if user.PassWord == "" {
		c.String(400, "Create Product Error: password not find")
		return
	}

	if user.IdAcess == 0 {
		c.String(400, "Create Product Error: idAcess not find")
		return
	}

	err = service.GetInstanceUser().CreateUser(context.Background(), user)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "" )
}

func EditProduct(c *gin.Context) {

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

	err = service.GetInstanceUser().EditUser(context.Background(), user)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func DeletProduct(c *gin.Context) {

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

	err = service.GetInstanceUser().DeleteUser(context.Background(), user.UserId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}

func GetByBarCode(c *gin.Context) {

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

	result, err := service.GetInstanceUser().GetInformationUser(context.Background(), user.UserId)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	log.Infof("[GetInformation] Object : %s \n", result, "")

	c.JSON(http.StatusOK, result)
}

func GetProductByName(c *gin.Context) {

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

func GetProducts(c *gin.Context) {

	c.JSON(http.StatusOK,"") //return list
}