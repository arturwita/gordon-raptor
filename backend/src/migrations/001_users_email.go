package migrations

import (
	"context"
	"gordon-raptor/src/internal/consts"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUsersEmailUniqueIndex(db *mongo.Database) error {
	indexName := "unique_email"
	usersCollection := db.Collection(consts.CollectionNames["users"])

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true).SetName(indexName),
	}

	_, err := usersCollection.Indexes().CreateOne(context.Background(), indexModel)
	return err
}
