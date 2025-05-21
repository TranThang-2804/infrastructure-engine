package utils

import "github.com/go-playground/validator/v10"

func ValidateStruct(s any) error {
	validator := validator.New()
	return validator.Struct(s)
}
