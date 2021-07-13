package service

import (
	repository "github.com/AlejandroWaiz/mvc-pattern/Repository"
)

type Service struct {
	repository repository.IRepository
}

func NewService(repository repository.IRepository) IService {

	service := Service{

		repository: repository,
	}

	return &service

}

type IService interface {
	CreatePokemonServ(jsonBody []byte) error
	GetAllPokemons() ([]byte, error)
}
