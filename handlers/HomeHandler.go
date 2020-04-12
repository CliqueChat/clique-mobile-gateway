package handlers

import (
	"fmt"
	"github.com/CliqueChat/clique-common-lib/"
	"github.com/gorilla/mux"
	"net/http"
)

func InitHomeHandles(r *mux.Router) {
	fmt.Println(helpers)
	r.HandleFunc("/", home).Methods(http.MethodGet)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serveer is up and running")
}
