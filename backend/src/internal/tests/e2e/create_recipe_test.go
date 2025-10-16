package tests_e2e

import (
	"fmt"
	"testing"

	"gordon-raptor/src/internal/app"
	"gordon-raptor/src/internal/config"
	"gordon-raptor/src/internal/consts"
	"gordon-raptor/src/internal/contracts"
	"gordon-raptor/src/internal/domains/users"
	tests_mocks "gordon-raptor/src/internal/tests/mocks"
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

	usersCollection := database.Collection(consts.CollectionNames["users"])
	userBuilder := tests_utils.NewGenericEntityBuilder(usersCollection, tests_mocks.DefaultUserMock)

	t.Run("saves the recipe in the database and returns 201", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)
		mockAdmin := userBuilder.WithID(tests_mocks.MockUserId1).OverrideProps(map[string]any{"role": users.AdminRole}).Build()
		mockAdminJwt := tests_utils.GenerateTestJWT(mockAdmin)

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
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", mockAdminJwt))
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

	t.Run("returns 401 if auth header is missing", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)

		// when
		req, _ := http.NewRequest(method, path, nil)
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusUnauthorized, response.Code)

		var responseBody contracts.ErrorResponse
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, "Unauthorized", responseBody.Message)
	})

	t.Run("returns 401 if auth header has invalid value", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)

		// when
		req, _ := http.NewRequest(method, path, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer INVALID_TOKEN")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusUnauthorized, response.Code)

		var responseBody contracts.ErrorResponse
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, "Unauthorized", responseBody.Message)
	})

	t.Run("returns 403 if a non-admin tries to perform the request", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)
		mockUser := userBuilder.WithID(tests_mocks.MockUserId2).OverrideProps(map[string]any{"role": users.UserRole}).Build()
		mockUserJwt := tests_utils.GenerateTestJWT(mockUser)

		// when
		req, _ := http.NewRequest(method, path, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", mockUserJwt))
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
