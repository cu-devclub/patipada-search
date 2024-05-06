package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"data-management/communication"
	mock "data-management/mock/communication"
)

func TestAuthMiddleware(t *testing.T) {
	mockGrpc := mock.NewMockgRPC()
	rabbitMQ := mock.MockRabbitMQ()
	comm := communication.NewCommunicationImpl(mockGrpc,rabbitMQ)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Set up the mock gRPC server
	g := &ginServer{
		comm: comm,
	}
	router.Use(g.AuthMiddleware("role"))

	// Define a test route
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Test route")
	})

	t.Run("Authorized Request", func(t *testing.T) {
		mock.SetAuthorizationResponse(true)
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "valid_token")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, "Test route", resp.Body.String())
	})

	t.Run("Unauthorized Request", func(t *testing.T) {
		mock.SetAuthorizationResponse(false)
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "invalid_token")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
	})
}
