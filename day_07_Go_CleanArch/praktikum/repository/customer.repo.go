package repository

import (
	"log"

	"github.com/devianwahyu/efishery/domain/model"
	"gorm.io/gorm"
)

type ICustomerRepository interface {
	GetCustomers() ([]model.Customer, error)
	GetCustomerByID(id uint) (model.Customer, error)
	AddCustomer(customer model.Customer) (model.Customer, error)
	UpdateCustomer(id uint, customer model.Customer) (model.Customer, error)
	DeleteCustomer(id uint) (string, error)
}

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db}
}

func (r *CustomerRepository) GetCustomers() ([]model.Customer, error) {
	var customers []model.Customer

	if err := r.db.Find(&customers).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}

	log.Println("Get customers success")
	return customers, nil
}

func (r *CustomerRepository) GetCustomerByID(id uint) (model.Customer, error) {
	var customer model.Customer

	if err := r.db.First(&customer, id).Error; err != nil {
		log.Println(err.Error())
		return customer, err
	}

	log.Println("Get customer success")
	return customer, nil
}

func (r *CustomerRepository) AddCustomer(customer model.Customer) (model.Customer, error) {
	if err := r.db.Create(&customer).Error; err != nil {
		log.Println(err.Error())
		return customer, err
	}

	log.Println("Add new customer success")
	return customer, nil
}

func (r *CustomerRepository) UpdateCustomer(id uint, c model.Customer) (model.Customer, error) {
	var customer model.Customer

	if err := r.db.First(&customer, id).Error; err != nil {
		log.Println(err.Error())
		return customer, err
	}

	customer = c

	if err := r.db.Where("id = ?", id).Updates(&customer).Error; err != nil {
		log.Println(err.Error())
		return customer, err
	}

	log.Println("Update customer success")
	return customer, nil
}

func (r *CustomerRepository) DeleteCustomer(id uint) (string, error) {
	var customer model.Customer

	if err := r.db.Where("id = ?", id).Delete(&customer).Error; err != nil {
		log.Println(err.Error())
		return "", err
	}

	log.Println("Delete customer success")
	return "Deleted successfuly", nil
}
