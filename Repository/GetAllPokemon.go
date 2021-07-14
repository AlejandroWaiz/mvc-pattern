package repository

import (
	models "github.com/AlejandroWaiz/mvc-pattern/Models"
)

func (r *Repository) GetAllPokemonsRepo() (AllPokemons []models.PokemonWithID, err error) {

	resp, err := r.db.Query("SELECT * FROM `allpokemons`")

	if err != nil {

		return nil, err
	}

	for resp.Next() {
		var pokemon models.PokemonWithID

		err = resp.Scan(pokemon.ID, pokemon.PokemonName, pokemon.PokemonNickname, pokemon.PrimaryType, pokemon.SecondaryType)

		if err != nil {
			return nil, err
		}

		AllPokemons = append(AllPokemons, pokemon)

		break
	}

	return AllPokemons, nil

}
