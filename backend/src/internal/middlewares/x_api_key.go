package middlewares

import (
	"gordon-raptor/src/internal/contracts"
	errors "gordon-raptor/src/internal/errors"

	"github.com/gin-gonic/gin"
)

func ApiKeyAuthMiddleware(adminApiKey string) gin.HandlerFunc {
	return func(context *gin.Context) {
		apiKey := context.GetHeader("x-api-key")
		if apiKey != adminApiKey {
			err := errors.CustomErrors.Auth.Forbidden
			context.AbortWithStatusJSON(err.Status, contracts.ErrorResponse{
				Message: err.Message,
			})
			return
		}
		context.Next()
	}
}