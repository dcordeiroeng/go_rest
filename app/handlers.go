package app

import (
	"encoding/json"
	"encoding/xml"
	"modulo/errors"
	"modulo/service"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomHandlers struct {
	service service.CustomerService
}

func (ch *CustomHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		chooseContentType(w, r, err)
		return
	}
	chooseContentType(w, r, customers)
}

func (ch *CustomHandlers) getCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomerById(id)
	if err != nil {
		chooseContentType(w, r, err)
		return
	} else {
		chooseContentType(w, r, customer)
	}
}

func (ch *CustomHandlers) deleteCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	err := ch.service.DeleteCustomerById(id)
	if err != nil {
		chooseContentType(w, r, err)
		return
	}
}

func (ch *CustomHandlers) createCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	err := ch.service.DeleteCustomerById(id)
	if err != nil {
		chooseContentType(w, r, err)
		return
	}
}

func chooseContentType(w http.ResponseWriter, r *http.Request, data interface{}) {
	if err, ok := data.(*errors.AppErrors); ok {
		w.WriteHeader(err.Code)
		data = err.AsMessage()
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(data)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}
