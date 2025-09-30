package tests_e2e

import (
	"testing"

	"gordon-raptor/src/internal/app"
	"gordon-raptor/src/internal/config"
	"gordon-raptor/src/internal/contracts"

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
		reqBody, _ := json.Marshal(contracts.CreateRecipeDto{
			Name: "spaghetti bolognese",
			Ingredients: map[string]string{
				"pasta": "100g",
				"meat":  "100g",
			},
		})

		// when
		req, _ := http.NewRequest(method, path, bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusCreated, response.Code)

		var responseBody contracts.CreateRecipeResponseDto
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)

		assert.NoError(t, err)
		assert.Equal(t, "success", responseBody.Recipe) // TODO: fix assertion
	})
}
