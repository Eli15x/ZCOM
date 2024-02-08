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
	instanceRepositoryGrupo RepositoryGrupo
	onceRepositoryGrupo    sync.Once
)

type RepositoryGrupo interface {
	FindOne(ctx context.Context, collName string, query map[string]interface{})  (model.Grupo, error)
	Find(ctx context.Context, collName string, query map[string]interface{},) ([]model.Grupo, error)
}

type repositoryGrupo struct{}

func GetInstanceGrupo() RepositoryGrupo{
	onceRepositoryGrupo.Do(func() {
		instanceRepositoryGrupo = &repositoryGrupo{}
	})
	return instanceRepositoryGrupo
}

func (rg *repositoryGrupo)Find(ctx context.Context, collName string, query map[string]interface{}) ([]model.Grupo, error) {

	cursor, err := client.GetInstance().Find(ctx, collName, query)
	if err != nil {
		return nil, errors.New("Error Repository: Error find query in mongoDB")
	}

	var content []model.Grupo
	if err = cursor.All(ctx, &content); err != nil {
		return nil, errors.New("Error Repository: Error Get Cursor information mongoDB")
	}

	return content, nil
}

func (rg *repositoryGrupo)FindOne(ctx context.Context, collName string, query map[string]interface{}) (model.Grupo, error){

	var Grupo model.Grupo
	result, err := client.GetInstance().FindOne(ctx, collName, query)
	fmt.Println(result)
	if err != nil {
		return Grupo, errors.New("Error Repository: Error find query in mongoDb")
	}
	result.Decode(&Grupo)

	return Grupo,nil
}



