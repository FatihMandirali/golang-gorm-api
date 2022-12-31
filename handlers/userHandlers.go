package handlers

import (
	. "appword-api/middleware"
	"github.com/gorilla/mux"
	"net/http"
)
import "appword-api/services"

func Run() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", IsAuthorized(services.GetUsers)).Methods("GET")
	r.HandleFunc("/api/login", services.Login).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}
