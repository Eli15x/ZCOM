package repository

import (
	"context"
	"errors"

	"github.com/Eli15x/ZCOM/src/model"
	"github.com/Eli15x/ZCOM/src/client"
)

type MongoDB interface {
	Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) (model.user, error)
}

func Find(ctx context.Context, collName string, query map[string]interface{}, doc interface{}) (model.user, error) {

	cursor, err := client.GetInstance().Find(ctx, collName, query, doc)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in mongoDb")
	}

	var user model.user
	if err = cursor.All(ctx, &user); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	return user, nil
}
