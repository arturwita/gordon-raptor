package tests_unit

import (
	"gordon-raptor/src/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPaginationMeta(t *testing.T) {
	t.Run("GetPaginationMeta", func(t *testing.T) {

		t.Run("calculates pagination metadata correctly (middle page)", func(t *testing.T) {
			params := &utils.GetBasePaginationMetaParams{
				TotalItems: 50,
				Page:       2,
				Limit:      10,
			}

			meta := utils.GetBasePaginationMeta(params)

			assert.Equal(t, 50, meta.TotalItems)
			assert.Equal(t, 2, meta.Page)
			assert.Equal(t, 10, meta.Limit)
			assert.Equal(t, 5, meta.TotalPages)
			assert.True(t, meta.HasNextPage)
			assert.True(t, meta.HasPrevPage)
			assert.Equal(t, 1, *meta.PrevPage)
			assert.Equal(t, 3, *meta.NextPage)
		})

		t.Run("handles first page correctly", func(t *testing.T) {
			params := &utils.GetBasePaginationMetaParams{
				TotalItems: 35,
				Page:       1,
				Limit:      10,
			}

			meta := utils.GetBasePaginationMeta(params)

			assert.Equal(t, 4, meta.TotalPages)
			assert.True(t, meta.HasNextPage)
			assert.False(t, meta.HasPrevPage)
			assert.Nil(t, meta.PrevPage)
			assert.Equal(t, 2, *meta.NextPage)
		})

		t.Run("handles last page correctly", func(t *testing.T) {
			params := &utils.GetBasePaginationMetaParams{
				TotalItems: 30,
				Page:       3,
				Limit:      10,
			}

			meta := utils.GetBasePaginationMeta(params)

			assert.Equal(t, 3, meta.TotalPages)
			assert.False(t, meta.HasNextPage)
			assert.True(t, meta.HasPrevPage)
			assert.Equal(t, 2, *meta.PrevPage)
			assert.Nil(t, meta.NextPage)
		})

		t.Run("handles case when Total < Limit", func(t *testing.T) {
			params := &utils.GetBasePaginationMetaParams{
				TotalItems: 5,
				Page:       1,
				Limit:      10,
			}

			meta := utils.GetBasePaginationMeta(params)

			assert.Equal(t, 1, meta.TotalPages)
			assert.False(t, meta.HasPrevPage)
			assert.False(t, meta.HasNextPage)
			assert.Nil(t, meta.PrevPage)
			assert.Nil(t, meta.NextPage)
		})

		t.Run("handles Page beyond total pages gracefully", func(t *testing.T) {
			params := &utils.GetBasePaginationMetaParams{
				TotalItems: 10,
				Page:       5,
				Limit:      2,
			}

			meta := utils.GetBasePaginationMeta(params)

			assert.Equal(t, 5, meta.TotalPages)
			assert.False(t, meta.HasNextPage)
			assert.True(t, meta.HasPrevPage)
			assert.Equal(t, 4, *meta.PrevPage)
			assert.Nil(t, meta.NextPage)
		})

	})
}
