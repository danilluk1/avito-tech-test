package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	DbConn string `env:"DB_CONN" env-default:"postgres://postgres:admin@localhost:5432/avito-tech"`
	AppEnv string `env:"APP_ENV"    env-default:"development"`
}

func New(isRequired bool) (*AppConfig, error) {
	err := godotenv.Load(".env")

	if err != nil {
		if isRequired {
			panic(err)
		}
	}

	var config AppConfig
	err = cleanenv.ReadEnv(&config)

	if err != nil {
		panic(err)
	}

	return &config, err
}
