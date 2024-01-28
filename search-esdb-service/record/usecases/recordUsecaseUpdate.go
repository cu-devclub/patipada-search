package usecases

import (
	"log"
	"search-esdb-service/record/helper"
	"search-esdb-service/record/models"
)

func (r *recordUsecaseImpl) UpdateRecord(record *models.UpdateRecord) error {
	log.Println("Update record usecase", record)

	updateRecordEntity := helper.UpdateRecordModelToEntity(record)
	if err := r.recordRepository.UpdateRecord(updateRecordEntity) ; err != nil {
		log.Println("Error update record: ", err)
		return err
	}
	
	return nil
}
