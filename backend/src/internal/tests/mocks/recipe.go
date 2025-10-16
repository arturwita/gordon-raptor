package tests_mocks

import (
	"gordon-raptor/src/internal/domains/recipes"
	"gordon-raptor/src/pkg/db"
)

var MockRecipeId1 = "68dc4669766f5f4c66451161"
var MockRecipeId2 = "68dc4669766f5f4c66451162"
var MockRecipeId3 = "68dc4669766f5f4c66451163"
var MockRecipeId4 = "68dc4669766f5f4c66451164"
var MockRecipeId5 = "68dc4669766f5f4c66451165"

var DefaultRecipeMock = recipes.RecipeModel{
	Id:        db.EnsureMongoId(MockRecipeId1),
	CreatedAt: MockTimestamp,
	UpdatedAt: MockTimestamp,
	Name:      "spaghetti bolognese",
	Ingredients: map[string]string{
		"pasta": "100g",
		"meat":  "100g",
	},
}
