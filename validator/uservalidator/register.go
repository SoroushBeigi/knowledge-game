package uservalidator

import (
	"fmt"
	"regexp"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/errmessage"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateRegisterRequest(req dto.RegisterRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateRegisterRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50).Error(errmessage.NameLength)),
		validation.Field(&req.Password, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z0-9!@#$%^&*()]{8,}$")).Error(errmessage.PasswordLength)),
		validation.Field(&req.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errmessage.PhoneNotValid), validation.By(v.checkPhoneUniqueness)),
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
