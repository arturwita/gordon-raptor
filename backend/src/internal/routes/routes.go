package routes

import (
	"gordon-raptor/src/internal/di"
	"gordon-raptor/src/internal/domains/auth/google"
	"gordon-raptor/src/internal/domains/recipes"
	"gordon-raptor/src/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, deps *di.DIContainer) {
	isLoggedIn := middlewares.IsLoggedInMiddleware(deps.Config.JwtSecret)
	isAdmin := middlewares.IsAdminMiddleware()

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.GET("/google/login", google.NewGoogleLoginHandler(deps.GoogleOauthConfig))
		authEndpoints.GET("/google/callback", google.NewGoogleCallbackHandler(deps.GoogleOauthConfig, deps.GoogleService, deps.UserService, deps.AuthService))
	}

	authorizedRoutes := router.Group("")
	authorizedRoutes.Use(isLoggedIn)

	recipesEndpoints := authorizedRoutes.Group("/recipes")
	{
		recipesEndpoints.GET("", recipes.NewGetRecipesHandler(deps.RecipeService))
		recipesEndpoints.POST("", isAdmin, recipes.NewCreateRecipeHandler(deps.RecipeService))
		recipesEndpoints.PUT("/:id", isAdmin, recipes.NewUpdateRecipeHandler(deps.RecipeService))
		recipesEndpoints.DELETE("/:id", isAdmin, recipes.NewDeleteRecipeHandler(deps.RecipeService))
	}

}
