package services

import (
	"api-coffee-app/models"
	"api-coffee-app/repositories"
	"api-coffee-app/requests"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type BaristaService interface {
	Register(req *requests.BaristaRequest) (*models.Barista, error)
	Login(username, password string) (*models.Barista, error)
	UpdateProfile(id string, req *requests.BaristaRequest) (*models.Barista, error)
}

type baristaService struct {
	repository repositories.BaristaRepository
}

func NewBaristaService(repository repositories.BaristaRepository) BaristaService {
	return &baristaService{repository}
}

func (s *baristaService) Register(req *requests.BaristaRequest) (*models.Barista, error) {

	barista := models.Barista{
		Fullname: req.Fullname,
		Username: req.Username,
		Email:    req.Email,
	}

	if req.Password == "" {
		return nil, errors.New("password cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	barista.Password = string(hashedPassword)
	err = s.repository.Create(&barista)
	if err != nil {
		return nil, err
	}

	return &barista, nil
}

func (s *baristaService) Login(username, password string) (*models.Barista, error) {
	barista, err := s.repository.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(barista.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return barista, nil
}

func (s *baristaService) UpdateProfile(id string, req *requests.BaristaRequest) (*models.Barista, error) {
	barista, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	barista.Fullname = req.Fullname
	barista.Username = req.Username
	barista.Email = req.Email

	updateData, err := s.repository.Update(barista)
	if err != nil {
		return updateData, err
	}

	return updateData, err

}
