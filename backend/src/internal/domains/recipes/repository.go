package recipes

import (
	"context"
	"time"

	"gordon-raptor/src/internal/contracts"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeRepository interface {
	CreateRecipe(dto contracts.CreateRecipeDto) (*RecipeModel, error)
}

type recipeRepository struct {
	collection *mongo.Collection
}

func NewRecipeRepository(database *mongo.Database) (RecipeRepository, error) {
	return &recipeRepository{database.Collection("recipes")}, nil
}

func (repo *recipeRepository) CreateRecipe(dto contracts.CreateRecipeDto) (*RecipeModel, error) {
	recipe := RecipeModel{
		Id:          primitive.NewObjectID(),
		Name:        dto.Name,
		Ingredients: dto.Ingredients,
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}
	_, err := repo.collection.InsertOne(context.Background(), recipe)

	if err != nil {
		return nil, err
	}

	return &recipe, nil
}
