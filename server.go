package main

import (
	"github.com/CliqueChat/clique-mobile-gateway/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", setupRoutes()))
	log.Println("STARTED APPLICATION IN PORT : " + "8080")
}

func setupRoutes() *mux.Router {
	r := mux.NewRouter()
	handlers.InitHomeHandles(r)
	handlers.InitUserHandles(r)
	return r
}
