package recipes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gordon-raptor/src/internal/contracts"
)

func CreateRecipeHandler(recipeService RecipeService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var dto contracts.CreateRecipeDto
		if err := context.BindJSON(&dto); err != nil {
			context.JSON(http.StatusBadRequest, contracts.ErrorResponse{Message: err.Error()})
			return
		}

		result, err := recipeService.CreateRecipe(dto)
		if err != nil {
			context.JSON(http.StatusInternalServerError, contracts.ErrorResponse{Message: err.Error()})
			return
		}

		context.JSON(http.StatusCreated, contracts.CreateRecipeResponseDto{
			Result: result,
		})
	}
}
