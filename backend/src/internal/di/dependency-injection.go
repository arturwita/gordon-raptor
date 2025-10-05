package di

import (
	"gordon-raptor/src/internal/config"
	"gordon-raptor/src/internal/domains/recipes"
	"gordon-raptor/src/pkg/db"
)

type DIContainer struct {
	RecipeRepository recipes.RecipeRepository
	RecipeService    recipes.RecipeService
	Config           *config.Config
}

func NewDIContainer(cfg *config.Config) (*DIContainer, error) {
	database, err := db.NewMongoDatabase(cfg.MongoURL)
	if err != nil {
		return nil, err
	}

	recipeRepository, err := recipes.NewRecipeRepository(database)
	if err != nil {
		return nil, err
	}

	recipeService, err := recipes.NewRecipeService(recipeRepository)
	if err != nil {
		return nil, err
	}

	return &DIContainer{
		RecipeRepository: recipeRepository,
		RecipeService:    recipeService,
		Config:           cfg,
	}, nil
}
