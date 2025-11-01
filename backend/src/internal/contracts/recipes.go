package contracts

type RecipeDto struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Ingredients map[string]string `json:"ingredients"`
	Picture     *string           `json:"picture,omitempty"`
	Description *string           `json:"description,omitempty"`
	CreatedAt   string            `json:"createdAt"`
	UpdatedAt   string            `json:"updatedAt"`
}

type CreateRecipeBodyDto struct {
	Name        string            `form:"name" json:"name" binding:"required"`
	Ingredients map[string]string `form:"ingredients" json:"ingredients" binding:"required"`
	Picture     *string           `form:"picture" json:"picture" binding:"omitempty,url"`
	Description *string           `form:"description" json:"description" binding:"omitempty"`
}

type CreateRecipeResponseDto struct {
	Recipe *RecipeDto `json:"recipe"`
}

type GetRecipesQueryDto struct {
	Page  int    `form:"page" json:"page" binding:"gte=1"`
	Limit int    `form:"limit" json:"limit" binding:"gte=1,lte=100"`
	Name  string `form:"name" json:"name" binding:"omitempty"`
}

type GetRecipesResponseDto struct {
	Recipes []*RecipeDto       `json:"recipes"`
	Meta    *BasePaginationDto `json:"meta"`
}

type UpdateRecipeBodyDto struct {
	Name        string            `form:"name" json:"name" binding:"required"`
	Description *string           `form:"description" json:"description" binding:"omitempty"`
	Ingredients map[string]string `form:"ingredients" json:"ingredients" binding:"required"`
}

type UpdateRecipeResponseDto struct {
	Recipe *RecipeDto `json:"recipe"`
}

type RecipeIdParamDto struct {
	Id string `uri:"id" json:"id" binding:"required,len=24,hexadecimal"`
}
