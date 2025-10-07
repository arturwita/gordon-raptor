package recipes

import (
	"context"
	"errors"
	"time"

	"gordon-raptor/src/internal/consts"
	"gordon-raptor/src/internal/contracts"
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

func (repo *recipeRepository) GetRecipes(query *contracts.GetRecipesQueryDto, ctx context.Context) ([]*RecipeModel, error) {
	skip := int64((query.Page - 1) * query.Limit)
	limit := int64(query.Limit)
	filter := bson.M{
		"name": bson.M{
			"$regex":   query.Name,
			"$options": "i",
		},
	}
	cursor, err := repo.collection.Find(ctx, filter, options.Find().SetSkip(skip).SetLimit(limit))

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
	var updatedRecipe RecipeModel
	err := repo.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": db.EnsureMongoId(id)},
		bson.M{
			"$set": bson.M{
				"name":        dto.Name,
				"ingredients": dto.Ingredients,
				"updatedAt":   primitive.NewDateTimeFromTime(time.Now()),
			},
		},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedRecipe)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("recipe not found")
		}
		return nil, err
	}

	return &updatedRecipe, nil
}

func (repo *recipeRepository) DeleteRecipe(id string, ctx context.Context) error {
	result, err := repo.collection.DeleteOne(ctx, bson.M{"_id": db.EnsureMongoId(id)})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("recipe not found")
	}

	return nil
}
