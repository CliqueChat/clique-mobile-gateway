package main

import (
	"github.com/CliqueChat/clique-common-lib/helpers"
	"github.com/CliqueChat/clique-mobile-gateway/handlers"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties"
	"log"
	"net/http"
)

var prop = properties.MustLoadFile("./properties/config.properties", properties.UTF8)

func main() {

	setupRoutes()

	// Read host and port from property file
	host, _ := prop.Get(helpers.HOST)
	port, _ := prop.Get(helpers.PORT)

	log.Println("STARTING APPLICATION IN " + host + ":" + port)
	log.Fatal(http.ListenAndServe(host+":"+port, setupRoutes()))

}

func setupRoutes() *mux.Router {
	r := mux.NewRouter()
	handlers.InitHomeHandles(r)
	handlers.InitUserHandles(r)
	return r
}
