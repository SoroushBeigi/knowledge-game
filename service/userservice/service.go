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

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	User         dto.UserInfo `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	const op = "userservice.Login"
	var defaultErr = errors.New("Phone number and password combination didn't work")

	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)

	if err != nil {
		log.Println("Service Login:", err)

		return LoginResponse{},
			richerror.New(op).
				WithErr(err).
				WithMetaData(map[string]any{"phone_number": req.PhoneNumber})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return LoginResponse{}, defaultErr
	}

	aToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		log.Println("Service Login, createToken ", err)

		return LoginResponse{}, defaultErr
	}

	rToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		log.Println("Service Login, createToken ", err)

		return LoginResponse{}, defaultErr
	}

	return LoginResponse{AccessToken: aToken,
		RefreshToken: rToken,
		User: dto.UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
		},
	}, nil
}

type GetProfileRequest struct {
	UserID uint `json:"id"`
}

type GetProfileResponse struct {
	Name string `json:"name"`
}

func (s Service) GetProfile(req GetProfileRequest) (GetProfileResponse, error) {
	const op = "userservice.GetProfile"

	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		log.Println("Service Profile:", err)

		return GetProfileResponse{},
			richerror.New(op).WithErr(err).WithMetaData(map[string]any{"req": req})
	}

	return GetProfileResponse{Name: user.Name}, nil
}
