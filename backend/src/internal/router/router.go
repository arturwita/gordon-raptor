package router

import (
	"gordon-raptor/src/internal/di"
	"gordon-raptor/src/internal/recipes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(deps *di.DIContainer) func(router *gin.Engine) {
	return func(router *gin.Engine) {
		api := router.Group("/recipes")
		{
			api.POST("/", recipes.CreateRecipe(deps.RecipeRepository))
		}
	}
}
