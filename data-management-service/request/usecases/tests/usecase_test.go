package request_usecase_test

import (
	mock_request "data-management/mock/request"
	"data-management/request/entities"
	"data-management/request/models"
	"data-management/request/usecases"
	validator "data-management/structValidator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupUsecase() usecases.UseCase {
	mockRepo := mock_request.NewMockRepositories()
	validate := validator.NewValidator()
	// Create the usecase
	usecase := usecases.NewRequestUsecase(&mockRepo, &validate)
	return usecase
}

func TestGetRequest(t *testing.T) {
	usecase := setupUsecase()
	// Test the GetRequest function
	t.Run("GetRequest Success", func(t *testing.T) {
		mock_request.SetValidateRecordResponse(true)
		expectedRequest := mock_request.GetMockRequestConstant()
		mock_request.SetMockRequestsValue([]*entities.Request{expectedRequest})

		requests, err := usecase.GetRequest("pending", "username", "REQ1", "sadads-1", "approvedBy")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(requests))
	})

	t.Run("GetRequest Success multiple request", func(t *testing.T) {
		mock_request.SetValidateRecordResponse(true)
		expectedRequest1 := mock_request.GetMockRequestConstant()
		expectedRequest2 := mock_request.GetMockRequestConstant()
		expectedRequest2.RequestID = "REQ2"

		mock_request.SetMockRequestsValue([]*entities.Request{expectedRequest1, expectedRequest2})

		requests, err := usecase.GetRequest("pending", "username", "", "sadads-1", "approvedBy")
		assert.Nil(t, err)
		assert.Equal(t, 2, len(requests))
	})

	t.Run("GetRequest Fail Invalid status", func(t *testing.T) {
		requests, err := usecase.GetRequest("invalid", "username", "REQ1", "sadads-1", "approvedBy")
		assert.NotNil(t, err)
		assert.Nil(t, requests)
	})
}

func TestGetLatestRequestOfRecord(t *testing.T) {
	usecase := setupUsecase()

	t.Run("GetLatestRequest Success", func(t *testing.T) {
		expectedRequest := mock_request.GetMockRequestConstant()
		mock_request.SetMockRequestsValue([]*entities.Request{expectedRequest})
		mock_request.SetValidateRecordResponse(true)

		request, err := usecase.GetLastestRequestOfRecord("sadads-1")
		assert.Nil(t, err)
		assert.Equal(t, expectedRequest.RequestID, request.RequestID)
	})

	t.Run("GetLatestRequest Success Multiple value", func(t *testing.T) {
		expectedRequest1 := mock_request.GetMockRequestConstant()
		expectedRequest2 := mock_request.GetMockRequestConstant()
		expectedRequest2.RequestID = "REQ2"
		expectedRequest2.UpdatedAt = expectedRequest2.UpdatedAt.Add(500)

		mock_request.SetMockRequestsValue([]*entities.Request{expectedRequest1, expectedRequest2})
		mock_request.SetValidateRecordResponse(true)

		request, err := usecase.GetLastestRequestOfRecord("sadads-1")
		assert.Nil(t, err)
		assert.Equal(t, expectedRequest2.RequestID, request.RequestID)
	})

	t.Run("GetLatestRequest Fail Invalid Record", func(t *testing.T) {
		mock_request.SetValidateRecordResponse(false)
		request, err := usecase.GetLastestRequestOfRecord("")
		assert.NotNil(t, err)
		assert.Nil(t, request)
	})
}

func TestInsertRequest(t *testing.T) {
	usecase := setupUsecase()
	// Test the InsertRequest function
	t.Run("InsertRequest Success", func(t *testing.T) {
		request := models.Request{}
		request.MockData()

		mock_request.SetValidateRecordResponse(true)

		err := usecase.InsertRequest(&request)
		assert.Nil(t, err)
	})

	t.Run("InsertRequest Fail Invalid Request", func(t *testing.T) {
		request := models.Request{}
		request.MockData()
		request.Index = ""

		err := usecase.InsertRequest(&request)
		assert.NotNil(t, err)
	})

	t.Run("InsertRequest Fail Invalid Status", func(t *testing.T) {
		request := models.Request{}
		request.MockData()
		request.Status = "invalid"

		err := usecase.InsertRequest(&request)
		assert.NotNil(t, err)
	})
}

