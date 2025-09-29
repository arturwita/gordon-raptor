package dtos

type CreateRecipeDto struct {
	Recipe string `form:"recipe" json:"recipe" binding:"required"`
}
