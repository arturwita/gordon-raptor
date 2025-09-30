package tests_utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func CleanDatabase(database *mongo.Database) {
	database.Drop(context.Background())
}
