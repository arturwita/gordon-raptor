package tests_utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CleanTestDatabase(database *mongo.Database) {
	database.Drop(context.Background())
}

func EnsureMongoId(hex string) primitive.ObjectID {
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		panic(err)
	}

	return id
}
