package tests_unit

import (
	"gordon-raptor/src/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomString(t *testing.T) {
	t.Run("GenerateRandomString", func(t *testing.T) {
		t.Run("generates a random string of the given length", func(t *testing.T) {
			// given
			length := 5

			// when
			str1 := utils.GenerateRandomString(length)
			str2 := utils.GenerateRandomString(length)

			// then
			assert.Equal(t, len(str1), length)
			assert.NotEqual(t, str1, "")
			assert.NotEqual(t, str1, str2)
		})

		t.Run("returns an empty string when n is negative", func(t *testing.T) {
			// given
			length := -1

			// when
			str := utils.GenerateRandomString(length)

			// then
			assert.Equal(t, str, "")
		})
	})
}
