package service

import (
	"encoding/json"
)

func (s *Service) GetAllPokemons() ([]byte, error) {

	AllPokemons, err := s.repository.GetAllPokemonsRepo()

	if err != nil {
		return nil, err
	}

	PokemonsRespBody, err := json.Marshal(AllPokemons)

	if err != nil {
		return nil, err
	}

	return PokemonsRespBody, nil

}
