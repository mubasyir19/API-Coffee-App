package services

import (
	"api-coffee-app/models"
	"api-coffee-app/repositories"
	"api-coffee-app/requests"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type CartService interface {
	GetItems(customerID string) ([]models.Cart, error)
	AddItem(cartInput *requests.CartInput) (*models.Cart, error)
	UpdateItem(cartInput *requests.CartInput) (*models.Cart, error)
	RemoveItem(id string) error
}

type cartService struct {
	repository   repositories.CartRepository
	productRepo  repositories.ProductRepository
	customerRepo repositories.CustomerRepository
}

func NewCartService(repository repositories.CartRepository, productRepo repositories.ProductRepository, customerRepo repositories.CustomerRepository) CartService {
	return &cartService{repository, productRepo, customerRepo}
}

func (s *cartService) GetItems(customerID string) ([]models.Cart, error) {
	if customerID == "" {
		return nil, errors.New("customer id required")
	}

	carts, err := s.repository.FindByUserID(customerID)

	if err != nil {
		return nil, errors.New("failed to get data cart")
	}

	return carts, nil
}

func (s *cartService) AddItem(cartInput *requests.CartInput) (*models.Cart, error) {
	if cartInput.Quantity <= 0 {
		cartInput.Quantity = 1
	}

	customer, err := s.customerRepo.FindByID(cartInput.CustomerID)
	log.Println("kenapa gk ketemu = ", err)
	if err != nil {
		return nil, errors.New("customer not found")
	}

	// checking, is product available?
	product, err := s.productRepo.FindByID(cartInput.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}
	itemTotalPrice := int(product.Price) * cartInput.Quantity

	// existingCart, err := s.repository.FindByUserAndProduct(cartInput.CustomerID, cartInput.ProductID)
	existingCart, err := s.repository.FindByUserAndProduct(customer.ID, cartInput.ProductID)
	if err == nil && existingCart != nil {
		existingCart.Quantity += cartInput.Quantity
		existingCart.TotalPrice = product.Price * float64(existingCart.Quantity)
		updateCart, err := s.repository.Update(existingCart)
		if err != nil {
			return nil, errors.New("failed to update cart")
		}
		return updateCart, nil
	}

	newCart := &models.Cart{
		ID:         uuid.NewString(),
		CustomerID: cartInput.CustomerID,
		ProductID:  cartInput.ProductID,
		Quantity:   cartInput.Quantity,
		TotalPrice: float64(itemTotalPrice),
	}
	createCart, err := s.repository.Create(newCart)
	if err != nil {
		return nil, errors.New("failed to update cart")
	}

	return createCart, nil
}

func (s *cartService) UpdateItem(cartInput *requests.CartInput) (*models.Cart, error) {
	if cartInput.Quantity <= 0 {
		return nil, errors.New("quantity must be more than zero")
	}

	// check customer
	if _, err := s.customerRepo.FindByID(cartInput.CustomerID); err != nil {
		fmt.Println("customer gk ketemu = ", err)
		return nil, errors.New("customer not found")
	}

	// check product
	if _, err := s.productRepo.FindByID(cartInput.ProductID); err != nil {
		return nil, errors.New("product not found")
	}

	// find chart base on input customerId dan productId
	existingCart, err := s.repository.FindByUserAndProduct(cartInput.CustomerID, cartInput.ProductID)
	if err != nil {
		return nil, err
	}

	existingCart.Quantity = cartInput.Quantity
	existingCart.TotalPrice = float64(cartInput.Quantity) * existingCart.Product.Price

	updatedCart, err := s.repository.Update(existingCart)
	if err != nil {
		fmt.Println("salah disini nih = ", err)
		return nil, errors.New("failed to update cart")
	}

	return updatedCart, nil
}

func (s *cartService) RemoveItem(id string) error {
	if id == "" {
		return errors.New("item not found")
	}

	if err := s.repository.Remove(id); err != nil {
		return errors.New("failed remove item")
	}

	return nil
}
