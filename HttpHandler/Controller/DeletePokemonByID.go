package controller

import (
	"log"
	"net/http"
	"strconv"

	web "github.com/AlejandroWaiz/mvc-pattern/HttpHandler/APIResponse"
	"github.com/gorilla/mux"
)

func (c *Controller) DeletePokemonByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	pokemonID, err := strconv.Atoi(vars["id"])

	if err != nil {

		web.ErrBadRequest.Send(w)

	}

	err = c.Service.DeletePokemonServByID(pokemonID)

	if err != nil {

		log.Println(err)
		web.ErrBadRequest.Send(w)

	}

	web.Success("Pokemon has been deleted succesfully!", 200).Send(w)

}
