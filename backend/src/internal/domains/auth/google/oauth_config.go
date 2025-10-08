package google

import (
	"fmt"
	"gordon-raptor/src/internal/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// todo: consider moving to AppConfig, to validate env vars

var GoogleOauthConfig = &oauth2.Config{
	ClientID:     config.GetStringEnv("GOOGLE_OAUTH_CLIENT_ID", ""),
	ClientSecret: config.GetStringEnv("GOOGLE_OAUTH_CLIENT_SECRET", ""),
	RedirectURL:  fmt.Sprintf("%s/auth/google/callback", config.GetStringEnv("BACKEND_URL", "")),
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}
