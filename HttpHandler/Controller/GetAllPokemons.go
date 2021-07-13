package controller

import (
	"net/http"

	web "github.com/AlejandroWaiz/mvc-pattern/HttpHandler/APIResponse"
)

func (c *Controller) GetAllPokemons(w http.ResponseWriter, r *http.Request) {

	allPokemons, err := c.Service.GetAllPokemons()

	if err != nil {
		web.ErrBadRequest.Send(w)
	}

	web.Success(allPokemons, http.StatusOK).Send(w)

}
