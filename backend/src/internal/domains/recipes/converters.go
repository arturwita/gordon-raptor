package recipes

import (
	"time"

	"gordon-raptor/src/internal/contracts"
)

func MapToRecipeDto(recipe *RecipeModel) *contracts.RecipeDto {
	return &contracts.RecipeDto{
		Id:          recipe.Id.Hex(),
		Name:        recipe.Name,
		Ingredients: recipe.Ingredients,
		Picture:     recipe.Picture,
		CreatedAt:   recipe.CreatedAt.Time().Format(time.RFC3339),
		UpdatedAt:   recipe.UpdatedAt.Time().Format(time.RFC3339),
	}
}

func MapRecipesToDtos(recipes []*RecipeModel) []*contracts.RecipeDto {
	dtos := make([]*contracts.RecipeDto, 0, len(recipes))
	for _, recipe := range recipes {
		dtos = append(dtos, MapToRecipeDto(recipe))
	}
	return dtos
}
