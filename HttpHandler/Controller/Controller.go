package controller

import (
	"net/http"

	service "github.com/AlejandroWaiz/mvc-pattern/Service"
)

type Controller struct {
	Service service.Service
}

func NewController(service service.Service) *Controller {

	controller := Controller{
		Service: service,
	}

	return &controller

}

type IController interface {
	CreatePokemonCont(w http.ResponseWriter, r *http.Request)
}
