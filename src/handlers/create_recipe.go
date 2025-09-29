package handlers

import (
	"gordon-raptor/src/dtos"
	"gordon-raptor/src/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewCreateRecipe(repo repositories.RecipeRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		var recipeDto dtos.CreateRecipeDto
		if err := context.BindJSON(&recipeDto); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := repo.CreateRecipe(recipeDto); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusCreated, gin.H{
			"message": "created",
			"body":    recipeDto,
		})
	}
}
