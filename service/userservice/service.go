package userservice

import (
	"github.com/SoroushBeigi/knowledge-game/entity"
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
