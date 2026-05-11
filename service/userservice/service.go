package userservice

import (
	"fmt"

	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(pn string) (bool, error)
	Register(u entity.User) (entity.User, error)
}

type Service struct {
	Repo Repository
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
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

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
	}

	createdUser, err := s.Repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error happened")
	}

	return RegisterResponse{User: createdUser}, nil

}
