package recipes

import "gordon-raptor/src/internal/contracts"

type RecipeService interface {
	CreateRecipe(dto contracts.CreateRecipeDto) (string, error)
}

type recipeService struct {
	repository RecipeRepository
}

func NewRecipeService(repository RecipeRepository) (RecipeService, error) {
	return &recipeService{repository}, nil
}

func (service *recipeService) CreateRecipe(dto contracts.CreateRecipeDto) (string, error) {
	service.repository.CreateRecipe(dto)

	return "success", nil
}
