package service

import (
	models "github.com/AlejandroWaiz/mvc-pattern/Models"
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
	CreatePokemonServ([]byte) error
	GetAllPokemons() ([]byte, error)
	DeletePokemonServByID(id int) error
	UpdatePokemonByID(data models.PokemonWithID) error
}
