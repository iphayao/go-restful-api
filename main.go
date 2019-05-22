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
	r.HandleFunc("/", handler.HomeHandler)
	r.HandleFunc("/customers", handler.CustomerHandler)

	fmt.Println("Server running ...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
