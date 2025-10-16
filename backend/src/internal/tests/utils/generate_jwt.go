package tests_utils

import (
	"fmt"
	"gordon-raptor/src/internal/config"
	"gordon-raptor/src/internal/domains/auth"
	"gordon-raptor/src/internal/domains/users"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTestJWT(user *users.UserModel) string {
	now := time.Now()

	claims := auth.JwtClaims{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Picture:   user.Picture,
		Role:      user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Id.Hex(),
			Issuer:    "gordon-raptor",
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(config.TestConfig.JwtExpirationMins) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(config.TestConfig.JwtSecret)
	if err != nil {
		fmt.Println("Failed to create test JWT", err)
		panic(err)
	}

	return token
}
