package config

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	Port         int    `validate:"required"`
	MongoURL     string `validate:"required"`
	TrustedProxy string `validate:"required"`
}

func LoadConfig() (*Config, error) {
	config := Config{
		Port:         GetIntEnv("PORT", 8000),
		MongoURL:     GetStringEnv("MONGO_URL", ""),
		TrustedProxy: GetStringEnv("TRUSTED_PROXY", "127.0.0.1"),
	}

	validate := validator.New()
	err := validate.Struct(config)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed validation rule '%s'\n", err.Field(), err.Tag())
			return nil, errors.New("Config validation failed")
		}
	}

	return &config, nil
}
