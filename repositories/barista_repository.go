package repositories

import (
	"api-coffee-app/models"

	"gorm.io/gorm"
)

type BaristaRepository interface {
	Create(barista *models.Barista) error
	FindByID(id string) (*models.Barista, error)
	FindByUsername(username string) (*models.Barista, error)
	Update(barista *models.Barista) (*models.Barista, error)
}

type baristaRepository struct {
	db *gorm.DB
}

func NewBaristaRepository(db *gorm.DB) BaristaRepository {
	return &baristaRepository{db}
}

func (r *baristaRepository) Create(barista *models.Barista) error {
	return r.db.Create(barista).Error
}

func (r *baristaRepository) FindByID(id string) (*models.Barista, error) {
	var barista models.Barista
	err := r.db.Where("id = ?", id).First(&barista).Error
	return &barista, err
}

func (r *baristaRepository) FindByUsername(username string) (*models.Barista, error) {
	var barista models.Barista
	err := r.db.Where("username = ?", username).First(&barista).Error
	return &barista, err
}

func (r *baristaRepository) Update(barista *models.Barista) (*models.Barista, error) {
	err := r.db.Save(&barista).Error
	if err != nil {
		return barista, err
	}

	return barista, nil
}
