package env

import (
	"os"
	"strconv"

	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/joho/godotenv"
)

var Env *EnvConfig

type EnvConfig struct {
	AppEnv         string
	ServerAddress  string
	ContextTimeout int
	GitToken       string
	CI             string
	NATS_URL       string

	GIT_MASTERDATA_REPO_OWNER string
	GIT_MASTERDATA_REPO_NAME  string
	GIT_RESOURCE_REPO_OWNER   string
	GIT_RESOURCE_REPO_NAME    string
}

func LoadEnv() {
	// Load .env file first (won't override real env vars)
	_ = godotenv.Load(".env")

	env := &EnvConfig{
		AppEnv:         getEnv("APP_ENV", "development"),
		ServerAddress:  getEnv("SERVER_ADDRESS", ":8080"),
		ContextTimeout: getIntEnv("CONTEXT_TIMEOUT", 30),
		GitToken:       getEnvOrPanic("GIT_TOKEN"), // Required
		CI:             getEnv("CI_TYPE", "github"),
		NATS_URL:       getEnv("NATS_URL", "nats://localhost:4222"),

		GIT_MASTERDATA_REPO_OWNER: getEnvOrPanic("GIT_MASTERDATA_REPO_OWNER"),
		GIT_MASTERDATA_REPO_NAME:  getEnvOrPanic("GIT_MASTERDATA_REPO_NAME"),
		GIT_RESOURCE_REPO_OWNER:   getEnvOrPanic("GIT_RESOURCE_REPO_OWNER"),
		GIT_RESOURCE_REPO_NAME:    getEnvOrPanic("GIT_RESOURCE_REPO_NAME"),
	}

	log.BaseLogger.Info("Loaded Config", "AppEnv", env.AppEnv)
	Env = env
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
	log.BaseLogger.Fatal("Required environment variable missing", "key", key)
	return ""
}
