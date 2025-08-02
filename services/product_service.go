package services

import (
	"api-coffee-app/models"
	"api-coffee-app/repositories"
	"api-coffee-app/requests"
	"errors"
)

type ProductService interface {
	FindAllProduct() ([]models.Product, error)
	FindProductByID(id string) (*models.Product, error)
	FindProductBySlug(slug string) (*models.Product, error)
	FindProductByName(name string) (*models.Product, error)
	FindProductByCategory(categoryID string) ([]models.Product, error)
	AddProduct(productInput *requests.ProductInput) (*models.Product, error)
}

type productService struct {
	repository repositories.ProductRepository
}

func NewProductService(repository repositories.ProductRepository) ProductService {
	return &productService{repository}
}

func (s *productService) FindAllProduct() ([]models.Product, error) {
	products, err := s.repository.FindAll()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *productService) FindProductByID(id string) (*models.Product, error) {
	return s.repository.FindByID(id)
}

func (s *productService) FindProductByName(name string) (*models.Product, error) {
	return s.repository.FindByName(name)
}

func (s *productService) FindProductBySlug(slug string) (*models.Product, error) {
	return s.repository.FindBySlug(slug)
}

func (s *productService) FindProductByCategory(categoryID string) ([]models.Product, error) {
	products, err := s.repository.FindByCategory(categoryID)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productService) AddProduct(productInput *requests.ProductInput) (*models.Product, error) {
	if productInput.Name == "" {
		return nil, errors.New("product name is required")
	}
	if productInput.CategoryID == "" {
		return nil, errors.New("product category is required")
	}
	if productInput.Description == "" {
		return nil, errors.New("product description is required")
	}
	if productInput.Price <= 0 {
		return nil, errors.New("product price is required")
	}

	product := models.Product{
		Name:        productInput.Name,
		CategoryID:  productInput.CategoryID,
		Description: productInput.Description,
		Price:       productInput.Price,
		Image:       productInput.Image,
	}

	newProduct, err := s.repository.Create(&product)
	if err != nil {
		return nil, err
	}

	return &newProduct, nil

}
