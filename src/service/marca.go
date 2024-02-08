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
	instanceServiceMarca ServiceMarca
	onceServiceMarca    sync.Once
)

type ServiceMarca interface {	
	GetMarca(ctx context.Context, id string) (model.Marca, error)
	GetMarcas(ctx context.Context) ([]model.Marca, error)
}

type marca struct{}

func GetInstanceMarca() ServiceMarca {
	onceServiceMarca.Do(func() {
		instanceServiceMarca = &marca{}
	})
	return instanceServiceMarca
}

func (m *marca) GetMarca(ctx context.Context, nome string) (model.Marca, error) {
	var marca model.Marca

	if err := client.GetInstance().Ping(context.Background()); err == nil {
		nome := map[string]interface{}{"NOME": nome}
		marca, err = repository.GetInstanceMarca().FindOne(ctx, "marca", nome)
		if err != nil {
			return marca, errors.New("Get marca: problem to Find Id into MongoDB")
		}
	} else {
		return marca,err
	}
	
	return marca, nil
}

func (m *marca) GetMarcas(ctx context.Context)([]model.Marca, error){

	marcas := []model.Marca{}
	all := map[string]interface{}{}

	if err := client.GetInstance().Initialize(context.Background()); err == nil {
		marcas, err = repository.GetInstanceMarca().Find(ctx, "marca", all)
		if err != nil {
			return nil, errors.New("Get Marcas: problem to Find Id into MongoDB")
		}
	}

	return marcas, nil
}

