package middlewares

import (
	"gordon-raptor/src/internal/contracts"
	"gordon-raptor/src/internal/custom_errors"
	"gordon-raptor/src/internal/domains/auth"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func IsLoggedInMiddleware(jwtSecret []byte) gin.HandlerFunc {
	return func(context *gin.Context) {
		unauthorizedError := custom_errors.HttpErrors.Auth.Unauthorized
		unauthorizedResponse := &contracts.ErrorResponse{Message: unauthorizedError.Message}

		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.AbortWithStatusJSON(unauthorizedError.Status, unauthorizedResponse)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			context.AbortWithStatusJSON(unauthorizedError.Status, unauthorizedResponse)
			return
		}

		claims := &auth.JwtClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			context.AbortWithStatusJSON(unauthorizedError.Status, unauthorizedResponse)
			return
		}

		context.Set("claims", claims)

		context.Next()
	}
}
