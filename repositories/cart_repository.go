package repositories

import (
	"api-coffee-app/models"
	"log"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindByUserID(customerID string) ([]models.Cart, error)
	FindByUserAndProduct(CustomerID, ProductID string) (*models.Cart, error)
	Create(cart *models.Cart) (*models.Cart, error)
	Update(cart *models.Cart) (*models.Cart, error)
	Remove(id string) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) FindByUserID(customerID string) ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("Customer").Preload("Product").Where("customer_id = ?", customerID).Find(&cart).Error
	return cart, err
}

func (r *cartRepository) FindByUserAndProduct(CustomerID, ProductID string) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Customer").Preload("Product").Where("customer_id = ? AND product_id = ?", CustomerID, ProductID).First(&cart).Error
	log.Println("terjadi error = ", err)
	if err != nil {
		return nil, err
	}

	return &cart, err
}

func (r *cartRepository) Create(cart *models.Cart) (*models.Cart, error) {
	if err := r.db.Create(cart).Error; err != nil {
		return nil, err
	}

	return cart, nil
}

func (r *cartRepository) Update(cart *models.Cart) (*models.Cart, error) {
	err := r.db.Save(&cart).Error
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (r *cartRepository) Remove(id string) error {
	return r.db.Where("id = ?", id).Delete(&id).Error
}
