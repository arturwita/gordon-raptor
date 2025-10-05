package routes

import (
	"gordon-raptor/src/internal/di"
	"gordon-raptor/src/internal/domains/recipes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, deps *di.DIContainer) {
	recipesEndpoints := router.Group("/recipes")
	{
		recipesEndpoints.POST("/", recipes.CreateRecipeHandler(deps.RecipeService))
		recipesEndpoints.GET("/", recipes.GetRecipesHandler(deps.RecipeService))
	}

	// authEndpoints := router.Group("/auth")
	// {
	// 	authEndpoints.POST("/login", auth.CreateLoginHandler(deps.AuthService))
	// }
}
