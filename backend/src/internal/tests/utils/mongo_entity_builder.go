package tests_utils

import (
	"context"
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
	builder.entity["_id"] = EnsureMongoId(id)
	return builder
}

func (builder *GenericEntityBuilder[T]) OverrideProps(props map[string]any) *GenericEntityBuilder[T] {
	maps.Copy(builder.entity, props)
	return builder
}

func (builder *GenericEntityBuilder[T]) Build() (*T, error) {
	_, err := builder.collection.InsertOne(context.Background(), builder.entity)
	if err != nil {
		return nil, err
	}

	resultBson := builder.entity
	var result T
	bsonBytes, _ := bson.Marshal(resultBson)
	if err := bson.Unmarshal(bsonBytes, &result); err != nil {
		return nil, err
	}

	builder.entity = make(bson.M)
	maps.Copy(builder.entity, builder.defaultValues)

	return &result, nil
}
