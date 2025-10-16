package config

import (
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewGoogleOauthConfig(cfg *AppConfig) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.GoogleOauthClientId,
		ClientSecret: cfg.GoogleOauthClientSecret,
		RedirectURL:  fmt.Sprintf("%s/auth/google/callback", cfg.BackendURL),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}
