package users

import (
	"context"

	"gordon-raptor/src/internal/consts"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(dto *CreateUserDto, ctx context.Context) (*UserModel, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) (UserRepository, error) {
	return &userRepository{database.Collection(consts.CollectionNames["users"])}, nil
}

func (repo *userRepository) CreateUser(dto *CreateUserDto, ctx context.Context) (*UserModel, error) {
	user := NewUser(dto)

	if _, err := repo.collection.InsertOne(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
