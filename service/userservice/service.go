package userservice

import (
	"errors"
	"fmt"
	"log"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(pn string) (entity.User, error)
	GetUserByID(id uint) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

func New(repo Repository, auth AuthGenerator) *Service {
	return &Service{repo: repo, auth: auth}
}

type Service struct {
	auth AuthGenerator
	repo Repository
}

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



func (s Service) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	const op = "userservice.Login"
	var defaultErr = errors.New("Phone number and password combination didn't work")

	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)

	if err != nil {
		log.Println("Service Login:", err)

		return dto.LoginResponse{},
			richerror.New(op).
				WithErr(err).
				WithMetaData(map[string]any{"phone_number": req.PhoneNumber})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return dto.LoginResponse{}, defaultErr
	}

	aToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		log.Println("Service Login, createToken ", err)

		return dto.LoginResponse{}, defaultErr
	}

	rToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		log.Println("Service Login, createToken ", err)

		return dto.LoginResponse{}, defaultErr
	}

	return dto.LoginResponse{AccessToken: aToken,
		RefreshToken: rToken,
		User: dto.UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
		},
	}, nil
}

func (s Service) GetProfile(req dto.GetProfileRequest) (dto.GetProfileResponse, error) {
	const op = "userservice.GetProfile"

	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		log.Println("Service Profile:", err)

		return dto.GetProfileResponse{},
			richerror.New(op).WithErr(err).WithMetaData(map[string]any{"req": req})
	}

	return dto.GetProfileResponse{Name: user.Name}, nil
}
