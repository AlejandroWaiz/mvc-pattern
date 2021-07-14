package service

import (
	"encoding/json"
	"errors"

	models "github.com/AlejandroWaiz/mvc-pattern/Models"
)

func (s *Service) CreatePokemonServ(jsonBody []byte) error {

	var Pokemon []models.Pokemon

	err := json.Unmarshal(jsonBody, &Pokemon)

	if err != nil {

		return errors.New("Invalid or null data.")

	}

	err = s.repository.CreatePokemonRepo(Pokemon)

	if err != nil {

		return err

	}

	return nil

}
