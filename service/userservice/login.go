package userservice

import (
	"errors"
	"log"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

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
