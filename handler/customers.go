package handler

import (
	"encoding/json"
	"net/http"
)

func responseJSON(w http.ResponseWriter, status int, payload string) {
	response, err := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(status)
	w.Write([]byte(response))
}

func CustomerHandler(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, 200, "hello, customer!")
}

func CustomerGetAllHandler(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, 200, "not implement yet!")
}

func CustomerGetByIdHandler(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, 200, "not implement yet!")
}

func CustomerGetByNameHandler(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, 200, "not implement yet!")
}

func CustomerPostHandler(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, 200, "not implement yet!")
}

func CustomerPutByIdHandler(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, 200, "not implement yet!")
}

func CustomerDeleteByIdHandler(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, 200, "not implement yet!")
}
