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
	fmt.Println("ENTERING BUILD FUNCTION", builder.entity)
	entity, err := builder.collection.InsertOne(context.Background(), builder.entity)
	fmt.Println("INSERTED ENTITY", entity.InsertedID)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to save entity with id: '%s'", entity.InsertedID), err)
		return nil
	}

	fmt.Println("MARSHALLING")
	resultBson := builder.entity
	var result T
	bsonBytes, _ := bson.Marshal(resultBson)
	fmt.Println("UNMARSHALLING")
	if err := bson.Unmarshal(bsonBytes, &result); err != nil {
		fmt.Println(fmt.Sprintf("Failed unmarshall entity with id: '%s'", entity.InsertedID), err)
		return nil
	}

	fmt.Println("MAKING JSON")
	builder.entity = make(bson.M)
	fmt.Println("COPYING")
	maps.Copy(builder.entity, builder.defaultValues)

	return &result
}
