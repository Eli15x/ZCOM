package repository

import (
	"context"
	"errors"

	"github.com/Eli15x/ZCOM/src/model"
	"github.com/Eli15x/ZCOM/src/client"
)

type MongoDB interface {
	Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) (model.User, error)
}

func Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) (model.User, error) {
	var user model.User

	cursor, err := client.GetInstance().Find(ctx, collName, query, doc)
	if err != nil {
		return user, errors.New("Error Repository: Error find query in mongoDb")
	}

	if err = cursor.All(ctx, &user); err != nil {
		return user, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	return user, nil
}

