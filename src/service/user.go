package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Eli15x/ZCOM/src/model"
	"github.com/Eli15x/ZCOM/src/client"
	"github.com/Eli15x/ZCOM/src/repository"
	"github.com/Eli15x/ZCOM/src/utils"
	"github.com/fatih/structs"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	instanceUser CommandUser
	onceUser    sync.Once
)

type CommandUser interface {
	ValidateUser(ctx context.Context, email string, password string) error
	CreateUser(ctx context.Context, name string, email string, password string) error
	GetInformationUser(ctx context.Context, id string) ([]bson.M, error)
	DeleteUser(ctx context.Context, id string) ([]bson.M, error)
}

type user struct{}

func GetInstanceUser() CommandUser {
	onceUser.Do(func() {
		instanceUser = &user{}
	})
	return instanceUser
}


func (u *user) ValidateUser(ctx context.Context, email string, password string) error {
	var user model.user
	
	emailValidate := map[string]interface{}{"email": email}

	result, err := repository.Find(ctx, "user", emailValidate, &user)
	if err != nil {
		return errors.New("Validate user: problem to get information into MongoDB")
	}

	passwordEncrypt := utils.Encrypt(password)
	if passwordEncrypt != result.PassWord{
		return errors.New("Password user: wrong password")
	}
	
	return nil
}


func (u *user) CreateUser(ctx context.Context, user models.userRequest) error {

	var userId = utils.CreateCodeId()
	user := &models.user{
		UserId:   userId,
		Name:     user.name,
		Email:    user.email,
		PassWord: utils.Encrypt(user.password),
		idAcess:  user.idAcess,
	}

	userInsert := structs.Map(user)

	_, err := client.GetInstance().Insert(ctx, "user", userInsert)
	if err != nil {
		return errors.New("Create user: problem to insert into MongoDB")
	}

	err = CreateFriendTable(ctx, userId)
	if err != nil {
		return err
	}

	return nil
}


func (u *user) Edituser(ctx context.Context, user models.user) error {

	user := &models.user{
		UserId:   user.userId,
		Name:     user.name,
		Email:    user.email,
		PassWord: user.password,
		idAcess:  user.idAcess,
	}

	userUpdate:= structs.Map(user)

	_, err := client.GetInstance().UpdateOne(ctx, "user", userInsert)
	if err != nil {
		return errors.New("Edit User: problem to update into MongoDB")
		//adicionar aqui retorno de erro (500)
	}

	return nil
}



func (u *user) GetInformationUser(ctx context.Context, id string) ([]bson.M, error) {
	var user models.user

	userId := map[string]interface{}{"UserId": id}

	result, err := client.Find(ctx, "user", userId, &user)
	if err != nil {
		return nil, errors.New("Add Information user: problem to Find Id into MongoDB")
	}

	return result, nil
}

func (u *user) DeleteUser(ctx context.Context, userId string) error {

	userId := map[string]interface{}{"UserId": id}

	_, err := client.GetInstance().Remove(ctx, "user", userId)
	if err != nil {
		return errors.New("Delete User: problem to delete into MongoDB")
	}


	return nil
}
