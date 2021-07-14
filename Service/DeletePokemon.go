package service

func (s *Service) DeletePokemonServ(id int) error {

	err := s.repository.DeletePokemonRepo(id)

	if err != nil {
		return err
	}

	return nil

}
