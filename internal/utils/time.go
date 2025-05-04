package utils

import (
	"time"

	"github.com/TranThang-2804/infrastructure-engine/internal/shared/env"
)

func GetContextTimeout() time.Duration {
	return time.Duration(env.Env.ContextTimeout) * time.Second
}
