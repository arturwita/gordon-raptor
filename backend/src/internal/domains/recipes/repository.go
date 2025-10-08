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
	GetRecipes(query *contracts.GetRecipesQueryDto, ctx context.Context) ([]*RecipeModel, error)
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
	recipe := RecipeModel{
		Name:        dto.Name,
		Ingredients: dto.Ingredients,
		Picture:     dto.Picture,
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}
	result, err := repo.collection.InsertOne(ctx, recipe)

	if err != nil {
		return nil, err
	}

	if recipeId, isOk := result.InsertedID.(primitive.ObjectID); isOk {
		recipe.Id = recipeId
	}

	return &recipe, nil
}

func (repo *recipeRepository) GetRecipes(query *contracts.GetRecipesQueryDto, ctx context.Context) ([]*RecipeModel, error) {
	skip := int64((query.Page - 1) * query.Limit)
	limit := int64(query.Limit)
	filter := bson.M{
		"name": bson.M{
			"$regex":   query.Name,
			"$options": "i",
		},
	}
	opts := options.Find().SetSkip(skip).SetLimit(limit)

	cursor, err := repo.collection.Find(ctx, filter, opts)
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

func (repo *recipeRepository) UpdateRecipe(id string, dto *contracts.UpdateRecipeBodyDto, ctx context.Context) (*RecipeModel, error) {
	filter := bson.M{"_id": db.EnsureMongoId(id)}
	update := bson.M{
		"$set": bson.M{
			"name":        dto.Name,
			"ingredients": dto.Ingredients,
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
