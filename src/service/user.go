package service

import (
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"context"
	"errors"
	"sync"
	"fmt"


	"github.com/Eli15x/ZCOM/src/model"
	"github.com/Eli15x/ZCOM/src/client"
	"github.com/Eli15x/ZCOM/src/repository"
	"github.com/Eli15x/ZCOM/src/utils"
	"github.com/fatih/structs"
)

var (
	instanceServiceUser ServiceUser
	onceServiceUser    sync.Once
)

type ServiceUser interface {
	ValidateUser(ctx context.Context, email string, password string) (string, error)
	CreateUser(ctx context.Context, user model.UserRequest) error
	EditUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, id string) (model.User, error)
	GetUserByName(ctx context.Context, name string) (model.User,error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUsersByAcess(ctx context.Context, idAcess int) ([]model.User, error)
	GetUsers(ctx context.Context) ([]model.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type user struct{}

func GetInstanceUser() ServiceUser {
	onceServiceUser.Do(func() {
		instanceServiceUser = &user{}
	})
	return instanceServiceUser
}

func (u *user) ValidateUser(ctx context.Context, email string, password string) (string, error) {
	
	emailValidate := map[string]interface{}{"Email": email}
	user, err := repository.GetInstanceUser().FindOne(ctx, "user", emailValidate)
	if err != nil {
		return "",errors.New("Validate user: problem to get information into MongoDB")
	}


	fmt.Println(user)
	passwordEncrypt := utils.Encrypt(password)
	fmt.Println(passwordEncrypt)
	fmt.Println(user.PassWord)
	if strings.Compare(passwordEncrypt, user.PassWord) != 0 {
		return "", errors.New("Password user: wrong password")
	}
	
	return user.UserId,nil
}

func (u *user) CreateUser(ctx context.Context, user model.UserRequest) error {

	userExist, _ := u.GetUserByEmail(ctx, user.Email)
	if userExist.UserId != "" {
		return errors.New("User Create: this email exists")
	}

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

	userExist, _ := u.GetUser(ctx, user.UserId)
	if userExist.UserId == "" {
		return errors.New("User Edit: this userId not exists")
	}

	userUpdate:= structs.Map(user)
	userId := map[string]interface{}{"UserId": user.UserId}
	change := bson.M{"$set": userUpdate}

	_, err := client.GetInstance().UpdateOne(ctx, "user", userId, change)
	if err != nil {
		return errors.New("Edit User: problem to update into MongoDB")
	}

	return nil
}

func (u *user) GetUser(ctx context.Context, id string) (model.User, error) {
	var user model.User

	userId := map[string]interface{}{"UserId": id}

	user, err := repository.GetInstanceUser().FindOne(ctx, "user", userId)
	if err != nil {
		return user, errors.New("Get user: problem to Find Id into MongoDB")
	}

	return user, nil
}

func (u *user) DeleteUser(ctx context.Context, id string) error {

	userExist, _ := u.GetUser(ctx, id)
	if userExist.UserId == "" {
		return errors.New("User Delete: this userId not exists")
	}
	userId := map[string]interface{}{"UserId": id}

	err := client.GetInstance().Remove(ctx, "user", userId)
	if err != nil {
		return errors.New("Delete User: problem to delete into MongoDB")
	}

	return nil
}

func (u *user) GetUserByName(ctx context.Context, name string) (model.User, error){

	Name := map[string]interface{}{"Name": name}
	user, err := repository.GetInstanceUser().FindOne(ctx, "user", Name)
	if err != nil {
	
		return user, errors.New("Get user by name: problem to Find Id into MongoDB")
	}

	return user, nil
}

func (u *user) GetUsersByAcess(ctx context.Context, idAcess int) ([]model.User, error){

	IdAcess := map[string]interface{}{"IdAcess": idAcess}

	users, err := repository.GetInstanceUser().Find(ctx, "user", IdAcess)
	if err != nil {
		return nil, errors.New("Get Users By Acess: problem to Find Id into MongoDB")
	}

	return users, nil
}

func (u *user) GetUserByEmail(ctx context.Context, email string) (model.User, error){

	Email := map[string]interface{}{"Email": email}

	user, err := repository.GetInstanceUser().FindOne(ctx, "user", Email)
	if err != nil {
		return model.User{}, errors.New("Get Users By Acess: problem to Find Id into MongoDB")
	}

	return user, nil
}


func (u *user) GetUsers(ctx context.Context) ([]model.User, error){

	all := map[string]interface{}{}

	users, err := repository.GetInstanceUser().Find(ctx, "user", all)
	if err != nil {
		return nil, errors.New("Get Users: problem to Find Id into MongoDB")
	}

	return users, nil
}