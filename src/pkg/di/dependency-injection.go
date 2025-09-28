package di

import (
	"gordon-raptor/src/config"
	"gordon-raptor/src/pkg/db"
	"gordon-raptor/src/repositories"
)

type DIContainer struct {
	RecipeRepository repositories.RecipeRepository
}

func DIContainerFactory(cfg *config.Config) (*DIContainer, error) {
	client, err := db.MongoClientFactory(cfg.MongoURL)
	if err != nil {
		return nil, err
	}

	recipeRepository := repositories.RecipeRepositoryFactory(client, cfg)

	return &DIContainer{
		RecipeRepository: recipeRepository,
	}, nil
}
