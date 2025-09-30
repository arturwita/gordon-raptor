package recipes_e2e_tests

import (
	"testing"

	"gordon-raptor/src/internal/app"
	"gordon-raptor/src/internal/config"
	"gordon-raptor/src/internal/recipes"

	"github.com/stretchr/testify/assert"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func TestCreateRecipe(t *testing.T) {
	var method = "POST"
	var path = "/recipes"
	server, _ := app.NewApp(config.TestConfig)

	t.Run("saves the recipe in the database and returns 201", func(t *testing.T) {
		// given
		reqBody, _ := json.Marshal(recipes.CreateRecipeDto{
			Recipe: "pasta",
		})

		// when
		req, _ := http.NewRequest(method, path, bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusCreated, response.Code)

		var responseBody recipes.CreateRecipeResponseDto
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)

		assert.NoError(t, err)
		assert.Equal(t, "success", responseBody.Result)
	})
}
