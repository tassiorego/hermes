package internalerrors

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()
	err := validate.Struct(s)

	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	validatorError := validationErrors[0]

	if validatorError.Tag() == "required" {
		return errors.New(validatorError.Field() + " is required")
	}

	if validatorError.Tag() == "email" {
		return errors.New(validatorError.Field() + " must be a valid email")
	}

	if validatorError.Tag() == "min" {
		return errors.New(validatorError.Field() + " must be at least " + validatorError.Param() + " characters long")
	}

	if validatorError.Tag() == "max" {
		return errors.New(validatorError.Field() + " must be at most " + validatorError.Param() + " characters long")
	}

	return nil
}
