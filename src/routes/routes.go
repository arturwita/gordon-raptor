package routes

import (
	"gordon-raptor/src/handlers"
	"gordon-raptor/src/pkg/di"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(deps *di.DIContainer) func(router *gin.Engine) {
	return func(router *gin.Engine) {
		api := router.Group("/api")
		{
			api.GET("/ping", handlers.Ping)
			api.POST("/recipes", handlers.NewCreateRecipe(deps.RecipeRepository))
		}
	}
}
