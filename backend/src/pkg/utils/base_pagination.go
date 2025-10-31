package utils

import (
	"gordon-raptor/src/internal/contracts"
	"math"
)

type GetBasePaginationMetaParams struct {
	Page       int
	Limit      int
	TotalItems int
}

func GetBasePaginationMeta(params *GetBasePaginationMetaParams) *contracts.BasePaginationDto {
	totalPages := int(math.Ceil(float64(params.TotalItems) / float64(params.Limit)))
	hasNextPage := params.Page >= 1 && params.Page < totalPages
	hasPrevPage := params.Page > 1 && params.Page <= totalPages+1

	var prevPage *int
	var nextPage *int

	if hasPrevPage {
		prevPage = ToIntPointer(params.Page - 1)
	}

	if hasNextPage {
		nextPage = ToIntPointer(params.Page + 1)
	}

	return &contracts.BasePaginationDto{
		TotalItems:  params.TotalItems,
		Page:        params.Page,
		Limit:       params.Limit,
		TotalPages:  totalPages,
		HasNextPage: hasNextPage,
		HasPrevPage: hasPrevPage,
		NextPage:    nextPage,
		PrevPage:    prevPage,
	}
}
