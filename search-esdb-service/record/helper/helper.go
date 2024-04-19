package helper

import (
	"search-esdb-service/proto/ml_gateway_proto"
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

func ConvertgRPCText2VecResonseToEntityResponses(grpcRes *ml_gateway_proto.Text2VecResponse) []*entities.Text2VecResponse {
	var res []*entities.Text2VecResponse
	for _, r := range grpcRes.Results {
		res = append(res, &entities.Text2VecResponse{
			Embedding:   r.Embedding,
			Name:        r.Name,
			ScoreWeight: r.Score,
		})
	}
	return res
}
