package service

import (
	"context"
	"errors"
	"sync"

	"ZCOM/src/model"
	"ZCOM/src/client"
	"ZCOM/src/repository"
)

var (
	instanceServiceUnidade ServiceUnidade
	onceServiceUnidade    sync.Once
)

type ServiceUnidade interface {	
	GetUnidade(ctx context.Context, sigla string) (model.Unidade, error)
	GetUnidades(ctx context.Context) ([]model.Unidade, error)
}

type unidade struct{}

func GetInstanceUnidade() ServiceUnidade{
	onceServiceUnidade.Do(func() {
		instanceServiceUnidade = &unidade{}
	})
	return instanceServiceUnidade
}

func (u *unidade) GetUnidade(ctx context.Context, sigla string) (model.Unidade, error) {
	var unidade model.Unidade

	if err := client.GetInstance().Ping(context.Background()); err == nil {
		sigla := map[string]interface{}{"SIGLA": sigla}
		unidade, err = repository.GetInstanceUnidade().FindOne(ctx, "unidade", sigla)
		if err != nil {
			return unidade, errors.New("Get unidade: problem to Find Id into MongoDB")
		}
	} else {
		return unidade,err
	}
	
	return unidade, nil
}

func (u *unidade) GetUnidades(ctx context.Context)([]model.Unidade, error){

	unidades := []model.Unidade{}
	all := map[string]interface{}{}

	if err := client.GetInstance().Initialize(context.Background()); err == nil {
		unidades, err = repository.GetInstanceUnidade().Find(ctx, "unidade", all)
		if err != nil {
			return nil, errors.New("Get Marcas: problem to Find Id into MongoDB")
		}
		
	}

	return unidades, nil
}

