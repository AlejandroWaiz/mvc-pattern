package infrastructure

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
)

func CreateRouter() (string, *mux.Router) {

	port := fmt.Sprintf(":%s", os.Getenv("muxport"))

	router := mux.NewRouter().StrictSlash(false)

	fmt.Println("Creating router...")

	return port, router

}
