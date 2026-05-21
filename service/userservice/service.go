package userservice

import (
	"errors"
	"fmt"
	"log"

	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/pkg/phonenumber"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	IsPhoneNumberUnique(pn string) (bool, error)
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

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserInfo struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}

type RegisterResponse struct {
	User UserInfo `json:"user"`
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid: %v", req.PhoneNumber)
	}

	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
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

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error happened")
	}

	return RegisterResponse{User: UserInfo{
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
	User         UserInfo `json:"user"`
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
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
		User: UserInfo{
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
