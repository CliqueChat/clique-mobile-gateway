package main

import (
	"github.com/CliqueChat/clique-mobile-gateway/handlers"
	"github.com/CliqueChat/clique-mobile-gateway/helpers"
	"github.com/CliqueChat/clique-mobile-gateway/resources"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

var prop = resources.GetApplicationProfile()

func main() {

	setupRoutes()

	// Read host and port from property file
	host, _ := prop.Get(helpers.HOST)
	port, _ := prop.Get(helpers.PORT)
	tcpPort, _ := prop.Get(helpers.TcpHost)

	log.Println("APPLICATION STARTED ON " + host + ":" + port)
	log.Printf("STARTED TCP CONNECTION ON PORT: " + tcpPort)

	log.Fatal(http.ListenAndServe(host+":"+port, setupRoutes()))
	log.Fatal(net.Listen("tcp_clique", "4000"))

}

func setupRoutes() *mux.Router {
	r := mux.NewRouter()
	handlers.InitHomeHandles(r)
	handlers.InitUserHandles(r)
	return r
}
