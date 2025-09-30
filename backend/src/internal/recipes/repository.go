package recipes

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeRepository interface {
	CreateRecipe(dto CreateRecipeDto) (string, error)
}

type recipeRepository struct {
	collection *mongo.Collection
}

func NewRecipeRepository(database *mongo.Database) (RecipeRepository, error) {
	return &recipeRepository{database.Collection("recipes")}, nil
}

func (repo *recipeRepository) CreateRecipe(dto CreateRecipeDto) (string, error) {
	value, err := repo.collection.InsertOne(context.Background(), dto)
	if err != nil {
		return "failure", err
	}

	return "success", nil
}
