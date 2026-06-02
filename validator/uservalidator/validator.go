package uservalidator

import "github.com/SoroushBeigi/knowledge-game/entity"

const (
	phoneNumberRegex = "^[0-9]{11}$"
)

type Repository interface {
	IsPhoneNumberUnique(pn string) (bool, error)
	GetUserByPhoneNumber(pn string) (entity.User, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) *Validator {
	return &Validator{repo: repo}
}
