package service

import (
	"log"

	"github.com/devianwahyu/efishery/domain/model"
	"github.com/devianwahyu/efishery/repository"
)

type ICustomerService interface {
	GetCustomersService() ([]model.Customer, error)
	GetCustomerByIDService(id uint) (model.Customer, error)
	AddCustomerService(customer model.Customer) (model.Customer, error)
	UpdateCustomerService(id int, customer model.Customer) (model.Customer, error)
	DeleteCustomerService(id uint) (string, error)
}

type CustomerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo}
}

func (s *CustomerService) GetCustomersService() ([]model.Customer, error) {
	customers, err := s.repo.GetCustomers()

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	log.Println("Get customers success")
	return customers, nil
}

func (s *CustomerService) GetCustomerByIDService(id uint) (model.Customer, error) {
	customer, err := s.repo.GetCustomerByID(id)

	if err != nil {
		log.Println(err.Error())
		return customer, err
	}

	log.Println("Get customer success")
	return customer, nil
}

func (s *CustomerService) AddCustomerService(customer model.Customer) (model.Customer, error) {
	customer, err := s.repo.AddCustomer(customer)

	if err != nil {
		log.Println(err.Error())
		return customer, err
	}

	log.Println("Add customer success")
	return customer, nil
}

func (s *CustomerService) UpdateCustomerService(id uint, c model.Customer) (model.Customer, error) {
	customer, err := s.repo.UpdateCustomer(id, c)

	if err != nil {
		log.Println(err.Error())
		return customer, err
	}

	log.Println("Update customer success")
	return customer, nil
}

func (s *CustomerService) DeleteCustomerService(id uint) (string, error) {
	message, err := s.repo.DeleteCustomer(id)

	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	log.Println("Delete customer success")
	return message, nil
}
