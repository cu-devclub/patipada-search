package helper

import (
	"data-management/constant"
	"data-management/request/entities"
	"data-management/request/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUpdatePreviousRequestsStatus(t *testing.T) {
	// Create sample requests
	requests := []*models.Request{
		{
			ID:         "1",
			RequestID:  "REQ1",
			Index:      "1",
			YoutubeURL: "https://youtube.com/video1",
			Question:   "Question 1",
			Answer:     "Answer 1",
			StartTime:  "start",
			EndTime:    "end",
			Status:     constant.REQUEST_STATUS_PENDING,
			CreatedAt:  time.Now().Add(-time.Hour),
			UpdatedAt:  time.Now().Add(-time.Hour),
			By:         "User1",
			ApprovedBy: "Admin",
		},
		{
			ID:         "2",
			RequestID:  "REQ2",
			Index:      "2",
			YoutubeURL: "https://youtube.com/video2",
			Question:   "Question 2",
			Answer:     "Answer 2",
			StartTime:  "start",
			EndTime:    "end",
			Status:     constant.REQUEST_STATUS_PENDING,
			CreatedAt:  time.Now().Add(-2 * time.Hour),
			UpdatedAt:  time.Now().Add(-2 * time.Hour),
			By:         "User2",
			ApprovedBy: "Admin",
		},
		{
			ID:         "3",
			RequestID:  "REQ3",
			Index:      "1",
			YoutubeURL: "https://youtube.com/video3",
			Question:   "Question 3",
			Answer:     "Answer 3",
			StartTime:  "start",
			EndTime:    "end",
			Status:     constant.REQUEST_STATUS_REVIEWED,
			CreatedAt:  time.Now().Add(-3 * time.Hour),
			UpdatedAt:  time.Now().Add(-3 * time.Hour),
			By:         "User3",
			ApprovedBy: "Admin",
		},
	}

	// Create source request
	sourceRequest := &models.Request{
		ID:         "4",
		RequestID:  "REQ4",
		Index:      "1",
		YoutubeURL: "https://youtube.com/video4",
		Question:   "Question 4",
		Answer:     "Answer 4",
		StartTime:  "start",
		EndTime:    "end",
		Status:     constant.REQUEST_STATUS_PENDING,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		By:         "User4",
		ApprovedBy: "Admin",
	}

	// Call the function
	neededUpdateRequest := UpdatePreviousRequestsStatus(requests, sourceRequest)

	// Assert the results
	assert.Len(t, neededUpdateRequest, 1)
	assert.Equal(t, constant.REQUEST_STATUS_REVIEWED, neededUpdateRequest[0].Status)
}
func TestGenerateRequestID(t *testing.T) {
	nextSeq := 1
	expectedRequestID := "REQ1"

	requestID := GenerateRequestID(nextSeq)

	assert.Equal(t, expectedRequestID, requestID)
}

func TestEntityToModels(t *testing.T) {
	// Create sample entity
	entity := &entities.Request{
		ID:         "1",
		RequestID:  "REQ1",
		Index:      "1",
		YoutubeURL: "https://youtube.com/video1",
		Question:   "Question 1",
		Answer:     "Answer 1",
		StartTime:  "start",
		EndTime:    "end",
		Status:     constant.REQUEST_STATUS_PENDING,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		By:         "User1",
		ApprovedBy: "Admin",
	}

	// Call the function
	model := EntityToModels(entity)

	// Assert the results
	assert.Equal(t, entity.ID, model.ID)
	assert.Equal(t, entity.RequestID, model.RequestID)
	assert.Equal(t, entity.Index, model.Index)
	assert.Equal(t, entity.YoutubeURL, model.YoutubeURL)
	assert.Equal(t, entity.Question, model.Question)
	assert.Equal(t, entity.Answer, model.Answer)
	assert.Equal(t, entity.StartTime, model.StartTime)
	assert.Equal(t, entity.EndTime, model.EndTime)
	assert.Equal(t, entity.Status, model.Status)
	assert.Equal(t, entity.CreatedAt, model.CreatedAt)
	assert.Equal(t, entity.UpdatedAt, model.UpdatedAt)
	assert.Equal(t, entity.By, model.By)
	assert.Equal(t, entity.ApprovedBy, model.ApprovedBy)
}

func TestRequestToRecordsEntity(t *testing.T) {
	// Create sample request
	request := &entities.Request{
		Index:      "1",
		YoutubeURL: "https://youtube.com/video1",
		Question:   "Question 1",
		Answer:     "Answer 1",
		StartTime:  "start",
		EndTime:    "end",
	}

	// Call the function
	record := RequestToRecordsEntity(request)

	// Assert the results
	assert.Equal(t, request.Index, record.Index)
	assert.Equal(t, request.YoutubeURL, record.YoutubeURL)
	assert.Equal(t, request.Question, record.Question)
	assert.Equal(t, request.Answer, record.Answer)
	assert.Equal(t, request.StartTime, record.StartTime)
	assert.Equal(t, request.EndTime, record.EndTime)
}
