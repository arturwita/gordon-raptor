package recipes

import (
	"context"
	"time"

	"gordon-raptor/src/internal/contracts"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RecipeRepository interface {
	CreateRecipe(dto contracts.CreateRecipeDto, ctx context.Context) (*RecipeModel, error)
	GetRecipes(paginationDto *contracts.PaginationDto, ctx context.Context) ([]*RecipeModel, error)
}

type recipeRepository struct {
	collection *mongo.Collection
}

func NewRecipeRepository(database *mongo.Database) (RecipeRepository, error) {
	return &recipeRepository{database.Collection("recipes")}, nil
}

func (repo *recipeRepository) CreateRecipe(dto contracts.CreateRecipeDto, ctx context.Context) (*RecipeModel, error) {
	recipe := RecipeModel{
		Id:          primitive.NewObjectID(),
		Name:        dto.Name,
		Ingredients: dto.Ingredients,
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}
	_, err := repo.collection.InsertOne(ctx, recipe)

	if err != nil {
		return nil, err
	}

	return &recipe, nil
}

func (repo *recipeRepository) GetRecipes(paginationDto *contracts.PaginationDto, ctx context.Context) ([]*RecipeModel, error) {
	skip := int64((paginationDto.Page - 1) * paginationDto.Limit)
	limit := int64(paginationDto.Limit)
	cursor, err := repo.collection.Find(ctx, bson.M{}, options.Find().SetSkip(skip).SetLimit(limit))

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var recipes []*RecipeModel
	for cursor.Next(ctx) {
		var recipe RecipeModel
		if err := cursor.Decode(&recipe); err != nil {
			return nil, err
		}
		recipes = append(recipes, &recipe)
	}

	return recipes, nil
}
