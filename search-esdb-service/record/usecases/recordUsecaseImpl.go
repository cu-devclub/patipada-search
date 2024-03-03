package usecases

import (
	recordRepository "search-esdb-service/record/repositories/recordRepository"
	mlRepository "search-esdb-service/record/repositories/mlRepository"
)

type recordUsecaseImpl struct {
	recordRepository recordRepository.RecordRepository
	mlRepository     mlRepository.MLRepository
}

func NewRecordUsecase(recordRepository recordRepository.RecordRepository, mlRepository mlRepository.MLRepository) RecordUsecase {
	return &recordUsecaseImpl{
		recordRepository: recordRepository,
		mlRepository:     mlRepository,
	}
}
