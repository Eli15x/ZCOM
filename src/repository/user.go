package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"errors"
	"sync"
	//"fmt"

	"github.com/Eli15x/ZCOM/src/model"
	"github.com/Eli15x/ZCOM/src/client"
)

var (
	instanceRepositoryUser RepositoryUser
	onceRepositoryUser    sync.Once
)

type RepositoryUser interface {
	FindOne(ctx context.Context, collName string, query map[string]interface{})  (model.User, error)
	Find(ctx context.Context, collName string, query map[string]interface{},) ([]bson.M, error)
}

type repositoryUser struct{}

func GetInstance() RepositoryUser {
	onceRepositoryUser.Do(func() {
		instanceRepositoryUser = &repositoryUser{}
	})
	return instanceRepositoryUser
}

func (ru *repositoryUser)Find(ctx context.Context, collName string, query map[string]interface{}) ([]bson.M, error) {

	cursor, err := client.GetInstance().Find(ctx, collName, query)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in mongoDB")
	}

	var content []bson.M
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	return content, nil
}

func (ru *repositoryUser)FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.User, error){

	var user model.User
	result, err := client.GetInstance().FindOne(ctx, collName, query)
	if err != nil {
		return user, errors.New("Error Repository: Error find query in mongoDb")
	}
	result.Decode(&user)

	return user,nil
}



