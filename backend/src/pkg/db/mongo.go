package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

func NewMongoClient(uri string) (*mongo.Client, error) {
	return mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
}

func NewMongoDatabase(uri string) (*mongo.Database, error) {
	client, err := NewMongoClient(uri)
	if err != nil {
		fmt.Println("Error creating MongoDB client:", err, uri)
		return nil, err
	}

	parsedDbUrl, err := connstring.ParseAndValidate(uri)
	if err != nil {
		fmt.Println("Error parsing MongoDB URL:", err, uri)
		return nil, err
	}

	return client.Database(parsedDbUrl.Database), nil
}
