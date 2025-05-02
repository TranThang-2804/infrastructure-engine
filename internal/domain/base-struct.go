package domain

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// BaseDomainStruct provides a common Validate method for all structs
type BaseDomainStruct struct{}

// Validate validates the struct passed to it
func (b *BaseDomainStruct) Validate() error {
	validate := validator.New()
  fmt.Print(b)
	return validate.Struct(b)
}
