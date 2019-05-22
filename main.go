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

	s.HandleFunc("/customers", handler.CustomerGetAllHandler).Methods("GET")
	s.HandleFunc("/customers/{id}", handler.CustomerGetByIdHandler).Methods("GET")
	s.HandleFunc("/customers", handler.CustomerGetByNameHandler).Queries("name", "{name}").Methods("GET")
	s.HandleFunc("/customers", handler.CustomerPostHandler).Methods("POST")
	s.HandleFunc("/customers/{id}", handler.CustomerPutByIdHandler).Methods("PUT")
	s.HandleFunc("/customers/{id}", handler.CustomerDeleteByIdHandler).Methods("DELETE")

	fmt.Println("Server running ...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
