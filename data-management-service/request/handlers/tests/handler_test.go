package request_handler_tests

import (
	mock_request "data-management/mock/request"
	"data-management/request/handlers"
	"data-management/request/models"
	"data-management/testutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setUpRequestHandler() *gin.Engine {
	ratingUsecase := mock_request.NewMockUsecase()
	handler := handlers.NewRequestHandler(&ratingUsecase)

	g := gin.Default()
	g.POST("/requests", handler.InsertRequest)
	g.GET("/requests", handler.GetRequest)
	g.GET("/request/latestRecord", handler.GetLastestRequestOfRecord)
	g.PUT("/request", handler.UpdateRequest)
	g.GET("/summary", handler.GetSummary)
	g.POST("/sync", handler.SyncRequestRecord)
	g.POST("/sync-all", handler.SyncAllRequestRecords)

	return g
}

func TestGetRequest(t *testing.T) {
	g := setUpRequestHandler()

	t.Run("Success Get Request : 200 OK", func(t *testing.T) {
		req := testutil.CreateNewRequest("GET", "/requests", nil)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("Success Get Request with RequestID filter : 200 OK", func(t *testing.T) {
		req := testutil.CreateNewRequest("GET", "/requests?requestID=mock", nil)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}

func TestGetLastestRequestOfRecord(t *testing.T) {
	g := setUpRequestHandler()

	t.Run("Success Get Lastest Request Of Record : 200 OK", func(t *testing.T) {
		req := testutil.CreateNewRequest("GET", "/request/latestRecord?index=demo", nil)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("Fail Get Lastest Request Of Record : 400 Bad Request : Missing index", func(t *testing.T) {
		req := testutil.CreateNewRequest("GET", "/request/latestRecord", nil)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
}

func TestGetSummary(t *testing.T) {
	g := setUpRequestHandler()

	t.Run("Success Get Summary : 200 OK", func(t *testing.T) {
		req := testutil.CreateNewRequest("GET", "/summary", nil)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}

func TestInsertRequest(t *testing.T) {
	g := setUpRequestHandler()

	t.Run("Success Insert Request : 201 Created", func(t *testing.T) {
		requestModel := models.Request{}
		requestModel.MockData()

		req := testutil.CreateNewRequest("POST", "/requests", requestModel)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Code)
	})

	t.Run("Fail Insert Request : 400 Bad Request : Body can't bind", func(t *testing.T) {
		req := testutil.CreateNewRequest("POST", "/requests", "invalid body")
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
}

func TestUpdateRequest(t *testing.T) {
	g := setUpRequestHandler()

	t.Run("Success Update Request : 200 OK", func(t *testing.T) {
		requestModel := models.Request{}
		requestModel.MockData()

		req := testutil.CreateNewRequest("PUT", "/request", requestModel)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("Fail Update Request : 400 Bad Request : Body can't bind", func(t *testing.T) {
		req := testutil.CreateNewRequest("PUT", "/request", "invalid body")
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
}

func TestSyncRequestRecord(t *testing.T) {
	g := setUpRequestHandler()

	t.Run("Success Sync Request Record : 200 OK", func(t *testing.T) {
		body := models.SyncRequestRecord{}
		body.RequestId = "mock"

		req := testutil.CreateNewRequest("POST", "/sync", body)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("Fail Sync Request Record : 400 Bad Request : Body can't bind", func(t *testing.T) {
		req := testutil.CreateNewRequest("POST", "/sync", "invalid body")
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
}

func TestSyncAllRequestRecords(t *testing.T) {
	g := setUpRequestHandler()

	t.Run("Success Sync All Request Records : 200 OK", func(t *testing.T) {
		req := testutil.CreateNewRequest("POST", "/sync-all", nil)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}
