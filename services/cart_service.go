package services

import (
	"api-coffee-app/models"
	"api-coffee-app/repositories"
	"api-coffee-app/requests"
	"errors"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartService interface {
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

	product, err := s.productRepo.FindByID(cartInput.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	newTotalPrice := product.Price * float64(cartInput.Quantity)

	existingCart, err := s.repository.FindByUserAndProduct(cartInput.CustomerID, cartInput.ProductID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err // error beneran selain record not found
	}

	existingCart.Quantity = cartInput.Quantity
	existingCart.TotalPrice = newTotalPrice

	updatedCart, err := s.repository.Update(existingCart)
	if err != nil {
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
