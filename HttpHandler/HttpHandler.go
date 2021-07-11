package httphandler

import (
	"database/sql"
	"log"
	"net/http"

	controller "github.com/AlejandroWaiz/mvc-pattern/HttpHandler/Controller"
	infrastructure "github.com/AlejandroWaiz/mvc-pattern/Infrastructure"
	repository "github.com/AlejandroWaiz/mvc-pattern/Repository"
	service "github.com/AlejandroWaiz/mvc-pattern/Service"
)

func RunRouter() {

	port, router := infrastructure.CreateRouter()

	db := infrastructure.ConnectToDataBase()

	controller := getController(db)

	router.HandleFunc("/create", controller.CreatePokemonCont)

	log.Fatal(http.ListenAndServe(port, router))

}

func getController(db *sql.DB) controller.IController {

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(service)

	return controller

}
