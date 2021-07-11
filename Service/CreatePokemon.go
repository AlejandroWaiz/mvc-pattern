package service

import (
	"encoding/json"
	"errors"

	models "github.com/AlejandroWaiz/mvc-pattern/Models"
)

func (s *Service) CreatePokemonServ(jsonBody []byte) error {

	var pokemonToCreate []models.PokemonToCreate

	err := json.Unmarshal(jsonBody, &pokemonToCreate)

	if err != nil {

		return errors.New("Invalid or null data.")

	}

	err = s.repository.CreatePokemonRepo(pokemonToCreate)

	if err != nil {

		return err

	}

	return nil

}
