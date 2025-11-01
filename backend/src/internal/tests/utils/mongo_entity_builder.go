package tests_utils

import (
	"context"
	"fmt"
	"gordon-raptor/src/pkg/db"
	"maps"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GenericEntityBuilder[T any] struct {
	collection    *mongo.Collection
	entity        bson.M
	defaultValues map[string]any
}

func NewGenericEntityBuilder[T any](collection *mongo.Collection, defaultValues T) *GenericEntityBuilder[T] {
	entity := make(bson.M)
	bsonBytes, _ := bson.Marshal(defaultValues)
	bson.Unmarshal(bsonBytes, &entity)

	return &GenericEntityBuilder[T]{
		collection:    collection,
		entity:        entity,
		defaultValues: entity,
	}
}

func (builder *GenericEntityBuilder[T]) WithID(id string) *GenericEntityBuilder[T] {
	builder.entity["_id"] = db.EnsureMongoId(id)
	return builder
}

func (builder *GenericEntityBuilder[T]) OverrideProps(props map[string]any) *GenericEntityBuilder[T] {
	maps.Copy(builder.entity, props)
	return builder
}

func (builder *GenericEntityBuilder[T]) Build() *T {
	entity, err := builder.collection.InsertOne(context.Background(), builder.entity)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to save entity with id: '%s'", entity.InsertedID), err)
		return nil
	}

	var result T
	resultBson := builder.entity
	bsonBytes, _ := bson.Marshal(resultBson)

	if err := bson.Unmarshal(bsonBytes, &result); err != nil {
		fmt.Println(fmt.Sprintf("Failed unmarshall entity with id: '%s'", entity.InsertedID), err)
		return nil
	}

	builder.entity = make(bson.M)
	maps.Copy(builder.entity, builder.defaultValues)

	return &result
}
