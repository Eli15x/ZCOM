package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Eli15x/MovieWorkNow/src/models"
	"github.com/Eli15x/MovieWorkNow/src/storage"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoDB interface {
	Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) ([]bson.M, error)
	FindFriend(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) (models.Friend, error)
}

func Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) ([]bson.M, error) {

	cursor, err := storage.GetInstance().Find(ctx, collName, query, doc)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in mongoDb")
	}

	var content []bson.M
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	return content, nil
}

func FindFriend(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) (models.Friend, error) {

	cursor, err := storage.GetInstance().Find(ctx, collName, query, doc)
	if err != nil {
		return models.Friend{}, errors.New("Error Repository: Error find query in mongoDb")
	}

	var content models.Friend
	if err = cursor.All(nil, content); err != nil {
		fmt.Println(err)
		return models.Friend{}, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	// a ideia era com que essa função retornasse para mim  o retorno que seria um userId_user e uma string de ids que seriam o userId.
	// essa ideia tem como questão o agrupamento de ids.

	return content, nil
}
package repository

import (
	"context"
	"errors"
	"fmt"

	"ZCOM/src/models"
	"ZCOM/src/client"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoDB interface {
	Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) (models.user, error)
}

func Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) (models.user, error) {

	cursor, err := client.GetInstance().Find(ctx, collName, query, doc)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in mongoDb")
	}

	var user models.user
	if err = cursor.All(ctx, &user); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	return user, nil
}
