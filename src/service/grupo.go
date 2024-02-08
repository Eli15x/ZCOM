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
	instanceServiceGrupo ServiceGrupo
	onceServiceGrupo    sync.Once
)

type ServiceGrupo  interface {	
	GetGrupo(ctx context.Context, grupo string) (model.Grupo, error)
	GetGrupos(ctx context.Context) ([]model.Grupo, error)
}

type grupo struct{}

func GetInstanceGrupo() ServiceGrupo {
	onceServiceGrupo.Do(func() {
		instanceServiceGrupo = &grupo{}
	})
	return instanceServiceGrupo
}

func (m *grupo) GetGrupo(ctx context.Context, grupo string) (model.Grupo, error) {
	var Grupo model.Grupo

	if err := client.GetInstance().Ping(context.Background()); err == nil {
		grupo_interface := map[string]interface{}{"GRUPO": grupo}
		Grupo, err = repository.GetInstanceGrupo().FindOne(ctx, "grupo", grupo_interface)
		if err != nil {
			return Grupo, errors.New("Get grupo: problem to Find Id into MongoDB")
		}
	} else {
		return Grupo, err
	}
	
	return Grupo, nil
}

func (m *grupo) GetGrupos(ctx context.Context)([]model.Grupo, error){

	grupos := []model.Grupo{}
	all := map[string]interface{}{}

	if err := client.GetInstance().Initialize(context.Background()); err == nil {
		grupos, err = repository.GetInstanceGrupo().Find(ctx, "grupo", all)
		if err != nil {
			return nil, errors.New("Get grupo: problem to Find Id into MongoDB")
		}
		
	}

	return grupos, nil
}

