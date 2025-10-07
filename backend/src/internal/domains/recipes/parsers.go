package recipes

import (
	"gordon-raptor/src/internal/contracts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParseGetRecipesQuery(context *gin.Context) *contracts.GetRecipesQueryDto {
	var params = &contracts.GetRecipesQueryDto{}
	_ = context.ShouldBindQuery(&params)

	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 || params.Limit > 100 {
		params.Limit = 20
	}

	return params
}

func ParseRecipeIdParam(context *gin.Context) string {
	var params = &contracts.RecipeIdParamDto{}
	if err := context.BindUri(&params); err != nil {
		context.JSON(http.StatusBadRequest, &contracts.ErrorResponse{Message: err.Error()})
		return ""
	}

	return params.Id
}

func ParseCreateRecipeBody(context *gin.Context) *contracts.CreateRecipeBodyDto {
	var body = &contracts.CreateRecipeBodyDto{}
	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, &contracts.ErrorResponse{Message: err.Error()})
		return nil
	}

	return body
}

func ParseUpdateRecipeBody(context *gin.Context) *contracts.UpdateRecipeBodyDto {
	var body = &contracts.UpdateRecipeBodyDto{}
	if err := context.BindJSON(&body); err != nil {
		context.JSON(http.StatusBadRequest, &contracts.ErrorResponse{Message: err.Error()})
		return nil
	}
	return body
}
