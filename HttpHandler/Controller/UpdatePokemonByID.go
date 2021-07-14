package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	web "github.com/AlejandroWaiz/mvc-pattern/HttpHandler/APIResponse"
	models "github.com/AlejandroWaiz/mvc-pattern/Models"
)

func (c *Controller) UpdatePokemonByID(w http.ResponseWriter, r *http.Request) {

	pokemonData, err := ioutil.ReadAll(r.Body)

	if err != nil {
		web.ErrBadRequest.Send(w)
	}

	var pokemonToCreate models.PokemonWithID

	err = json.Unmarshal(pokemonData, &pokemonToCreate)

	if err != nil {
		web.ErrInvalidJSON.Send(w)
	}

	err = c.Service.UpdatePokemonByID(pokemonToCreate)

	if err != nil {
		web.ErrBadRequest.Send(w)
	}

	web.Success("Pokemon has been updated succesfully!", 200)

}
