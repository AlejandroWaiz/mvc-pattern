package repository

import models "github.com/AlejandroWaiz/mvc-pattern/Models"

func (r *Repository) CreatePokemonRepo(pokemons []models.PokemonToCreate) error {

	for _, newPokemon := range pokemons {

		q := "INSERT INTO `allcards` (ID, PokemonName, PokemonNickname, PrimaryType, SecondaryType) VALUES (?, ?, ?, ?, ?)"

		insert, err := r.db.Prepare(q)

		defer insert.Close()

		if err != nil {

			return err

		}

		_, err = insert.Exec(newPokemon.PokemonName, newPokemon.PokemonNickname, newPokemon.PrimaryType,
			newPokemon.SecondaryType)

		if err != nil {

			return err

		}
	}

	return nil

}
