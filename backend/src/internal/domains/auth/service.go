package auth

import (
	"gordon-raptor/src/internal/config"
	"gordon-raptor/src/internal/domains/users"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	GenerateJWT(user *users.UserModel) (string, error)
}

type authService struct {
	config *config.AppConfig
}

func NewAuthService(config *config.AppConfig) (AuthService, error) {
	return &authService{config}, nil
}

func (service *authService) GenerateJWT(user *users.UserModel) (string, error) {
	now := time.Now()

	claims := JwtClaims{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Picture:   user.Picture,
		Role:      user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Id.Hex(),
			Issuer:    "gordon-raptor",
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(service.config.JwtExpirationMins) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(service.config.JwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
