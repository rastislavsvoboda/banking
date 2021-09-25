package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/rastislavsvoboda/banking/service"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
