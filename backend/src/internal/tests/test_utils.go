package recipes_e2e_tests

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func CleanDatabase(database *mongo.Database) {
	database.Drop(context.Background())
}
