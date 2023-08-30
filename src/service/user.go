package service

import (
	"context"
	"errors"
	"sync"

	"github.com/Eli15x/ZCOM/src/model"
	"github.com/Eli15x/ZCOM/src/client"
	"github.com/Eli15x/ZCOM/src/repository"
	"github.com/Eli15x/ZCOM/src/utils"
	"github.com/fatih/structs"
)

var (
	instanceUser CommandUser
	onceUser    sync.Once
)

type CommandUser interface {
	ValidateUser(ctx context.Context, email string, password string) (string, error)
	CreateUser(ctx context.Context, user model.UserRequest) error
	EditUser(ctx context.Context, user model.User) error
	GetInformationUser(ctx context.Context, id string) (model.User, error)
	GetUserByName(ctx context.Context, name string) (model.User, error)
	GetUsersByAcess(ctx context.Context, idAcess int) ([]model.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type user struct{}

func GetInstanceUser() CommandUser {
	onceUser.Do(func() {
		instanceUser = &user{}
	})
	return instanceUser
}

func (u *user) ValidateUser(ctx context.Context, email string, password string) (string, error) {
	var user model.User
	
	emailValidate := map[string]interface{}{"email": email}

	result, err := repository.Find(ctx, "user", emailValidate, &user)
	if err != nil {
		return "",errors.New("Validate user: problem to get information into MongoDB")
	}

	passwordEncrypt := utils.Encrypt(password)
	if passwordEncrypt != result.PassWord{
		return "", errors.New("Password user: wrong password")
	}
	
	return user.UserId,nil
}

func (u *user) CreateUser(ctx context.Context, user model.UserRequest) error {

	var userId = utils.CreateCodeId()
	userModel := &model.User{
		UserId:   userId,
		Name:     user.Name,
		Email:    user.Email,
		PassWord: utils.Encrypt(user.PassWord),
		IdAcess:  user.IdAcess,
	}

	userInsert := structs.Map(userModel)

	_, err := client.GetInstance().Insert(ctx, "user", userInsert)
	if err != nil {
		return errors.New("Create user: problem to insert into MongoDB")
	}


	return nil
}

func (u *user) EditUser(ctx context.Context, user model.User) error {
	userModel := &model.User{
		UserId:   user.UserId,
		Name:     user.Name,
		Email:    user.Email,
		PassWord: user.PassWord,
		IdAcess:  user.IdAcess,
	}

	userUpdate:= structs.Map(userModel)

	_, err := client.GetInstance().UpdateOne(ctx, "user", userUpdate, &model.User{})
	if err != nil {
		return errors.New("Edit User: problem to update into MongoDB")
		//adicionar aqui retorno de erro (500)
	}

	return nil
}

func (u *user) GetInformationUser(ctx context.Context, id string) (model.User, error) {
	var user model.User

	userId := map[string]interface{}{"UserId": id}

	result, err := repository.Find(ctx, "user", userId, &user)
	if err != nil {
		return user, errors.New("Add Information user: problem to Find Id into MongoDB")
	}

	return result, nil
}

func (u *user) DeleteUser(ctx context.Context, id string) error {

	userId := map[string]interface{}{"UserId": id}

	err := client.GetInstance().Remove(ctx, "user", userId)
	if err != nil {
		return errors.New("Delete User: problem to delete into MongoDB")
	}

	return nil
}

func (u *user) GetUserByName(ctx context.Context, name string) (model.User, error){

	var user model.User

	Name := map[string]interface{}{"Name": name}

	result, err := repository.Find(ctx, "user", Name, &user)
	if err != nil {
		return user, errors.New("Add Information user: problem to Find Id into MongoDB")
	}

	return result, nil
}

func (u *user) GetUsersByAcess(ctx context.Context, idAcess int) ([]model.User, error){
	/*var user []model.User

	userId := map[string]interface{}{"UserId": id}

	/*result, err := repository.Find(ctx, "user", userId, &user)
	if err != nil {
		return user, errors.New("Add Information user: problem to Find Id into MongoDB")
	}*/

	return []model.User{}, nil
}