package tests_e2e

import (
	"fmt"
	"testing"
	"time"

	"gordon-raptor/src/internal/app"
	"gordon-raptor/src/internal/config"
	"gordon-raptor/src/internal/consts"
	"gordon-raptor/src/internal/contracts"
	tests_mocks "gordon-raptor/src/internal/tests/mocks"
	tests_utils "gordon-raptor/src/internal/tests/utils"
	"gordon-raptor/src/pkg/db"

	"github.com/stretchr/testify/assert"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func TestUpdateRecipe(t *testing.T) {
	var method = "PUT"
	var path = "/recipes"
	server, _ := app.NewApp(config.TestConfig)
	database, _ := db.NewMongoDatabase(config.TestConfig.MongoURL)
	collection := database.Collection(consts.CollectionNames["recipes"])

	recipesBuilder := tests_utils.NewGenericEntityBuilder(collection, tests_mocks.DefaultRecipeMock)

	t.Run("updates the recipe in the database and returns 200", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)

		// given
		recipesBuilder.WithID(tests_mocks.MockRecipeId1).OverrideProps(map[string]any{
			"name": "spaghetti bolognese",
			"ingredients": map[string]string{
				"pasta": "100g",
				"meat": "100g"},
			}).Build()

		expected := contracts.UpdateRecipeBodyDto{
			Name: "pizza",
			Ingredients: map[string]string{
				"mozarella": "200g",
				"tomatoes": "100g",
			},
		}
		reqBody, _ := json.Marshal(expected)

		// when
		req, _ := http.NewRequest(method, fmt.Sprintf("%s/%s", path, tests_mocks.MockRecipeId1), bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusOK, response.Code)

		var responseBody contracts.CreateRecipeResponseDto
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, tests_mocks.MockRecipeId1, responseBody.Recipe.Id)
		assert.Equal(t, expected.Name, responseBody.Recipe.Name)
		assert.Equal(t, expected.Ingredients, responseBody.Recipe.Ingredients)
		assert.Equal(t, tests_mocks.MockISOTimestamp, responseBody.Recipe.CreatedAt)

		createdTime, _ := time.Parse(time.RFC3339, responseBody.Recipe.CreatedAt)
		updatedTime, _ := time.Parse(time.RFC3339, responseBody.Recipe.UpdatedAt)
		assert.True(t, updatedTime.After(createdTime), "expected updatedAt to be after createdAt")
	})

	t.Run("returns 404 when the recipe does not exist", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)

		// given
		expected := contracts.UpdateRecipeBodyDto{
			Name: "pizza",
			Ingredients: map[string]string{
				"mozarella": "200g",
				"tomatoes": "100g",
			},
		}
		reqBody, _ := json.Marshal(expected)

		// when
		req, _ := http.NewRequest(method, fmt.Sprintf("%s/%s", path, tests_mocks.MockRecipeId1), bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusNotFound, response.Code)

		var responseBody contracts.ErrorResponse
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, "recipe not found", responseBody.Message)
	})
}
