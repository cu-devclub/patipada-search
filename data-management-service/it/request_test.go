package tests

import (
	"data-management/config"
	"data-management/logging"
	mock_communication "data-management/mock/communication"
	test_container_database "data-management/mock/testcontainer/database"
	requestHandler "data-management/request/handlers"
	"data-management/request/models"
	requestRepository "data-management/request/repositories"
	requestUsecase "data-management/request/usecases"
	validator "data-management/structValidator"
	"data-management/testutil"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setUpRequestTestEnvironment() (requestHandler.Handlers, func()) {
	logging.NewSLogger()

	err := config.LoadConfig("../.")
	if err != nil {
		panic(err)
	}

	db, dbCleanUp, err := test_container_database.NewMockMongoClient()
	if err != nil {
		panic(err)
	}

	v := validator.NewValidator()

	comm := mock_communication.MockCommunication()

	// request Construct
	requestRepositories := requestRepository.NewRepositories(db, &comm)

	requestUsecase := requestUsecase.NewRequestUsecase(&requestRepositories, &v)

	requestHandlers := requestHandler.NewRequestHandler(&requestUsecase)

	return requestHandlers, dbCleanUp
}

func setupRequestGinEngine(handler requestHandler.Handlers) *gin.Engine {
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

func TestRequests(t *testing.T) {
	requestHandlers, dbCleanUp := setUpRequestTestEnvironment()
	defer dbCleanUp()

	g := setupRequestGinEngine(requestHandlers)

	t.Run("Test Success GetRequests", func(t *testing.T) {
		// requestHandlers.GetRequests()
		req := testutil.CreateNewRequest("GET", "/requests", nil)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("Success Get Lastest Request Of Record : 200 OK", func(t *testing.T) {
		mock_communication.SetSearchResponse(true)
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

	t.Run("Success Get Summary : 200 OK", func(t *testing.T) {
		req := testutil.CreateNewRequest("GET", "/summary", nil)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

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

	t.Run("Success Update Request : 200 OK", func(t *testing.T) {
		requestModel := models.Request{}
		requestModel.MockData()
		req := testutil.CreateNewRequest("POST", "/requests", requestModel)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal([]byte(w.Body.Bytes()), &response)
		if err != nil {
			t.Error("Error Unmarshalling Response Body : ", err)
		}

		requestID := response["requestID"].(string)
		id := response["id"].(string)

		requestModel.ID = id
		requestModel.RequestID = requestID
		requestModel.Status = "reviewed"
		requestModel.ApprovedBy = "admin"
		reqq := testutil.CreateNewRequest("PUT", "/request", requestModel)
		ww := httptest.NewRecorder()

		g.ServeHTTP(ww, reqq)

		assert.Equal(t, 200, ww.Code)
	})

	t.Run("Fail Update Request : 400 Bad Request : Body can't bind", func(t *testing.T) {
		req := testutil.CreateNewRequest("PUT", "/request", "invalid body")
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("Insert-Update-Sync", func(t *testing.T) {
		requestModel := models.Request{}
		requestModel.MockData()
		req := testutil.CreateNewRequest("POST", "/requests", requestModel)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal([]byte(w.Body.Bytes()), &response)
		if err != nil {
			t.Error("Error Unmarshalling Response Body : ", err)
		}

		requestID := response["requestID"].(string)
		id := response["id"].(string)

		requestModel.ID = id
		requestModel.RequestID = requestID
		requestModel.Status = "reviewed"
		requestModel.ApprovedBy = "admin"
		reqq := testutil.CreateNewRequest("PUT", "/request", requestModel)
		ww := httptest.NewRecorder()

		g.ServeHTTP(ww, reqq)

		assert.Equal(t, 200, ww.Code)

		body := models.SyncRequestRecord{}
		body.RequestId = requestID

		reqqq := testutil.CreateNewRequest("POST", "/sync", body)
		www := httptest.NewRecorder()

		g.ServeHTTP(www, reqqq)

		assert.Equal(t, 200, www.Code)
	})

	t.Run("Fail Sync Request Record : 400 Bad Request : Body can't bind", func(t *testing.T) {
		req := testutil.CreateNewRequest("POST", "/sync", "invalid body")
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("Success Sync All Request Records : 200 OK", func(t *testing.T) {
		req := testutil.CreateNewRequest("POST", "/sync-all", nil)
		w := httptest.NewRecorder()

		g.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}
