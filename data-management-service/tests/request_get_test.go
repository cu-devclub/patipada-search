package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetRequest(t *testing.T) {
	requestHandlers := setUpTestEnvironment()
	g := gin.Default()
	g.GET("/requests", requestHandlers.GetRequest)
	t.Run("Success Get Reques ; No filter ; all result : 200 OK", func(t *testing.T) {
		req := createNewRequestFromDataType("GET", "/requests", nil)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}
