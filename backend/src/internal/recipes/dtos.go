package recipes

type CreateRecipeDto struct {
	Recipe string `form:"recipe" json:"recipe" binding:"required"`
}

type CreateRecipeResponseDto struct {
	Result string 
}
