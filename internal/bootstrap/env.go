package bootstrap

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
)

type EnvConfig struct {
	AppEnv                 string
	ServerAddress          string
	ContextTimeout         int
	AccessTokenExpiryHour  int
	RefreshTokenExpiryHour int
	AccessTokenSecret      string
	RefreshTokenSecret     string
	GitToken               string
}

func NewEnv() *EnvConfig {
	// Load .env file first (won't override real env vars)
	_ = godotenv.Load(".env")

	// Helper to get int env
	getInt := func(key string, defaultVal int) int {
		val := os.Getenv(key)
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
		return defaultVal
	}

	env := &EnvConfig{
		AppEnv:                 os.Getenv("APP_ENV"),
		ServerAddress:          os.Getenv("SERVER_ADDRESS"),
		ContextTimeout:         getInt("CONTEXT_TIMEOUT", 30),
		AccessTokenExpiryHour:  getInt("ACCESS_TOKEN_EXPIRY_HOUR", 1),
		RefreshTokenExpiryHour: getInt("REFRESH_TOKEN_EXPIRY_HOUR", 72),
		AccessTokenSecret:      os.Getenv("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret:     os.Getenv("REFRESH_TOKEN_SECRET"),
		GitToken:               os.Getenv("GIT_TOKEN"),
	}

	log.Logger.Info("Loaded Config", "AppEnv", env.AppEnv)
	return env
}

