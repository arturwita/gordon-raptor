package users

import (
	"context"

	"gordon-raptor/src/internal/consts"
	"gordon-raptor/src/internal/custom_errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(dto *CreateUserDto, ctx context.Context) (*UserModel, error)
	GetUserByEmail(email string, ctx context.Context) (*UserModel, error)
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

func (repo *userRepository) GetUserByEmail(email string, ctx context.Context) (*UserModel, error) {
	filter := bson.M{"email": email}

	var user UserModel
	if err := repo.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, custom_errors.DomainErrors.User.NotFound
		}
		return nil, err
	}

	return &user, nil
}
