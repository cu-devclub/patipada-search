package usecases

import (
	"search-esdb-service/data"
	"search-esdb-service/record/repositories"
)

type recordUsecaseImpl struct {
	recordRepository repositories.RecordRepository
	dataI            data.Data
}

func NewRecordUsecase(recordRepository repositories.RecordRepository, dataI data.Data) RecordUsecase {
	return &recordUsecaseImpl{
		recordRepository: recordRepository,
		dataI:            dataI,
	}
}
