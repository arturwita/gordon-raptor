package routes

import (
	"gordon-raptor/src/internal/di"
	"gordon-raptor/src/internal/recipes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, deps *di.DIContainer) {
	router.POST("/recipes", recipes.CreateRecipeHandler(deps.RecipeService))
}
