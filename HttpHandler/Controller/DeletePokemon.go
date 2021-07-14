package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	web "github.com/AlejandroWaiz/mvc-pattern/HttpHandler/APIResponse"
	models "github.com/AlejandroWaiz/mvc-pattern/Models"
)

func (c *Controller) DeletePokemon(w http.ResponseWriter, r *http.Request) {

	dataBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		web.ErrBadRequest.Send(w)

	}

	var pokemonToDelete models.Pokemon

	err = json.Unmarshal(dataBody, &pokemonToDelete)

	if err != nil {

		web.ErrInvalidJSON.Send(w)

	}

}
