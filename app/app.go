package app

import (
	"github.com/gorilla/mux"
	"github.com/rastislavsvoboda/banking/domain"
	"github.com/rastislavsvoboda/banking/service"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	// customerHandlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerHandlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", customerHandlers.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
