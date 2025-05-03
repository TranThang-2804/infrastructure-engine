package utils

import "github.com/go-playground/validator/v10"

func ValidateStruct(s interface{}) error {
	validator := validator.New()
	return validator.Struct(s)
}
