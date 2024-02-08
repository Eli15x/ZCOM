package repository

import (
	"context"
	"errors"
	"sync"
	"fmt"

	"ZCOM/src/model"
	"ZCOM/src/client"
)

var (
	instanceRepositoryUnidade RepositoryUnidade
	onceRepositoryUnidade    sync.Once
)

type RepositoryUnidade interface {
	FindOne(ctx context.Context, collName string, query map[string]interface{})  (model.Unidade, error)
	Find(ctx context.Context, collName string, query map[string]interface{},) ([]model.Unidade, error)
}

type repositoryUnidade struct{}

func GetInstanceUnidade() RepositoryUnidade{
	onceRepositoryUnidade.Do(func() {
		instanceRepositoryUnidade = &repositoryUnidade{}
	})
	return instanceRepositoryUnidade
}

func (ru *repositoryUnidade)Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Unidade, error) {

	cursor, err := client.GetInstance().Find(ctx, collName, query)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in mongoDB")
	}

	var content []model.Unidade
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	return content, nil
}

func (ru *repositoryUnidade)FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Unidade, error){

	var Unidade model.Unidade
	result, err := client.GetInstance().FindOne(ctx, collName, query)
	fmt.Println(result)
	if err != nil {
		return Unidade, errors.New("Error Repository: Error find query in mongoDb")
	}
	result.Decode(&Unidade)

	return Unidade,nil
}