func TestUpdateRequest(t *testing.T) {
	usecase := setupUsecase()
	// Test the UpdateRequest function
	t.Run("UpdateRequest Success", func(t *testing.T) {
		request := models.Request{}
		request.MockData()
		request.Status = "reviewed"
		request.ApprovedBy = "username"

		err := usecase.UpdateRequest(&request)
		assert.Nil(t, err)
	})

	t.Run("UpdateRequest Fail Invalid Request", func(t *testing.T) {
		request := models.Request{}
		request.MockData()
		request.Index = ""

		err := usecase.UpdateRequest(&request)
		assert.NotNil(t, err)
	})

	t.Run("UpdateRequest Fail Invalid Status", func(t *testing.T) {
		request := models.Request{}
		request.MockData()
		request.Status = "invalid"

		err := usecase.UpdateRequest(&request)
		assert.NotNil(t, err)
	})
}

func TestSummaryRequest(t *testing.T) {
	usecase := setupUsecase()
	// Test the SummaryRequest function
	t.Run("SummaryRequest Success", func(t *testing.T) {
		// mock the requests
		expectedRequest1 := mock_request.GetMockRequestConstant()
		expectedRequest2 := mock_request.GetMockRequestConstant()
		expectedRequest2.RequestID = "REQ2"
		expectedRequest2.Status = "reviewed"

		mock_request.SetMockRequestsValue([]*entities.Request{expectedRequest1, expectedRequest2})

		// mock record counter
		mock_request.SetRecordCounterResponse(&entities.RecordCounter{
			RecordAmount:      12,
			YoutubeClipAmount: 10,
		})

		summary, err := usecase.SummaryData()

		assert.Nil(t, err)
		assert.Equal(t, 2, summary.RequestSummary.RequestAmount)
		assert.Equal(t, 1, summary.RequestSummary.ReviewedAmount)
		assert.Equal(t, 1, summary.RequestSummary.PendingAmount)
		assert.Equal(t, 12, summary.RecordSummary.RecordAmount)
		assert.Equal(t, 10, summary.RecordSummary.YouTubeClipAmount)
	})
}

func TestSyncRequestRecord(t *testing.T) {
	usecase := setupUsecase()
	// Test the SyncRequestRecord function
	t.Run("SyncRequestRecord Success", func(t *testing.T) {
		syncRequestRecord := models.SyncRequestRecord{}
		syncRequestRecord.RequestId = "REQ1"

		expectedRecord := mock_request.GetMockRequestConstant()
		expectedRecord.RequestID = "REQ1"
		expectedRecord.Status = "reviewed"

		mock_request.SetMockRequestsValue([]*entities.Request{expectedRecord})

		err := usecase.SyncRequestRecord(&syncRequestRecord)
		assert.Nil(t, err)
	})

	t.Run("SyncRequestRecord Fail Invalid Request (no request)", func(t *testing.T) {
		syncRequestRecord := models.SyncRequestRecord{}
		syncRequestRecord.RequestId = "REQ3"
		mock_request.SetMockRequestsValue([]*entities.Request{})

		err := usecase.SyncRequestRecord(&syncRequestRecord)
		assert.NotNil(t, err)
	})

	t.Run("SyncRequestRecord Fail Invalid Request (Status)", func(t *testing.T) {
		syncRequestRecord := models.SyncRequestRecord{}
		syncRequestRecord.RequestId = "REQ1"

		expectedRecord := mock_request.GetMockRequestConstant()
		expectedRecord.RequestID = "REQ1"
		expectedRecord.Status = "pending"

		mock_request.SetMockRequestsValue([]*entities.Request{expectedRecord})

		err := usecase.SyncRequestRecord(&syncRequestRecord)
		assert.NotNil(t, err)
	})
}

func TestSyncAllRequestRecords(t *testing.T) {
	usecase := setupUsecase()
	// Test the SyncAllRequestRecords function
	t.Run("SyncAllRequestRecords Success", func(t *testing.T) {
		err := usecase.SyncAllRequestRecords()
		assert.Nil(t, err)
	})
}
