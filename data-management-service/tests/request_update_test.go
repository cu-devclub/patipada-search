package tests

import (
	"data-management/request/models"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUpdateRequest(t *testing.T) {
	requestHandlers := setUpTestEnvironment()
	g := gin.Default()
	g.PUT("/request", requestHandlers.UpdateRequest)
	g.POST("/requests", requestHandlers.InsertRequest)
	t.Run("Success Update Request : 200 OK", func(t *testing.T) {
		// Insert data
		m := models.Request{}
		m.MockData()
		req := createNewRequestFromDataType("POST", "/requests", m)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Code)

		// Get the request ID
		var response map[string]interface{}
		err := json.Unmarshal([]byte(w.Body.Bytes()), &response)
		if err != nil {
			t.Error("Error Unmarshalling Response Body : ", err)
		}

		data := response["data"]
		id := data.(map[string]interface{})["id"].(string)
		requestID := data.(map[string]interface{})["request_id"].(string)

		// Update data
		m.ID = id
		m.RequestID = requestID
		m.Question = "Question Updated"
		m.Status = "approved"
		m.ApprovedBy = "super-admin"
		req = createNewRequestFromDataType("PUT", "/request", m)
		w = httptest.NewRecorder()

		g.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})
}
