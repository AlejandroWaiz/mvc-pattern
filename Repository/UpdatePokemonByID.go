package repository

import models "github.com/AlejandroWaiz/mvc-pattern/Models"

func (r *Repository) UpdatePokemonByID(toUpdate models.PokemonWithID) error {

	q := "UPDATE `allpokemons` WHERE ID = ? SET PokemonName = ?, PokemonNickName = ?, PrimaryType = ?, SecondaryType = ?"

	_, err := r.db.Exec(q, toUpdate.ID, toUpdate.PokemonName, toUpdate.PokemonNickname, toUpdate.PrimaryType, toUpdate.SecondaryType)

	if err != nil {

		return err

	}

	return nil

}
