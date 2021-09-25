package domain

import "github.com/rastislavsvoboda/banking/domain/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) ById(id string) (*Customer, *errs.AppError) {
	for _, c := range s.customers {
		if c.Id == id {
			return &c, nil
		}
	}

	return nil, errs.NewNotFoundError("Customer not found")

}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Ashish", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "Rob", City: "New Delhi", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
	}

	return CustomerRepositoryStub{
		customers: customers,
	}
}
