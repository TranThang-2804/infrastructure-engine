package utils

import (
	"github.com/google/uuid"
)

func GenerateUUIDv7() (string, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	return uuid.String(), err
}
