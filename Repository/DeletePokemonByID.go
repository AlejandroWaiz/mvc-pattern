package repository

func (r *Repository) DeletePokemonRepoByID(id int) error {

	q := "DELETE FROM `allpokemons` WHERE ID = ?"

	_, err := r.db.Exec(q, id)

	if err != nil {
		return err
	}

	return nil

}
