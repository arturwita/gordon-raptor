package tests_e2e

import (
	"testing"

	"gordon-raptor/src/internal/app"
	"gordon-raptor/src/internal/config"
	"gordon-raptor/src/internal/contracts"
	tests_utils "gordon-raptor/src/internal/tests/utils"
	"gordon-raptor/src/pkg/db"

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
	database, _ := db.NewMongoDatabase(config.TestConfig.MongoURL)

	t.Run("saves the recipe in the database and returns 201", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)

		// given
		expected := contracts.CreateRecipeBodyDto{
			Name: "spaghetti bolognese",
			Ingredients: map[string]string{
				"pasta": "100g",
				"meat":  "100g",
			},
		}
		reqBody, _ := json.Marshal(expected)

		// when
		req, _ := http.NewRequest(method, path, bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", config.TestConfig.AdminApiKey)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusCreated, response.Code)

		var responseBody contracts.CreateRecipeResponseDto
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, expected.Name, responseBody.Recipe.Name)
		assert.Equal(t, expected.Ingredients, responseBody.Recipe.Ingredients)
		assert.NotEmpty(t, responseBody.Recipe.Id)
		assert.NotEmpty(t, responseBody.Recipe.CreatedAt)
		assert.NotEmpty(t, responseBody.Recipe.UpdatedAt)
	})

	t.Run("returns 403 if x-api-key header is missing", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)

		// when
		req, _ := http.NewRequest(method, path, nil)
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusForbidden, response.Code)

		var responseBody contracts.ErrorResponse
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, "You're not allowed to perform this action", responseBody.Message)
	})


	t.Run("returns 403 if x-api-key header has invalid value", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)

		// when
		req, _ := http.NewRequest(method, path, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", "invalid-api-key")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusForbidden, response.Code)

		var responseBody contracts.ErrorResponse
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, "You're not allowed to perform this action", responseBody.Message)
	})
}
