package recipes

import "go.mongodb.org/mongo-driver/bson/primitive"

type RecipeModel struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" binding:"required"`
	Name        string             `bson:"name" json:"name" binding:"required"`
	Ingredients map[string]string  `bson:"ingredients" json:"ingredients" binding:"required"`
	CreatedAt   primitive.DateTime `bson:"createdAt" json:"createdAt" binding:"required"`
	UpdatedAt   primitive.DateTime `bson:"updatedAt" json:"updatedAt" binding:"required"`
}
