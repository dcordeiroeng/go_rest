package app

import (
	"log"
	"modulo/domain"
	"modulo/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	ch := CustomHandlers{service.NewCostumerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id}", ch.getCustomerById).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id}", ch.deleteCustomerById).Methods(http.MethodDelete)
	router.HandleFunc("/customers", ch.createCustomer).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8080", router), nil)
}
