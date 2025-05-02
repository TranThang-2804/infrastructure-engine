package domain

import (
	"github.com/go-playground/validator/v10"
)

// baseDomainStruct provides a common Validate method for all structs
type baseDomainStruct struct{}

// Validate validates the struct passed to it
func (b *baseDomainStruct) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}
