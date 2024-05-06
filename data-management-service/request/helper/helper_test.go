package helper

import (
	"data-management/constant"
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
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			By:         "User1",
			ApprovedBy: "Admin",
		},
		{
			ID:         "2",
			RequestID:  "REQ2",
			Index:      "1",
			YoutubeURL: "https://youtube.com/video2",
			Question:   "Question 2",
			Answer:     "Answer 2",
			StartTime:  "start",
			EndTime:    "end",
			Status:     constant.REQUEST_STATUS_PENDING,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			By:         "User2",
			ApprovedBy: "Admin",
		},
	}

	// Create source request
	sourceRequest := &models.Request{
		ID:         "3",
		RequestID:  "REQ3",
		Index:      "1",
		YoutubeURL: "https://youtube.com/video3",
		Question:   "Question 3",
		Answer:     "Answer 3",
		StartTime:  "start",
		EndTime:    "end",
		Status:     constant.REQUEST_STATUS_PENDING,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		By:         "User3",
		ApprovedBy: "Admin",
	}

	// Call the function
	neededUpdateRequest := UpdatePreviousRequestsStatus(requests, sourceRequest)

	// Assert the results
	assert.Len(t, neededUpdateRequest, 2)
	assert.Equal(t, constant.REQUEST_STATUS_REVIEWED, neededUpdateRequest[0].Status)
	assert.Equal(t, constant.REQUEST_STATUS_REVIEWED, neededUpdateRequest[1].Status)
}
func TestGenerateRequestID(t *testing.T) {
	nextSeq := 1
	expectedRequestID := "REQ1"

	requestID := GenerateRequestID(nextSeq)

	assert.Equal(t, expectedRequestID, requestID)
}