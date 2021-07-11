package service

import (
	"encoding/json"
	"errors"

	models "github.com/AlejandroWaiz/mvc-pattern/Models"
)

func (s *Service) CreatePokemonServ(jsonBody []byte) ([]byte, error) {

	var pokemonToCreate models.PokemonToCreate

	err := json.Unmarshal(jsonBody, &pokemonToCreate)

	if err != nil {

		return nil, errors.New("Invalid or null data.")

	}

	return nil, nil

}