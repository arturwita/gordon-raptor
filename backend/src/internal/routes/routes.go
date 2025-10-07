package routes

import (
	"gordon-raptor/src/internal/di"
	"gordon-raptor/src/internal/domains/recipes"
	"gordon-raptor/src/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, deps *di.DIContainer) {
	apiKeyMiddleware := middlewares.ApiKeyAuthMiddleware(deps.Config.AdminApiKey)
	
	recipesEndpoints := router.Group("/recipes")
	{
		recipesEndpoints.GET("", recipes.GetRecipesHandler(deps.RecipeService))
		recipesEndpoints.POST("", apiKeyMiddleware, recipes.CreateRecipeHandler(deps.RecipeService))
		recipesEndpoints.PUT("/:id", apiKeyMiddleware, recipes.UpdateRecipeHandler(deps.RecipeService))
		recipesEndpoints.DELETE("/:id", apiKeyMiddleware, recipes.DeleteRecipeHandler(deps.RecipeService))
	}

	// authEndpoints := router.Group("/auth")
	// {
	// 	authEndpoints.POST("/login", auth.CreateLoginHandler(deps.AuthService))
	// }
}
