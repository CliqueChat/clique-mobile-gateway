package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func InitHomeHandles(r *mux.Router) {
	r.HandleFunc("/", home).Methods(http.MethodGet)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serveer is up and running")
}
