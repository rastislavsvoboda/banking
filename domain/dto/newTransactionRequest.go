package dto

import (
	"strings"

	"github.com/rastislavsvoboda/banking/domain/errs"
)

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

// TODO: rename to TransactionRequest and merge to one file dto/Transaction.go
type NewTransactionRequest struct {
	CustomerId      string  `json:"customer_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (r NewTransactionRequest) IsTransactionTypeWithdrawal() bool {
	if r.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}

func (r NewTransactionRequest) Validate() *errs.AppError {
	if strings.ToLower(r.TransactionType) != WITHDRAWAL && strings.ToLower(r.TransactionType) != DEPOSIT {
		return errs.NewValidationError("Transaction type can only be deposit or withdrawal")
	}

	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less that zero")
	}

	return nil
}
