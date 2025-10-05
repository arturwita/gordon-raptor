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
		var dto contracts.CreateRecipeDto
		if err := context.BindJSON(&dto); err != nil {
			context.JSON(http.StatusBadRequest, &contracts.ErrorResponse{Message: err.Error()})
			return
		}

		recipe, err := recipeService.CreateRecipe(dto, ctx)
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
		paginationDto := contracts.BindPagination(context)
		
		recipes, err := recipeService.GetRecipes(paginationDto, ctx)
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
