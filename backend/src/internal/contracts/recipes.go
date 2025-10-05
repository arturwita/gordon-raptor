package contracts

type RecipeDto struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Ingredients map[string]string `json:"ingredients"`
	CreatedAt   string            `json:"createdAt"`
	UpdatedAt   string            `json:"updatedAt"`
}

type CreateRecipeDto struct {
	Name        string            `form:"name" json:"name" binding:"required"`
	Ingredients map[string]string `form:"ingredients" json:"ingredients" binding:"required"`
}

type CreateRecipeResponseDto struct {
	Recipe *RecipeDto `json:"recipe"`
}

type GetRecipesResponseDto struct {
	Recipes []*RecipeDto `json:"recipes"`
}

type DeleteRecipeDto struct {
  Id string `uri:"id" binding:"required,len=24,hexadecimal"`
}