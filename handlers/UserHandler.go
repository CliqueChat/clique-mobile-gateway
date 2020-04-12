package handlers

import (
	"github.com/CliqueChat/clique-mobile-gateway/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

func InitUserHandles(r *mux.Router) {
	r.HandleFunc("/user/profile", getUserProfile).Methods(helpers.GET.String())
	r.HandleFunc("/user/create", createUser).Methods(helpers.POST.String())

}

func createUser(w http.ResponseWriter, r *http.Request) {

}

func getUserProfile(w http.ResponseWriter, r *http.Request) {

}
