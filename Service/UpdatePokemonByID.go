package service

import models "github.com/AlejandroWaiz/mvc-pattern/Models"

func (s *Service) UpdatePokemonByID(data models.PokemonWithID) error {

	err := s.repository.UpdatePokemonByID(data)

	if err != nil {

		return err

	}

	return nil

}
