package helper

import (
	"data-management/constant"
	"data-management/request/entities"
	"data-management/request/models"
	"fmt"
	"time"
)

// This will generate a unique RequestID for each new record, like "REQ1", "REQ2", etc.
// using IncrementRecordCounter() from repositories
func GenerateRequestID(nextSeq int) string {
	nextRequestID := fmt.Sprintf("REQ%d", nextSeq)
	return nextRequestID
}

func ModelsToEntity(m *models.Request) (e *entities.Request) {
	e = &entities.Request{
		ID:         m.ID,
		RequestID:  m.RequestID,
		Index:      m.Index,
		YoutubeURL: m.YoutubeURL,
		Question:   m.Question,
		Answer:     m.Answer,
		StartTime:  m.StartTime,
		EndTime:    m.EndTime,
		Status:     m.Status,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		By:         m.By,
		ApprovedBy: m.ApprovedBy,
	}
	return e
}

func EntityToModels(e *entities.Request) (m *models.Request) {
	m = &models.Request{
		ID:         e.ID,
		RequestID:  e.RequestID,
		Index:      e.Index,
		YoutubeURL: e.YoutubeURL,
		Question:   e.Question,
		Answer:     e.Answer,
		StartTime:  e.StartTime,
		EndTime:    e.EndTime,
		Status:     e.Status,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
		By:         e.By,
		ApprovedBy: e.ApprovedBy,
	}
	return m
}

func RequestToRecordsEntity(r *entities.Request) *entities.Record {
	return &entities.Record{
		Index:      r.Index,
		YoutubeURL: r.YoutubeURL,
		Question:   r.Question,
		Answer:     r.Answer,
		StartTime:  r.StartTime,
		EndTime:    r.EndTime,
	}
}

func UpdatePreviousRequestsStatus(requests []*models.Request, sourceRequest *models.Request) []*entities.Request {
	neededUpdateRequest := make([]*entities.Request, 0)
	// ---  Update all request that come before the current request by setting status to "reviewed"
	for _, req := range requests {
		if req.UpdatedAt.After(sourceRequest.UpdatedAt) {
			continue
		}

		if req.Index != sourceRequest.Index || req.RequestID == sourceRequest.RequestID || req.Status == constant.REQUEST_STATUS_REVIEWED {
			continue
		}

		req.Status = constant.REQUEST_STATUS_REVIEWED
		requestEntitiy := ModelsToEntity(req)
		requestEntitiy.UpdatedAt = time.Now()
		neededUpdateRequest = append(neededUpdateRequest, requestEntitiy)
	}
	return neededUpdateRequest
}
