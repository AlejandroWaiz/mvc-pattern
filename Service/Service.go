package service

import (
	repository "github.com/AlejandroWaiz/mvc-pattern/Repository"
)

type Service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *Service {

	service := Service{
		repository: repository,
	}

	return &service

}
