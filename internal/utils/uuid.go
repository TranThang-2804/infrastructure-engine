package utils

import (
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/google/uuid"
)

func GenerateUUIDv7() (string, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		log.Logger.Error("Error GenerateUUIDv7", "error", err.Error())
		return "", err
	}

	return uuid.String(), err
}
