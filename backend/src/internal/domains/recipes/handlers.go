package recipes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gordon-raptor/src/internal/contracts"
)

func NewCreateRecipeHandler(recipeService RecipeService) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := context.Request.Context()
		body := ParseCreateRecipeBody(context)

		recipe, err := recipeService.CreateRecipe(body, ctx)
		if err != nil {
			context.JSON(http.StatusInternalServerError, &contracts.ErrorResponse{Message: err.Error()})
			return
		}

		context.JSON(http.StatusCreated, &contracts.CreateRecipeResponseDto{
			Recipe: MapToRecipeDto(recipe),
		})
	}
}

func NewGetRecipesHandler(recipeService RecipeService) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := context.Request.Context()
		query := ParseGetRecipesQuery(context)

		recipes, err := recipeService.GetRecipes(query, ctx)
		if err != nil {
			context.JSON(http.StatusInternalServerError, &contracts.ErrorResponse{Message: err.Error()})
			return
		}

		context.JSON(http.StatusOK, &contracts.GetRecipesResponseDto{
			Recipes: MapRecipesToDtos(recipes),
		})
	}
}

func NewUpdateRecipeHandler(recipeService RecipeService) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := context.Request.Context()
		recipeId := ParseRecipeIdParam(context)
		body := ParseUpdateRecipeBody(context)

		recipe, err := recipeService.UpdateRecipe(recipeId, body, ctx)
		if err != nil {
			context.JSON(http.StatusNotFound, &contracts.ErrorResponse{Message: err.Error()})
			return
		}

		context.JSON(http.StatusOK, &contracts.UpdateRecipeResponseDto{
			Recipe: MapToRecipeDto(recipe),
		})
	}
}

func NewDeleteRecipeHandler(recipeService RecipeService) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := context.Request.Context()
		recipeId := ParseRecipeIdParam(context)

		err := recipeService.DeleteRecipe(recipeId, ctx)
		if err != nil {
			context.JSON(http.StatusNotFound, &contracts.ErrorResponse{Message: err.Error()})
			return
		}

		context.JSON(http.StatusNoContent, nil)
	}
}
