package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iphayao/go-restful-api/handler"
)

func main() {
	r := mux.NewRouter()

	s := r.PathPrefix("/api").Subrouter()

	ch := handler.CustomerHandler{}
	ch.Initialize()

	s.HandleFunc("/customers/{id}", ch.CustomerGetByIdHandler).Methods("GET")
	s.HandleFunc("/customers", ch.CustomerGetByNameHandler).Queries("name", "{name}").Methods("GET")
	s.HandleFunc("/customers", ch.CustomerGetAllHandler).Methods("GET")
	s.HandleFunc("/customers", ch.CustomerPostHandler).Methods("POST")
	s.HandleFunc("/customers/{id}", ch.CustomerPutByIdHandler).Methods("PUT")
	s.HandleFunc("/customers/{id}", ch.CustomerDeleteByIdHandler).Methods("DELETE")

	fmt.Println("Server running ...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
