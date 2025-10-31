package tests_e2e

import (
	"fmt"
	"testing"

	"gordon-raptor/src/internal/app"
	"gordon-raptor/src/internal/config"
	"gordon-raptor/src/internal/consts"
	"gordon-raptor/src/internal/contracts"
	tests_mocks "gordon-raptor/src/internal/tests/mocks"
	tests_utils "gordon-raptor/src/internal/tests/utils"
	"gordon-raptor/src/pkg/db"

	"github.com/stretchr/testify/assert"

	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func TestGetRecipes(t *testing.T) {
	var method = "GET"
	var path = "/recipes"
	server, _ := app.NewApp(config.TestConfig)
	database, _ := db.NewMongoDatabase(config.TestConfig.MongoURL)

	recipesCollection := database.Collection(consts.CollectionNames["recipes"])
	usersCollection := database.Collection(consts.CollectionNames["users"])

	recipesBuilder := tests_utils.NewGenericEntityBuilder(recipesCollection, tests_mocks.DefaultRecipeMock)
	userBuilder := tests_utils.NewGenericEntityBuilder(usersCollection, tests_mocks.DefaultUserMock)

	t.Run("returns empty list with status 200 when there are no recipes", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)
		mockUser := userBuilder.WithID(tests_mocks.MockUserId1).Build()
		mockJwt := tests_utils.GenerateTestJWT(mockUser)

		// when
		req, _ := http.NewRequest(method, path, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", mockJwt))
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusOK, response.Code)

		var responseBody contracts.GetRecipesResponseDto
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, 0, len(responseBody.Recipes))
		assert.NotEmpty(t, responseBody.Meta)
	})

	t.Run("returns the recipes with status 200", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)
		mockUser := userBuilder.WithID(tests_mocks.MockUserId1).Build()
		mockJwt := tests_utils.GenerateTestJWT(mockUser)

		// given
		recipesBuilder.WithID(tests_mocks.MockRecipeId1).OverrideProps(map[string]any{"name": "spaghetti"}).Build()
		recipesBuilder.WithID(tests_mocks.MockRecipeId2).OverrideProps(map[string]any{"name": "pizza"}).Build()

		// when
		req, _ := http.NewRequest(method, path, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", mockJwt))
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusOK, response.Code)

		var responseBody contracts.GetRecipesResponseDto
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, 2, len(responseBody.Recipes))
		assert.Equal(t, "spaghetti", responseBody.Recipes[0].Name)
		assert.Equal(t, "pizza", responseBody.Recipes[1].Name)
		assert.NotEmpty(t, responseBody.Meta)
	})

	t.Run("properly handles page/limit parameters", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)
		mockUser := userBuilder.WithID(tests_mocks.MockUserId1).Build()
		mockJwt := tests_utils.GenerateTestJWT(mockUser)

		// given
		page := 2
		limit := 1
		recipesBuilder.WithID(tests_mocks.MockRecipeId1).OverrideProps(map[string]any{"name": "spaghetti"}).Build()
		recipesBuilder.WithID(tests_mocks.MockRecipeId2).OverrideProps(map[string]any{"name": "pizza"}).Build()

		// when
		req, _ := http.NewRequest(method, fmt.Sprintf("%s?page=%d&limit=%d", path, page, limit), nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", mockJwt))
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusOK, response.Code)

		var responseBody contracts.GetRecipesResponseDto
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, 1, len(responseBody.Recipes))
		assert.Equal(t, "pizza", responseBody.Recipes[0].Name)

		assert.Equal(t, 2, responseBody.Meta.Page)
		assert.Equal(t, 1, responseBody.Meta.Limit)
		assert.Equal(t, 2, responseBody.Meta.TotalItems)
		assert.Equal(t, 2, responseBody.Meta.TotalPages)
		assert.False(t, responseBody.Meta.HasNextPage)
		assert.True(t, responseBody.Meta.HasPrevPage)
		assert.Nil(t, responseBody.Meta.NextPage)
		assert.Equal(t, 1, int(*responseBody.Meta.PrevPage))
	})

	t.Run("returns only recipes containing the given search string", func(t *testing.T) {
		tests_utils.CleanTestDatabase(database)
		mockUser := userBuilder.WithID(tests_mocks.MockUserId1).Build()
		mockJwt := tests_utils.GenerateTestJWT(mockUser)

		// given
		name := "piz"
		recipesBuilder.WithID(tests_mocks.MockRecipeId1).OverrideProps(map[string]any{"name": "spaghetti"}).Build()
		recipesBuilder.WithID(tests_mocks.MockRecipeId2).OverrideProps(map[string]any{"name": "pizza"}).Build()

		// when
		req, _ := http.NewRequest(method, fmt.Sprintf("%s?name=%s", path, name), nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", mockJwt))
		response := httptest.NewRecorder()
		server.ServeHTTP(response, req)

		// then
		assert.Equal(t, http.StatusOK, response.Code)

		var responseBody contracts.GetRecipesResponseDto
		err := json.Unmarshal(response.Body.Bytes(), &responseBody)
		assert.NoError(t, err)

		assert.Equal(t, 1, len(responseBody.Recipes))
		assert.Equal(t, "pizza", responseBody.Recipes[0].Name)
		assert.NotEmpty(t, responseBody.Meta)
	})

	t.Run("returns 401 when auth header is missing", func(t *testing.T) {
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

	t.Run("returns 401 when auth header has invalid value", func(t *testing.T) {
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
}
