package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	Ping(context)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.JSONEq(t, `{"message":"pong"}`, response.Body.String())
}
