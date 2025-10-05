package recipes

import (
	"context"
	"gordon-raptor/src/internal/contracts"
)

type RecipeService interface {
	CreateRecipe(dto contracts.CreateRecipeDto, ctx context.Context) (*RecipeModel, error)
	GetRecipes(paginationDto *contracts.PaginationDto, ctx context.Context) ([]*RecipeModel, error)
}

type recipeService struct {
	repository RecipeRepository
}

func NewRecipeService(repository RecipeRepository) (RecipeService, error) {
	return &recipeService{repository}, nil
}

func (service *recipeService) CreateRecipe(dto contracts.CreateRecipeDto, ctx context.Context) (*RecipeModel, error) {
	recipe, err := service.repository.CreateRecipe(dto, ctx)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (service *recipeService) GetRecipes(paginationDto *contracts.PaginationDto, ctx context.Context) ([]*RecipeModel, error) {
	recipes, err := service.repository.GetRecipes(paginationDto, ctx)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}
