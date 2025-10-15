package middlewares

import (
	"fmt"
	"gordon-raptor/src/internal/contracts"
	"gordon-raptor/src/internal/custom_errors"
	"gordon-raptor/src/internal/domains/auth"
	"gordon-raptor/src/internal/domains/users"

	"github.com/gin-gonic/gin"
)

func IsAdminMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		customError :=custom_errors.HttpErrors.Auth.Forbidden

		claims, exists := context.Get("claims")
		if !exists {
			context.AbortWithStatusJSON(customError.Status, contracts.ErrorResponse{Message: customError.Message})
			return
		}

		jwtClaims, ok := claims.(*auth.JwtClaims)
		if !ok {
			fmt.Println("failed to cast claims", claims)
			context.AbortWithStatusJSON(customError.Status, contracts.ErrorResponse{Message: customError.Message})
			return
		}

		if jwtClaims.Role != users.AdminRole {
			fmt.Println("failed to assert role")
			context.AbortWithStatusJSON(customError.Status, contracts.ErrorResponse{Message: customError.Message})
			return
		}

		context.Next()
	}
}
