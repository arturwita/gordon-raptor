package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email     string             `bson:"email" json:"email"`
	FirstName *string            `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName  *string            `bson:"lastName,omitempty" json:"lastName,omitempty"`
	Picture   *string            `bson:"picture,omitempty" json:"picture,omitempty"`
	Role      Role               `bson:"role" json:"role"`
	CreatedAt primitive.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt primitive.DateTime `bson:"updatedAt" json:"updatedAt"`
}

func NewUser(dto *CreateUserDto) *UserModel {
	now := primitive.NewDateTimeFromTime(time.Now())

	return &UserModel{
		Id:        primitive.NewObjectID(),
		Email:     dto.Email,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Picture:   dto.Picture,
		Role:      UserRole,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
