package repository

import (
	"context"
	"errors"
	"sync"
	"fmt"

	"github.com/Eli15x/ZCOM/src/model"
	"github.com/Eli15x/ZCOM/src/client"
)

var (
	instanceRepositoryProduct RepositoryProduct
	onceRepositoryProduct    sync.Once
)

type RepositoryProduct interface {
	FindOne(ctx context.Context, collName string, query map[string]interface{})  (model.Product, error)
	Find(ctx context.Context, collName string, query map[string]interface{},) ([]model.Product, error)
}

type repositoryProduct struct{}

func GetInstanceProduct() RepositoryProduct{
	onceRepositoryProduct.Do(func() {
		instanceRepositoryProduct = &repositoryProduct{}
	})
	return instanceRepositoryProduct
}

func (rp *repositoryProduct)Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Product, error) {

	cursor, err := client.GetInstance().Find(ctx, collName, query)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in mongoDB")
	}

	var content []model.Product
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	return content, nil
}

func (rp *repositoryProduct)FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Product, error){

	var product model.Product
	result, err := client.GetInstance().FindOne(ctx, collName, query)
	fmt.Println(result)
	if err != nil {
		return product, errors.New("Error Repository: Error find query in mongoDb")
	}
	result.Decode(&product)

	return product,nil
}



