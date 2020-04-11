package handlers

import (
	"fmt"
	HTTPMethod "github.com/CliqueChat/clique-mobile-gateway/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

func InitHomeHandles(r *mux.Router) {

	r.HandleFunc("/", home).Methods(HTTPMethod.GET.String())
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serveer is up and running")
}
