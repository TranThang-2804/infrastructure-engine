package bootstrap

import (
	"os"
	"strconv"

	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/joho/godotenv"
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

	env := &EnvConfig{
		AppEnv:                 getEnv("APP_ENV", "development"),
		ServerAddress:          getEnv("SERVER_ADDRESS", ":8080"),
		ContextTimeout:         getIntEnv("CONTEXT_TIMEOUT", 30),
		AccessTokenExpiryHour:  getIntEnv("ACCESS_TOKEN_EXPIRY_HOUR", 1),
		RefreshTokenExpiryHour: getIntEnv("REFRESH_TOKEN_EXPIRY_HOUR", 72),
		GitToken:               getEnvOrPanic("GIT_TOKEN"), // Required
	}

	log.Logger.Info("Loaded Config", "AppEnv", env.AppEnv)
	return env
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func getIntEnv(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return defaultVal
}

// getEnvOrPanic gets an environment variable or panics if not set
func getEnvOrPanic(key string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	log.Logger.Fatal("Required environment variable missing", "key", key)
	return ""
}
