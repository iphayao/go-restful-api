package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iphayao/go-restful-api/app/handler"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {

}

func (a *App) setRounter() {
	a.Router = mux.NewRouter()

	s := a.Router.PathPrefix("/api").Subrouter()

	ch := handler.CustomerHandler{}
	ch.Initialize()

	s.HandleFunc("/customers/{id}", ch.CustomerGetByIdHandler).Methods("GET")
	s.HandleFunc("/customers", ch.CustomerGetByNameHandler).Queries("name", "{name}").Methods("GET")
	s.HandleFunc("/customers", ch.CustomerGetAllHandler).Methods("GET")
	s.HandleFunc("/customers", ch.CustomerPostHandler).Methods("POST")
	s.HandleFunc("/customers/{id}", ch.CustomerPutByIdHandler).Methods("PUT")
	s.HandleFunc("/customers/{id}", ch.CustomerDeleteByIdHandler).Methods("DELETE")

}

func (a *App) Run(port string) {
	a.setRounter()

	fmt.Printf("Server running on port %s ...", port)
	log.Fatal(http.ListenAndServe(port, a.Router))
}
