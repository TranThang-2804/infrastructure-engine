package bootstrap

import (
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Logger.Fatal("Can't find the file .env : ", "err", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Logger.Fatal("Environment can't be loaded: ", "err", err)
	}

	if env.AppEnv == "development" {
		log.Logger.Info("The App is running in development env")
	}

	return &env
}
