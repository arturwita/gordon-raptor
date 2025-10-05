package contracts

import "github.com/gin-gonic/gin"

type PaginationDto struct {
	Page  int `form:"page" json:"page" binding:"gte=1"`
	Limit int `form:"limit" json:"limit" binding:"gte=1,lte=100"`
}

func BindPagination(c *gin.Context) *PaginationDto {
	var params = &PaginationDto{}
	_ = c.ShouldBindQuery(&params)

	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 || params.Limit > 100 {
		params.Limit = 20
	}

	return params
}
