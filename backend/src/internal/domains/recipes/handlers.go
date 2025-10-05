package recipes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gordon-raptor/src/internal/contracts"
)

func CreateRecipeHandler(recipeService RecipeService) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := context.Request.Context()

		var body contracts.CreateRecipeBodyDto
		if err := context.BindJSON(&body); err != nil {
			context.JSON(http.StatusBadRequest, &contracts.ErrorResponse{Message: err.Error()})
			return
		}

		recipe, err := recipeService.CreateRecipe(body, ctx)
		if err != nil {
			fmt.Println("Failed to create recipe", err)
			context.JSON(http.StatusInternalServerError, &contracts.ErrorResponse{Message: err.Error()})
			return
		}

		context.JSON(http.StatusCreated, &contracts.CreateRecipeResponseDto{
			Recipe: MapToRecipeDto(recipe),
		})
	}
}

func GetRecipesHandler(recipeService RecipeService) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := context.Request.Context()
		query := contracts.BindPagination(context)

		recipes, err := recipeService.GetRecipes(query, ctx)
		if err != nil {
			fmt.Println("Failed to get recipes", err)
			context.JSON(http.StatusInternalServerError, &contracts.ErrorResponse{Message: err.Error()})
			return
		}

		context.JSON(http.StatusOK, &contracts.GetRecipesResponseDto{
			Recipes: MapRecipesToDtos(recipes),
		})
	}
}

func UpdateRecipeHandler(recipeService RecipeService) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := context.Request.Context()

		var params contracts.RecipeIdParamDto
		if err := context.BindUri(&params); err != nil {
			context.JSON(http.StatusBadRequest, &contracts.ErrorResponse{Message: err.Error()})
			return
		}

		var body contracts.UpdateRecipeBodyDto
		if err := context.BindJSON(&body); err != nil {
			context.JSON(http.StatusBadRequest, &contracts.ErrorResponse{Message: err.Error()})
			return
		}

		recipe, err := recipeService.UpdateRecipe(params.Id, body, ctx)
		if err != nil {
			context.JSON(http.StatusNotFound, &contracts.ErrorResponse{Message: err.Error()})
			return
		}

		context.JSON(http.StatusOK, &contracts.UpdateRecipeResponseDto{
			Recipe: MapToRecipeDto(recipe),
		})
	}
}

func DeleteRecipeHandler(recipeService RecipeService) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := context.Request.Context()

		var params contracts.RecipeIdParamDto
		if err := context.BindUri(&params); err != nil {
			context.JSON(http.StatusBadRequest, &contracts.ErrorResponse{Message: err.Error()})
			return
		}

		err := recipeService.DeleteRecipe(params.Id, ctx)
		if err != nil {
			context.JSON(http.StatusNotFound, &contracts.ErrorResponse{Message: err.Error()})
			return
		}

		context.JSON(http.StatusNoContent, nil)
	}
}
