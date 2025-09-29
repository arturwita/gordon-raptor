package di

import (
	"gordon-raptor/src/pkg/config"
	"gordon-raptor/src/pkg/db"
	"gordon-raptor/src/repositories"
)

type DIContainer struct {
	RecipeRepository repositories.RecipeRepository
	Config           *config.Config
}

func NewDIContainer(cfg *config.Config) (*DIContainer, error) {
	client, err := db.NewMongoClient(cfg.MongoURL)
	if err != nil {
		return nil, err
	}

	recipeRepository, err := repositories.NewRecipeRepository(client, cfg)
	if err != nil {
		return nil, err
	}

	return &DIContainer{
		RecipeRepository: recipeRepository,
		Config:           cfg,
	}, nil
}
