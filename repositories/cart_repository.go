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
	// err := r.db.Save(cart).Error
	err := r.db.Model(&models.Cart{}).Where("id = ?", cart.ID).
		Updates(map[string]any{
			"quantity":    cart.Quantity,
			"total_price": cart.TotalPrice,
		}).Error
	if err != nil {
		log.Println("errornya? = ", err)
		return cart, err
	}

	return cart, nil
}

func (r *cartRepository) Remove(id string) error {
	err := r.db.Delete(&models.Cart{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil

	// if err := r.db.Where("id = ?", id).Delete(&models.Cart{}).Error; err != nil {
	// 	return err
	// }

	// return nil
}
