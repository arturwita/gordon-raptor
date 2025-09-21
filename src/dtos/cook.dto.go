package dtos

type CookDto struct {
	Recipe string `form:"recipe" json:"recipe" binding:"required"`
}
