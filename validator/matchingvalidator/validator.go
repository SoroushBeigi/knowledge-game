package matchingvalidator

import (
	"fmt"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/entity"
	"github.com/SoroushBeigi/knowledge-game/pkg/errmessage"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Validator struct {
}

func New() *Validator {
	return &Validator{}
}

func (v Validator) ValidateAddToMatchingRequest(req dto.AddToWaitingListRequest) (map[string]string, error) {
	const op = "matchingvalidator.ValidateAddToMatchingRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required, validation.By(v.isCategoryValid)),
	); err != nil {

		fieldErrors := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.
			New(op).
			WithMessage(errmessage.InvalidInput).
			WithCode(richerror.InvalidCode).
			WithErr(err)
	}

	return nil, nil
}

func (v Validator) isCategoryValid(value interface{}) error {
	cat := value.(entity.Category)

	if !cat.IsValid() {
		return fmt.Errorf(errmessage.InvalidCategory)
	}

	return nil
}
