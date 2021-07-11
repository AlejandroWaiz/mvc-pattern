package service

import (
	repository "github.com/AlejandroWaiz/mvc-pattern/Repository"
)

type Service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) IService {

	service := Service{

		repository: repository,
	}

	return &service

}

type IService interface {
	CreatePokemonServ(jsonBody []byte) ([]byte, error)
}
