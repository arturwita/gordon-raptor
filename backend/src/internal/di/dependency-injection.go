package di

import (
	"gordon-raptor/src/internal/config"
	"gordon-raptor/src/internal/recipes"
	"gordon-raptor/src/pkg/db"
)

type DIContainer struct {
	RecipeRepository recipes.RecipeRepository
	Config           *config.Config
}

func NewDIContainer(cfg *config.Config) (*DIContainer, error) {
	client, err := db.NewMongoClient(cfg.MongoURL)
	if err != nil {
		return nil, err
	}

	recipeRepository, err := recipes.NewRecipeRepository(client, cfg)
	if err != nil {
		return nil, err
	}

	return &DIContainer{
		RecipeRepository: recipeRepository,
		Config:           cfg,
	}, nil
}
