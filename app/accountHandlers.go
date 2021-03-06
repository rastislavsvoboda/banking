package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rastislavsvoboda/banking/dto"
	"github.com/rastislavsvoboda/banking/service"
)

type AccountHandlers struct {
	service service.AccountService
}

func (ah AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		newAccountResponse, err := ah.service.NewAccount(request)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, newAccountResponse)
		}
	}
}

func (ah AccountHandlers) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	accountId := vars["account_id"]
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.CustomerId = customerId
		account, err := ah.service.MakeTransaction(request)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}
