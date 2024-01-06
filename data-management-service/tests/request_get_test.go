package tests

import (
	"data-management/request/models"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllRequests(t *testing.T) {
	requestHandlers := setUpTestEnvironment()
	// cfg := config.GetConfig()
	g := gin.Default()
	g.GET("/requests", requestHandlers.GetAllRequests)

	t.Run("Success Get All Request : 200 OK", func(t *testing.T) {
		req := createNewRequestFromRawJSON("GET", "/requests", "")
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}

func TestGetRequestRequest(t *testing.T) {
	requestHandlers := setUpTestEnvironment()
	g := gin.Default()
	g.GET("/requests/:requestID", requestHandlers.GetRequestByRequestID)
	g.POST("/requests", requestHandlers.InsertRequest)
	t.Run("Success Get Request : 200 OK", func(t *testing.T) {
		//* Insert first
		m := models.Request{}
		jsonBody := m.CreateMockJSON()

		req := createNewRequestFromRawJSON("POST", "/requests", jsonBody)
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
		requestID := data.(map[string]interface{})["request_id"].(string)

		// making GET request
		req = createNewRequestFromRawJSON("GET", fmt.Sprintf("/requests/%s", requestID), "")
		w = httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		// check the result is the same as the inserted request
		var response2 map[string]interface{}
		err = json.Unmarshal([]byte(w.Body.Bytes()), &response2)
		if err != nil {
			t.Error("Error Unmarshalling Response Body : ", err)
		}
		data2 := response["data"]
		requestID2 := data2.(map[string]interface{})["request_id"]
		assert.Equal(t, requestID, requestID2)
	})

	t.Run("Fail Get Request : 404 Not Found : Request ID not exists", func(t *testing.T) {
		req := createNewRequestFromRawJSON("GET", "/requests/123", "")
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
	})
}
