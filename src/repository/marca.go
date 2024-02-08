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
	instanceRepositoryMarca RepositoryMarca
	onceRepositoryMarca    sync.Once
)

type RepositoryMarca interface {
	FindOne(ctx context.Context, collName string, query map[string]interface{})  (model.Marca, error)
	Find(ctx context.Context, collName string, query map[string]interface{},) ([]model.Marca, error)
}

type repositoryMarca struct{}

func GetInstanceMarca() RepositoryMarca{
	onceRepositoryMarca.Do(func() {
		instanceRepositoryMarca = &repositoryMarca{}
	})
	return instanceRepositoryMarca
}

func (rm *repositoryMarca)Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Marca, error) {

	cursor, err := client.GetInstance().Find(ctx, collName, query)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in mongoDB")
	}

	var content []model.Marca
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	return content, nil
}

func (rm *repositoryMarca)FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Marca, error){

	var Marca model.Marca
	result, err := client.GetInstance().FindOne(ctx, collName, query)
	fmt.Println(result)
	if err != nil {
		return Marca, errors.New("Error Repository: Error find query in mongoDb")
	}
	result.Decode(&Marca)

	return Marca,nil
}



