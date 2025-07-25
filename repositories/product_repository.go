package repositories

import (
	"api-coffee-app/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(id string) (*models.Product, error)
	FindByName(name string) (*models.Product, error)
	FindByCategory(categoryID string) ([]models.Product, error)
	Create(product *models.Product) (models.Product, error)
	Update(product *models.Product) (*models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id string) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	return &product, err
}

func (r *productRepository) FindByName(name string) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("name = ?", name).First(&product).Error
	return &product, err
}

func (r *productRepository) FindByCategory(categoryID string) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("category_id = ?", categoryID).First(&products).Error
	return products, err
}

func (r *productRepository) Create(product *models.Product) (models.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return models.Product{}, err
	}
	return *product, nil
}

func (r *productRepository) Update(product *models.Product) (*models.Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
