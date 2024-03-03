package usecases

import (
	"search-esdb-service/record/helper"
	"search-esdb-service/record/models"
	mlRepository "search-esdb-service/record/repositories/mlRepository"
	recordRepository "search-esdb-service/record/repositories/recordRepository"
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

func (r *recordUsecaseImpl) GetAllRecords(indexName string) ([]*models.Record, error) {
	records, err := r.recordRepository.GetAllRecords(indexName)
	if err != nil {
		return nil, err
	}

	responseRecords := make([]*models.Record, 0)
	for _, r := range records {
		responseRecords = append(responseRecords, helper.RecordEntityToModels(r))
	}

	return responseRecords, nil
}
