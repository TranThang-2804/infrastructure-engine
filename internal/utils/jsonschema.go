package utils

import (
	"errors"

	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/xeipuuv/gojsonschema"
)

func ValidateJsonSchema(spec string, schema string) error {
	// Load the schema and spec as JSON loaders
	schemaLoader := gojsonschema.NewStringLoader(schema)
	specLoader := gojsonschema.NewStringLoader(spec)

	// Validate the spec against the schema
	result, err := gojsonschema.Validate(schemaLoader, specLoader)
	if err != nil {
		log.Logger.Error("ValidateJsonSchema", "err", "Error validating JSON schema:", "err detail", err)
		return err // Return error if validation process fails
	}

	// Check if the validation result is valid
	if !result.Valid() {
		// Collect all validation errors
		var validationErrors string
		for _, desc := range result.Errors() {
			validationErrors += desc.String() + "\n"
		}
		return errors.New("JSON schema validation failed:\n" + validationErrors)
	}

	return nil // Return nil if validation is successful
}
