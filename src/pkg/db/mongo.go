package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoClientFactory(uri string) (*mongo.Client, error) {
	return mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
}
