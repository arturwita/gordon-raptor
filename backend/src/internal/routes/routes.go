package routes

import (
	"gordon-raptor/src/internal/di"
	"gordon-raptor/src/internal/domains/auth/google"
	"gordon-raptor/src/internal/domains/recipes"
	"gordon-raptor/src/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, deps *di.DIContainer) {
	apiKeyMiddleware := middlewares.ApiKeyAuthMiddleware(deps.Config.AdminApiKey)

	recipesEndpoints := router.Group("/recipes")
	{
		recipesEndpoints.GET("", recipes.NewGetRecipesHandler(deps.RecipeService))
		recipesEndpoints.POST("", apiKeyMiddleware, recipes.NewCreateRecipeHandler(deps.RecipeService))
		recipesEndpoints.PUT("/:id", apiKeyMiddleware, recipes.NewUpdateRecipeHandler(deps.RecipeService))
		recipesEndpoints.DELETE("/:id", apiKeyMiddleware, recipes.NewDeleteRecipeHandler(deps.RecipeService))
	}

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.GET("/google/login", google.NewGoogleLoginHandler())
		authEndpoints.GET("/google/callback", google.NewGoogleCallbackHandler())
	}
}
