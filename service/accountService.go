package service

import (
	"time"

	"github.com/rastislavsvoboda/banking/domain"
	"github.com/rastislavsvoboda/banking/domain/dto"
	"github.com/rastislavsvoboda/banking/domain/errs"
	"github.com/rastislavsvoboda/banking/logger"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
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

	logger.Info(account.CustomerId)

	newAccount, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}

	resp := newAccount.ToNewAccountResponseDto()

	return &resp, nil
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{
		repo: repository,
	}
}
