package tests

import (
	"data-management/request/models"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInsertRequest(t *testing.T) {
	requestHandlers := setUpTestEnvironment()
	g := gin.Default()
	g.POST("/requests", requestHandlers.InsertRequest)

	t.Run("Success Insert Request : 201 Created", func(t *testing.T) {
		m := models.Request{}
		jsonBody := m.CreateMockJSON()

		req := createNewRequestFromRawJSON("POST", "/requests", jsonBody)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Code)
	})

	t.Run("Fail Insert Request : 400 Bad Request : Body can't bind", func(t *testing.T) {
		m := models.Request{}
		m.MockData()
		m.Status ="invalid status"

		req := createNewRequestFromDataType("POST", "/requests", m)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("Fail Insert Request : 400 Bad Request : Index not exists", func(t *testing.T) {
		//TODO : Implementing after implementing index validation
	})
}
