package controller

import (
	"io/ioutil"
	"net/http"

	web "github.com/AlejandroWaiz/mvc-pattern/HttpHandler/APIResponse"
)

func (c *Controller) CreatePokemonCont(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {

		web.ErrInvalidJSON.Send(w)

	}

	createdPokemon, err := c.Service.CreatePokemonServ(data)

	if err != nil {

		web.ErrBadRequest.Send(w)

	}

	web.Success(createdPokemon, http.StatusOK).Send(w)

}
