package controller

import (
	"net/http"

	service "github.com/AlejandroWaiz/mvc-pattern/Service"
)

type Controller struct {
	Service service.IService
}

func NewController(service service.IService) *Controller {

	controller := Controller{

		Service: service,
	}

	return &controller

}

type IController interface {
	CreatePokemonCont(w http.ResponseWriter, r *http.Request)
}
