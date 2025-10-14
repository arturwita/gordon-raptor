package di

import (
	"gordon-raptor/src/internal/config"
	"gordon-raptor/src/internal/domains/auth"
	"gordon-raptor/src/internal/domains/auth/google"
	"gordon-raptor/src/internal/domains/recipes"
	"gordon-raptor/src/internal/domains/users"
	"gordon-raptor/src/pkg/db"

	"golang.org/x/oauth2"
)

type DIContainer struct {
	RecipeRepository  recipes.RecipeRepository
	RecipeService     recipes.RecipeService
	UserRepository    users.UserRepository
	UserService       users.UserService
	GoogleService     google.GoogleService
	AuthService       auth.AuthService
	Config            *config.AppConfig
	GoogleOauthConfig *oauth2.Config
}

func NewDIContainer(cfg *config.AppConfig) (*DIContainer, error) {
	database, err := db.NewMongoDatabase(cfg.MongoURL)
	if err != nil {
		return nil, err
	}

	googleOauthConfig := config.NewGoogleOauthConfig(cfg)

	recipeRepository, err := recipes.NewRecipeRepository(database)
	if err != nil {
		return nil, err
	}

	recipeService, err := recipes.NewRecipeService(recipeRepository)
	if err != nil {
		return nil, err
	}

	userRepository, err := users.NewUserRepository(database)
	if err != nil {
		return nil, err
	}

	userService, err := users.NewUserService(userRepository)
	if err != nil {
		return nil, err
	}

	googleService, err := google.NewGoogleService(googleOauthConfig)
	if err != nil {
		return nil, err
	}

	authService, err := auth.NewAuthService(cfg)
	if err != nil {
		return nil, err
	}

	return &DIContainer{
		RecipeRepository:  recipeRepository,
		RecipeService:     recipeService,
		UserRepository:    userRepository,
		UserService:       userService,
		GoogleService:     googleService,
		AuthService:       authService,
		Config:            cfg,
		GoogleOauthConfig: googleOauthConfig,
	}, nil
}
