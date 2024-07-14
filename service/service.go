package service

import (
	"modulo/domain"
	"modulo/errors"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerById(string) (*domain.Customer, *errors.AppErrors)
	DeleteCustomerById(string) *errors.AppErrors
}

func (s DefaultCostumerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCostumerService) GetCustomerById(id string) (*domain.Customer, *errors.AppErrors) {
	return s.repo.ById(id)
}

func (s DefaultCostumerService) DeleteCustomerById(id string) *errors.AppErrors {
	return s.repo.DeleteById(id)
}

type DefaultCostumerService struct {
	repo domain.CustomerRepository
}

func NewCostumerService(repository domain.CustomerRepository) DefaultCostumerService {
	return DefaultCostumerService{repository}
}
