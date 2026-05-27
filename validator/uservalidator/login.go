package uservalidator

import (
	"fmt"
	"regexp"

	"github.com/SoroushBeigi/knowledge-game/dto"
	"github.com/SoroushBeigi/knowledge-game/pkg/errmessage"
	"github.com/SoroushBeigi/knowledge-game/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateLoginRequest(req dto.LoginRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateLoginRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errmessage.PhoneNotValid), validation.By(v.doesPhoneExist)),
		validation.Field(&req.Password, validation.Required),
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

func (v Validator) doesPhoneExist(value interface{}) error {
	pn := value.(string)

	_, err := v.repo.GetUserByPhoneNumber(pn)
	if err != nil {

		//security practice: not exposing the existence of phone number in database
		return fmt.Errorf(errmessage.IncorrectLogin)
	}

	return nil
}
