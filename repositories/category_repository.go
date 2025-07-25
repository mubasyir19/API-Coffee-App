package repositories

import (
	"api-coffee-app/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]models.Category, error)
	FindByID(id string) (*models.Category, error)
	Create(category *models.Category) (models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) FindByID(id string) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("id = ?", id).Preload("Products").First(&category).Error
	return &category, err
}

func (r *categoryRepository) Create(category *models.Category) (models.Category, error) {
	if err := r.db.Create(category).Error; err != nil {
		return models.Category{}, err
	}
	return *category, nil
}
