package utils

import (
	"errors"

	"github.com/xeipuuv/gojsonschema"
)

func ValidateJsonSchema(spec map[string]any, schema map[string]any) error {
	// Load the schema and spec as JSON loaders
	schemaLoader := gojsonschema.NewGoLoader(schema)
	specLoader := gojsonschema.NewGoLoader(spec)

	// Validate the spec against the schema
	result, err := gojsonschema.Validate(schemaLoader, specLoader)
	if err != nil {
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
