package config

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Port         int    `validate:"required"`
	MongoURL     string `validate:"required"`
	TrustedProxy string `validate:"required"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	config := Config{
		Port:         GetIntEnv("PORT", 8000),
		MongoURL:     GetStringEnv("MONGO_URL", ""),
		TrustedProxy: GetStringEnv("TRUSTED_PROXY", "127.0.0.1"),
	}

	if err := validator.New().Struct(config); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed validation rule '%s'\n", err.Field(), err.Tag())
			return nil, errors.New("Config validation failed")
		}
	}

	return &config, nil
}
