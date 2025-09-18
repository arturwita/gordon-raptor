package config

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	Port     int    `validate:"required"`
	MongoURL string `validate:"required"`
	Test     string `validate:"required,email"`
}

func LoadConfig() (Config, error) {
	config := Config{
		Port:     GetIntEnv("PORT", 8000),
		MongoURL: GetStringEnv("MONGO_URL", "example"),
		Test:     "lol",
	}

	validate := validator.New()
	err := validate.Struct(config)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed validation rule '%s'\n", err.Field(), err.Tag())
			return config, errors.New("Config validation failed")
		}
	}

	return config, nil
}
