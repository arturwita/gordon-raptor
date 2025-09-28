package routes

import (
	"gordon-raptor/src/handlers"
	"gordon-raptor/src/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterRoutesFactory(recipeRepo repositories.RecipeRepository) func(router *gin.Engine) {
	return func(router *gin.Engine) {
		api := router.Group("/api")
		{
			api.GET("/ping", handlers.Ping)
			api.POST("/recipes", handlers.CreateRecipeFactory(recipeRepo))
		}
	}
}
