package config

import (
	"errors"
	"fmt"
	"gordon-raptor/src/pkg/utils"

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
	JwtSecret               []byte `validate:"required"`
	JwtExpirationMins       int    `validate:"required"`
}

func LoadConfig() (*AppConfig, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	config := AppConfig{
		// App
		Port:         utils.GetIntEnv("PORT", 8000),
		MongoURL:     utils.GetStringEnv("MONGO_URL", ""),
		TrustedProxy: utils.GetStringEnv("TRUSTED_PROXY", "127.0.0.1"),
		FrontendURL:  utils.GetStringEnv("FRONTEND_URL", "http://localhost:5173"),
		BackendURL:   utils.GetStringEnv("BACKEND_URL", "http://localhost:8000"),

		// Google
		GoogleOauthClientId:     utils.GetStringEnv("GOOGLE_OAUTH_CLIENT_ID", ""),
		GoogleOauthClientSecret: utils.GetStringEnv("GOOGLE_OAUTH_CLIENT_SECRET", ""),

		// JWT
		JwtSecret:         []byte(utils.GetStringEnv("JWT_SECRET", "")),
		JwtExpirationMins: utils.GetIntEnv("JWT_EXPIRATION_MINS", 24*60),
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
	JwtSecret:               []byte("jwt-secret"),
	JwtExpirationMins:       15,
}
