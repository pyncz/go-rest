package utils

import (
	"pyncz/go-rest/models/validation"

	"github.com/go-playground/validator/v10"
)

func Validate[T any](object *T) *validation.ValidationErrors {
	validate := validator.New()

	err := validate.Struct(object)
	if err != nil {
		errors := validation.ValidationErrors{}
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Error()
		}
		return &errors
	}

	return nil
}
