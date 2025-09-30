package recipes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRecipeHandler(recipeService RecipeService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var dto CreateRecipeDto
		if err := context.BindJSON(&dto); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := recipeService.CreateRecipe(dto)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusCreated, CreateRecipeResponseDto{
			Result: result,
		})
	}
}
