package recipes

import (
	"context"
	"errors"
	"time"

	"gordon-raptor/src/internal/consts"
	"gordon-raptor/src/internal/contracts"
	"gordon-raptor/src/internal/custom_errors"
	"gordon-raptor/src/pkg/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RecipeRepository interface {
	CreateRecipe(dto *contracts.CreateRecipeBodyDto, ctx context.Context) (*RecipeModel, error)
	GetRecipes(query *contracts.GetRecipesQueryDto, ctx context.Context) ([]*RecipeModel, int, error)
	UpdateRecipe(id string, dto *contracts.UpdateRecipeBodyDto, ctx context.Context) (*RecipeModel, error)
	DeleteRecipe(id string, ctx context.Context) error
}

type recipeRepository struct {
	collection *mongo.Collection
}

func NewRecipeRepository(database *mongo.Database) (RecipeRepository, error) {
	return &recipeRepository{database.Collection(consts.CollectionNames["recipes"])}, nil
}

func (repo *recipeRepository) CreateRecipe(dto *contracts.CreateRecipeBodyDto, ctx context.Context) (*RecipeModel, error) {
	recipe := NewRecipe(dto)

	if _, err := repo.collection.InsertOne(ctx, recipe); err != nil {
		return nil, err
	}

	return recipe, nil
}

func (repo *recipeRepository) GetRecipes(query *contracts.GetRecipesQueryDto, ctx context.Context) ([]*RecipeModel, int, error) {
	skip := int64((query.Page - 1) * query.Limit)
	limit := int64(query.Limit)
	filter := bson.M{
		"name": bson.M{
			"$regex":   query.Name,
			"$options": "i",
		},
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{
			Key: "$facet",
			Value: bson.M{
				"items": mongo.Pipeline{
					{{Key: "$skip", Value: skip}},
					{{Key: "$limit", Value: limit}},
				},
				"total": mongo.Pipeline{
					{{Key: "$count", Value: "count"}},
				},
			},
		}},
	}

	cursor, err := repo.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var result []struct {
		Items []*RecipeModel `bson:"items"`
		Total []struct {
			Count int `bson:"count"`
		} `bson:"total"`
	}

	if err := cursor.All(ctx, &result); err != nil {
		return nil, 0, err
	}

	total := 0
	if len(result[0].Total) > 0 {
		total = result[0].Total[0].Count
	}

	return result[0].Items, total, nil
}

func (repo *recipeRepository) UpdateRecipe(id string, dto *contracts.UpdateRecipeBodyDto, ctx context.Context) (*RecipeModel, error) {
	filter := bson.M{"_id": db.EnsureMongoId(id)}
	update := bson.M{
		"$set": bson.M{
			"name":        dto.Name,
			"ingredients": dto.Ingredients,
			"description": dto.Description,
			"updatedAt":   primitive.NewDateTimeFromTime(time.Now()),
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedRecipe RecipeModel
	err := repo.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedRecipe)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, custom_errors.DomainErrors.Recipe.NotFound
		}
		return nil, err
	}

	return &updatedRecipe, nil
}

func (repo *recipeRepository) DeleteRecipe(id string, ctx context.Context) error {
	filter := bson.M{"_id": db.EnsureMongoId(id)}

	result, err := repo.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return custom_errors.DomainErrors.Recipe.NotFound
	}

	return nil
}
