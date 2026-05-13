package userservice

import (
	"errors"
	"fmt"
	"log"

	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/pkg/phonenumber"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	IsPhoneNumberUnique(pn string) (bool, error)
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(pn string) (entity.User, error)
}

func New(repo Repository) *Service {
	return &Service{Repo: repo}
}

type Service struct {
	Repo Repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User entity.User
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid: %v", req.PhoneNumber)
	}

	if isUnique, err := s.Repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {

			return RegisterResponse{}, fmt.Errorf("unexpected error happened")
		}

		if !isUnique {

			return RegisterResponse{}, fmt.Errorf("phone number is already registered: %v", req.PhoneNumber)
		}
	}

	if len(req.Name) < 2 {
		return RegisterResponse{}, fmt.Errorf("name should be at least 2 characters")
	}

	if len(req.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf("password should be at least 8 characters")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("password should be less than 72 characters")
	}

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    string(passwordHash),
	}

	createdUser, err := s.Repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error happened")
	}

	return RegisterResponse{User: createdUser}, nil

}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	var defaultErr = errors.New("Phone number and password combination didn't work")

	user, err := s.Repo.GetUserByPhoneNumber(req.PhoneNumber)

	if err != nil {
		log.Println("Service Login:", err)

		return LoginResponse{}, defaultErr
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return LoginResponse{}, defaultErr
	}

	return LoginResponse{}, nil
}
