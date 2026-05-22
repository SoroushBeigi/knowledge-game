package uservalidator

import (
	"fmt"
	"regexp"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/errmessage"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Repository interface {
	IsPhoneNumberUnique(pn string) (bool, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) *Validator {
	return &Validator{repo: repo}
}

func (v Validator) ValidateRegisterRequest(req dto.RegisterRequest) error {
	const op = "uservalidator.ValidateRegisterRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.Password, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z]{10,}$"))),
		validation.Field(&req.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{9}$")), validation.By(v.checkPhoneUniqueness)),
	); err != nil {
		return richerror.
			New(op).
			WithMessage(errmessage.InvalidInput).
			WithCode(richerror.InvalidCode).
			WithErr(err)
	}

	return nil
}

func (v Validator) checkPhoneUniqueness(value interface{}) error {
	pn := value.(string)

	if isUnique, err := v.repo.IsPhoneNumberUnique(pn); err != nil || !isUnique {
		if err != nil {
			return err
		}

		if !isUnique {
			return fmt.Errorf(errmessage.PhoneNotUnique)
		}
	}

	return nil
}
