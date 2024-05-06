package repositories

import (
	"data-management/constant"
	"data-management/request/entities"
)

func (r *repositoryImpl) UpdateRecord(record *entities.Record) error {
	entity := &entities.UpdateRecord{
		DocumentID: record.Index,
		Question:   record.Question,
		Answer:     record.Answer,
		StartTime:  record.StartTime,
		EndTime:    record.EndTime,
	}
	err := r.communicationClient.PublishUpdateRecordsToRabbitMQ(
		constant.UPDATE_RECORD_PAYLOAD_NAME,
		entity,
	)
	if err != nil {
		return err
	}
	return nil
}
