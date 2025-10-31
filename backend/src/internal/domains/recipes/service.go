package recipes

import (
	"context"
	"gordon-raptor/src/internal/contracts"
)

type RecipeService interface {
	CreateRecipe(dto *contracts.CreateRecipeBodyDto, ctx context.Context) (*RecipeModel, error)
	GetRecipes(query *contracts.GetRecipesQueryDto, ctx context.Context) ([]*RecipeModel, int, error)
	UpdateRecipe(id string, dto *contracts.UpdateRecipeBodyDto, ctx context.Context) (*RecipeModel, error)
	DeleteRecipe(id string, ctx context.Context) error
}

type recipeService struct {
	repository RecipeRepository
}

func NewRecipeService(repository RecipeRepository) (RecipeService, error) {
	return &recipeService{repository}, nil
}

func (service *recipeService) CreateRecipe(dto *contracts.CreateRecipeBodyDto, ctx context.Context) (*RecipeModel, error) {
	return service.repository.CreateRecipe(dto, ctx)
}

func (service *recipeService) GetRecipes(query *contracts.GetRecipesQueryDto, ctx context.Context) ([]*RecipeModel, int, error) {
	return service.repository.GetRecipes(query, ctx)
}

func (service *recipeService) UpdateRecipe(id string, dto *contracts.UpdateRecipeBodyDto, ctx context.Context) (*RecipeModel, error) {
	return service.repository.UpdateRecipe(id, dto, ctx)
}

func (service *recipeService) DeleteRecipe(id string, ctx context.Context) error {
	return service.repository.DeleteRecipe(id, ctx)
}
