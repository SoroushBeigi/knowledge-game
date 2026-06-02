package authzservice

import (
	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
)

type Repository interface {
	GetUserPermissions(userID uint, role entity.Role) ([]string, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) HasAccess(userID uint, role entity.Role, permissions ...string) (bool, error) {
	const op = "authzService.HasAccess"
	permTitles, err := s.repo.GetUserPermissions(userID, role)
	if err != nil {
		return false, richerror.New(op).WithErr(err)
	}

	for _, pt := range permTitles {
		for _, p := range permissions {
			if p == pt {
				return true, nil
			}
		}
	}
	return false, nil

}
