package middlewares

import (
	// "gordon-raptor/src/internal/contracts"
	// "gordon-raptor/src/internal/custom_errors"

	"github.com/gin-gonic/gin"
)

func ApiKeyAuthMiddleware(adminApiKey string) gin.HandlerFunc {
	return func(context *gin.Context) {
		// apiKey := context.GetHeader("x-api-key")
		// if apiKey != adminApiKey {
		// 	err := custom_errors.HttpErrors.Auth.Forbidden
		// 	context.AbortWithStatusJSON(err.Status, contracts.ErrorResponse{
		// 		Message: err.Message,
		// 	})
		// 	return
		// }
		context.Next()
	}
}
