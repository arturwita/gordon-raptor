package repositories

import (
	"context"

	"gordon-raptor/src/config"
	"gordon-raptor/src/dtos"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type RecipeRepository interface {
	CreateRecipe(recipe dtos.CreateRecipeDto) error
}

type recipeRepository struct {
	collection *mongo.Collection
}

func RecipeRepositoryFactory(client *mongo.Client, cfg *config.Config) RecipeRepository {
	parsedDbUrl, err := connstring.ParseAndValidate(cfg.MongoURL)
	if err != nil {
		return nil
	}

	return &recipeRepository{client.Database(parsedDbUrl.Database).Collection("recipes")}
}

func (repo *recipeRepository) CreateRecipe(recipe dtos.CreateRecipeDto) error {
	_, err := repo.collection.InsertOne(context.Background(), recipe)
	return err
}
