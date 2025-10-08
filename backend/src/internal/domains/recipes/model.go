package recipes

import "go.mongodb.org/mongo-driver/bson/primitive"

type RecipeModel struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Ingredients map[string]string  `bson:"ingredients" json:"ingredients"`
	Picture     *string            `bson:"picture,omitempty" json:"picture,omitempty"`
	CreatedAt   primitive.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt   primitive.DateTime `bson:"updatedAt" json:"updatedAt"`
}
