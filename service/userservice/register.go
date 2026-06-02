package userservice

import (
	"fmt"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/entity"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("password should be less than 72 characters")
	}

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    string(passwordHash),
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("unexpected error happened")
	}

	return dto.RegisterResponse{User: dto.UserInfo{
		ID:          createdUser.ID,
		PhoneNumber: createdUser.PhoneNumber,
		Name:        createdUser.Name,
	}}, nil

}
