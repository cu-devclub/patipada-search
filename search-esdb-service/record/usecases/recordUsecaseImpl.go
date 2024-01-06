package usecases

import (
	"search-esdb-service/record/repositories"
)

type recordUsecaseImpl struct {
	recordRepository repositories.RecordRepository
}

func NewRecordUsecase(recordRepository repositories.RecordRepository) RecordUsecase {
	return &recordUsecaseImpl{
		recordRepository: recordRepository,
	}
}
