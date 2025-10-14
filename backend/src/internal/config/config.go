package config

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port                    int    `validate:"required"`
	MongoURL                string `validate:"required"`
	TrustedProxy            string `validate:"required"`
	FrontendURL             string `validate:"required"`
	BackendURL              string `validate:"required"`
	GoogleOauthClientId     string `validate:"required"`
	GoogleOauthClientSecret string `validate:"required"`
	JwtSecret               string `validate:"required"`
	JwtExpirationMins       int    `validate:"required"`
}

func LoadConfig() (*AppConfig, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	config := AppConfig{
		// App
		Port:         GetIntEnv("PORT", 8000),
		MongoURL:     GetStringEnv("MONGO_URL", ""),
		TrustedProxy: GetStringEnv("TRUSTED_PROXY", "127.0.0.1"),
		FrontendURL:  GetStringEnv("FRONTEND_URL", "http://localhost:5173"),
		BackendURL:   GetStringEnv("BACKEND_URL", "http://localhost:8000"),

		// Google
		GoogleOauthClientId:     GetStringEnv("GOOGLE_OAUTH_CLIENT_ID", ""),
		GoogleOauthClientSecret: GetStringEnv("GOOGLE_OAUTH_CLIENT_SECRET", ""),

		// JWT
		JwtSecret:         GetStringEnv("JWT_SECRET", ""),
		JwtExpirationMins: GetIntEnv("JWT_EXPIRATION_MINS", 24*60),
	}

	if err := validator.New().Struct(config); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed validation rule '%s'\n", err.Field(), err.Tag())
			return nil, errors.New("config validation failed")
		}
	}

	return &config, nil
}

var TestConfig = &AppConfig{
	Port:                    8000,
	MongoURL:                "mongodb://localhost:27017/gordon_test",
	TrustedProxy:            "127.0.0.1",
	FrontendURL:             "http://localhost:5173",
	BackendURL:              "http://localhost:8000",
	GoogleOauthClientId:     "google-client-id",
	GoogleOauthClientSecret: "google-client-secret",
	JwtSecret:               "jwt-secret",
	JwtExpirationMins:       15,
}
