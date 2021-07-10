package controller

import (
	service "github.com/AlejandroWaiz/mvc-pattern/Service"
)

type Controller struct {
	service service.Service
}

func NewService(service service.Service) *Controller {

	controller := Controller{
		service: service,
	}

	return &controller

}
