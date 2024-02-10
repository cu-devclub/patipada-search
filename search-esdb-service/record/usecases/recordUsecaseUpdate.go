package usecases

import (
	"log"
	"search-esdb-service/errors"
	"search-esdb-service/record/helper"
	"search-esdb-service/record/models"
)

func (r *recordUsecaseImpl) UpdateRecord(record *models.UpdateRecord) *errors.RequestError {
	log.Println("UpdateRecord with record: ", record.ToString())
	updateRecordEntity := helper.UpdateRecordModelToEntity(record)
	if err := r.recordRepository.UpdateRecord(updateRecordEntity); err != nil {
		return err
	}

	log.Println("Record updated successfully")

	return nil
}
