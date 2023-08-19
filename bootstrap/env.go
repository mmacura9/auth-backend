package bootstrap

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv             string        `mapstructure:"APP_ENV"`
	ServerAddress      string        `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout     int           `mapstructure:"CONTEXT_TIMEOUT"`
	DBDriver           string        `mapstructure:"DB_DRIVER"`
	DBHost             string        `mapstructure:"DB_HOST"`
	DBPort             string        `mapstructure:"DB_PORT"`
	DBUser             string        `mapstructure:"DB_USER"`
	DBPass             string        `mapstructure:"DB_PASS"`
	DBName             string        `mapstructure:"DB_NAME"`
	DBSource           string        `mapstructure:"DB_SOURCE"`
	AccessTokenExpiry  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRY"`
	RefreshTokenExpiry time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRY"`
	AccessTokenSecret  string        `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string        `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	env := Env{}
	viper.AddConfigPath("..")
	viper.SetConfigFile("app.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
