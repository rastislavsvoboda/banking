package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rastislavsvoboda/banking/domain"
	"github.com/rastislavsvoboda/banking/logger"
	"github.com/rastislavsvoboda/banking/service"
)

func Start() {
	router := mux.NewRouter()

	// customerHandlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerHandlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", customerHandlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandlers.getCustomer).Methods(http.MethodGet)

	// starting server
	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		logger.Error(err.Error())
	}
}
