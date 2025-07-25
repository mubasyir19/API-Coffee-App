package repositories

import (
	"api-coffee-app/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *models.Customer) error
	FindByID(id string) (*models.Customer, error)
	FindByEmail(email string) (*models.Customer, error)
	FindByUsername(username string) (*models.Customer, error)
	Update(customer *models.Customer) (*models.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) Create(customer *models.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) FindByID(id string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("id = ?", id).First(&customer).Error
	return &customer, err
}

func (r *customerRepository) FindByEmail(email string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("email = ?", email).First(&customer).Error
	return &customer, err
}

func (r *customerRepository) FindByUsername(username string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("username = ?", username).First(&customer).Error
	return &customer, err
}

func (r *customerRepository) Update(customer *models.Customer) (*models.Customer, error) {
	err := r.db.Save(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}
