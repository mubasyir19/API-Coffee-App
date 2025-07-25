package services

import (
	"api-coffee-app/helpers"
	"api-coffee-app/models"
	"api-coffee-app/repositories"
	"api-coffee-app/requests"
	"errors"

	"github.com/google/uuid"
)

type OrderService interface {
	Checkout(input *requests.OrderRequest) (*models.Order, error)
}

type orderService struct {
	repository repositories.OrderRepository
}

func NewOrderService(repository repositories.OrderRepository) OrderService {
	return &orderService{repository}
}

func (s *orderService) Checkout(input *requests.OrderRequest) (*models.Order, error) {
	order := models.Order{
		ID:            uuid.New().String(),
		OrderCode:     helpers.GenerateCodeOrder(),
		CustomerID:    input.CustomerID,
		Address:       input.Address,
		PhoneNumber:   input.PhoneNumber,
		PaymentMethod: input.PaymentMethod,
		TotalPrice:    input.TotalPrice,
		AdminFee:      input.AdminFee,
	}

	var details []models.OrderDetail
	for _, item := range input.Items {
		details = append(details, models.OrderDetail{
			ID:        uuid.New().String(),
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			SubTotal:  item.SubTotal,
		})
	}

	newOrder, err := s.repository.CreateOrder(&order, details)
	if err != nil {
		return nil, errors.New("failed create order")
	}

	for _, item := range input.Items {
		if err := s.repository.UpdateProductStock(item.ProductID, item.Quantity); err != nil {
			return nil, err
		}
	}

	if err := s.repository.ClearUserCart(input.CustomerID); err != nil {
		return nil, err
	}

	order.OrderDetails = details
	return newOrder, nil
}
