package repositories

import (
	"api-coffee-app/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetByCustomerID(customerID string) ([]models.Order, error)
	CreateOrder(order *models.Order, details []models.OrderDetail) (*models.Order, error)
	UpdateProductStock(productID string, quantity int) error
	ClearUserCart(userID string) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) GetByCustomerID(customerID string) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("customerID = ?", customerID).Find(&orders).Error
	return orders, err
}

func (r *orderRepository) CreateOrder(order *models.Order, details []models.OrderDetail) (*models.Order, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		for i := range details {
			details[i].OrderID = order.ID
			if err := tx.Create(&details[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	order.OrderDetails = details
	return order, nil
}

func (r *orderRepository) UpdateProductStock(productID string, quantity int) error {
	return r.db.Model(&models.Product{}).Where("id = ?", productID).Update("stock", gorm.Expr("stock - ?", quantity)).Error
}

func (r *orderRepository) ClearUserCart(userID string) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Cart{}).Error
}
