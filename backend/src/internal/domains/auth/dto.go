package auth

import (
	"gordon-raptor/src/internal/domains/users"

	"github.com/golang-jwt/jwt/v5"
)

type GoogleOauthUser struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Sub           string `json:"sub"`
}

type JwtClaims struct {
	Email     string     `json:"email"`
	FirstName *string    `json:"firstName,omitempty"`
	LastName  *string    `json:"lastName,omitempty"`
	Picture   *string    `json:"picture,omitempty"`
	Role      users.Role `json:"role"`
	jwt.RegisteredClaims
}
