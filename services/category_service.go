package services

import (
	"api-coffee-app/models"
	"api-coffee-app/repositories"
	"api-coffee-app/requests"
	"errors"
)

type CategoryService interface {
	GetAllCategories() ([]models.Category, error)
	GetCategory(id string) (*models.Category, error)
	CreateCategory(request *requests.CategoryRequest) (*models.Category, error)
}

type categoryService struct {
	repository repositories.CategoryRepository
}

func NewCategoryService(repository repositories.CategoryRepository) CategoryService {
	return &categoryService{repository}
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	categories, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
func (s *categoryService) GetCategory(id string) (*models.Category, error) {
	category, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) CreateCategory(request *requests.CategoryRequest) (*models.Category, error) {
	if request.Name == "" {
		return nil, errors.New("category name is required")
	}

	category := models.Category{
		Name: request.Name,
	}

	newCategory, err := s.repository.Create(&category)
	if err != nil {
		return nil, err
	}

	return &newCategory, nil
}
