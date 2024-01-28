package helper

import (
	"search-esdb-service/record/entities"
	"search-esdb-service/record/models"
)

func RecordEntityToModels(e *entities.Record) *models.Record {
	return &models.Record{
		Index:      e.Index,
		YoutubeURL: e.YoutubeURL,
		Question:   e.Question,
		Answer:     e.Answer,
		StartTime:  e.StartTime,
		EndTime:    e.EndTime,
	}
}

func UpdateRecordModelToEntity(m *models.UpdateRecord) *entities.UpdateRecord {
	return &entities.UpdateRecord{
		DocumentID: m.DocumentID,
		StartTime:  m.StartTime,
		EndTime:    m.EndTime,
		Question:   m.Question,
		Answer:     m.Answer,
	}
}
