package services

import (
	"api-coffee-app/models"
	"api-coffee-app/repositories"
	"api-coffee-app/requests"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type CustomerService interface {
	Register(req *requests.CustomerRequest) (*models.Customer, error)
	Login(input requests.CustomerLogin) (*models.Customer, error)
	UpdateProfile(id string, req *requests.CustomerRequest) (*models.Customer, error)
	GenerateToken(customer *models.Customer) (string, error)
}

type customerService struct {
	repository repositories.CustomerRepository
}

func NewCustomerService(repository repositories.CustomerRepository) CustomerService {
	return &customerService{repository: repository}
}

func (s *customerService) Register(req *requests.CustomerRequest) (*models.Customer, error) {

	customer := models.Customer{
		Fullname:    req.Fullname,
		Username:    req.Username,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
	}

	if req.Password == "" {
		return nil, errors.New("password cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	customer.Password = string(hashedPassword)
	err = s.repository.Create(&customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (s *customerService) GenerateToken(customer *models.Customer) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("missing JWT secret key")
	}

	claims := jwt.MapClaims{
		"user_id":     customer.ID,
		"fullname":    customer.Fullname,
		"username":    customer.Username,
		"phoneNumber": customer.PhoneNumber,
		"email":       customer.Email,
		"exp":         time.Now().Add(time.Hour * 24).Unix(), // Token berlaku 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *customerService) Login(input requests.CustomerLogin) (*models.Customer, error) {
	username := input.Username
	password := input.Password

	customer, err := s.repository.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *customerService) UpdateProfile(id string, req *requests.CustomerRequest) (*models.Customer, error) {
	customer, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	customer.Fullname = req.Fullname
	customer.Username = req.Username
	customer.Email = req.Email
	customer.PhoneNumber = req.PhoneNumber
	customer.Address = req.Address

	updateData, err := s.repository.Update(customer)
	if err != nil {
		return updateData, err
	}

	return updateData, err

}
