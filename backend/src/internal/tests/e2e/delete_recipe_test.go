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

	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func TestDeleteRecipe(t *testing.T) {
	var method = "DELETE"
	var path = fmt.Sprintf("/recipes/%s", tests_mocks.MockRecipeId1)
	server, _ := app.NewApp(config.TestConfig)
	database, _ := db.NewMongoDatabase(config.TestConfig.MongoURL)

	recipesCollection := database.Collection(consts.CollectionNames["recipes"])
	usersCollection := database.Collection(consts.CollectionNames["users"])

	userBuilder := tests_utils.NewGenericEntityBuilder(usersCollection, tests_mocks.DefaultUserMock)
	recipesBuilder := tests_utils.NewGenericEntityBuilder(recipesCollection, tests_mocks.DefaultRecipeMock)

	t.Run("returns 404 when the recipe does not exist", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)
		mockAdmin := userBuilder.WithID(tests_mocks.MockUserId1).OverrideProps(map[string]any{"role": users.AdminRole}).Build()
		mockAdminJwt := tests_utils.GenerateTestJWT(mockAdmin)

		// when
		req, _ := http.NewRequest(method, path, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", mockAdminJwt))
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusNotFound, response.Code)

		var responseBody contracts.ErrorResponse
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, "recipe not found", responseBody.Message)
	})

	t.Run("returns 204 when the recipe is deleted", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)
		mockAdmin := userBuilder.WithID(tests_mocks.MockUserId1).OverrideProps(map[string]any{"role": users.AdminRole}).Build()
		mockAdminJwt := tests_utils.GenerateTestJWT(mockAdmin)

		// given
		recipesBuilder.WithID(tests_mocks.MockRecipeId1).Build()

		// when
		req, _ := http.NewRequest(method, path, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", mockAdminJwt))
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusNoContent, response.Code)
		assert.Equal(t, 0, len(response.Body.Bytes()))
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

	t.Run("returns 403 if non-admin tries to perform the request", func(t *testing.T) {
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
