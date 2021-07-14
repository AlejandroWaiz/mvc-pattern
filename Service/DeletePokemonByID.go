package service

func (s *Service) DeletePokemonServByID(id int) error {

	err := s.repository.DeletePokemonRepoByID(id)

	if err != nil {
		return err
	}

	return nil

}
