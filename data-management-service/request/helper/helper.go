package helper

import (
	"data-management/request/entities"
	"data-management/request/models"
	"data-management/request/repositories"
	"fmt"
)

// This will generate a unique RequestID for each new record, like "REQ1", "REQ2", etc.
// using IncrementRecordCounter() from repositories
func GenerateRequestID(r repositories.Repositories) (string, error) {
	nextSeq, err := r.IncrementRequestCounter()
	if err != nil {
		return "", err
	}
	nextRequestID := fmt.Sprintf("REQ%d", nextSeq)
	return nextRequestID, nil
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
