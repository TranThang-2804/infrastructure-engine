package bootstrap

import (
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/spf13/viper"
)

type EnvConfig struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	GitToken               string `mapstructure:"GIT_TOKEN"`
}

func NewEnv() *EnvConfig {
	env := EnvConfig{}
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	// Unmarshal the configuration into the Env struct
	if err := viper.Unmarshal(&env); err != nil {
		log.Logger.Fatal("Environment can't be loaded: ", "err", err)
	}

	// Log the environment mode
	if env.AppEnv == "development" {
		log.Logger.Info("The App is running in development env")
	}

	return &env
}
