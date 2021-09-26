package service

import (
	"time"

	"github.com/rastislavsvoboda/banking/domain"
	"github.com/rastislavsvoboda/banking/domain/dto"
	"github.com/rastislavsvoboda/banking/domain/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	account := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		Status:      "1",
	}

	newAccount, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}

	resp := newAccount.ToNewAccountResponseDto()

	return &resp, nil
}

func (s DefaultAccountService) MakeTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	// incoming request validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	// server side validation for checking the available balance in the account
	if req.IsTransactionTypeWithdrawal() {
		// TODO: rename to FindBy
		account, err := s.repo.ById(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}

	// if all is well, build the domain object & save the transaction
	transaction := domain.Transaction{
		// TransactionId:   "",
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		// TODO: extract format dbTSLayout
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	newTransaction, err := s.repo.SaveTransaction(transaction)
	if err != nil {
		return nil, err
	}

	// TODO: rename to ToDto
	resp := newTransaction.ToNewTransactionResponseDto()

	return &resp, nil
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{
		repo: repository,
	}
}