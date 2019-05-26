package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/iphayao/go-restful-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type CustomerHandler struct {
	DB *gorm.DB
}

func (c *CustomerHandler) Initialize() {
	db, err := gorm.Open("mysql", "webservice:P@ssw0rd@tcp(127.0.0.1:3306)/db_webservice?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Customer{})

	c.DB = db
}

func responseJSON(w http.ResponseWriter, status int, payload interface{}) {
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

func responseError(w http.ResponseWriter, code int, message string) {
	responseJSON(w, code, map[string]string{"error": message})
}

func (h *CustomerHandler) CustomerGetAllHandler(w http.ResponseWriter, r *http.Request) {
	customers := []model.Customer{}
	h.DB.Find(&customers)

	responseJSON(w, http.StatusOK, customers)
}

func (h *CustomerHandler) CustomerGetByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	customer := model.Customer{}
	if err := h.DB.First(&customer, id).Error; err != nil {
		responseError(w, http.StatusNotFound, err.Error())
		return
	}

	responseJSON(w, http.StatusOK, customer)
}

func (h *CustomerHandler) CustomerGetByNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	customer := model.Customer{}
	if err := h.DB.First(&customer, model.Customer{FirstName: name}).Error; err != nil {
		responseError(w, http.StatusNotFound, err.Error())
		return
	}

	responseJSON(w, http.StatusOK, customer)
}

func (h *CustomerHandler) CustomerPostHandler(w http.ResponseWriter, r *http.Request) {
	customer := model.Customer{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&customer); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.DB.Save(&customer).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJSON(w, http.StatusCreated, customer)
}

func (h *CustomerHandler) CustomerPutByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	customer := model.Customer{}
	if err := h.DB.First(&customer, id).Error; err != nil {
		responseError(w, http.StatusNotFound, err.Error())
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&customer); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.DB.Save(&customer).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJSON(w, http.StatusOK, customer)
}

func (h *CustomerHandler) CustomerDeleteByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	customer := model.Customer{}
	if err := h.DB.First(&customer, id).Error; err != nil {
		responseError(w, http.StatusNotFound, err.Error())
		return
	}

	if err := h.DB.Delete(&customer).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJSON(w, http.StatusNoContent, nil)
}
