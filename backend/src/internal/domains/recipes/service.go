package recipes

import "gordon-raptor/src/internal/contracts"

type RecipeService interface {
	CreateRecipe(dto contracts.CreateRecipeDto) (*RecipeModel, error)
}

type recipeService struct {
	repository RecipeRepository
}

func NewRecipeService(repository RecipeRepository) (RecipeService, error) {
	return &recipeService{repository}, nil
}

func (service *recipeService) CreateRecipe(dto contracts.CreateRecipeDto) (*RecipeModel, error) {
	recipe, err := service.repository.CreateRecipe(dto)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}
