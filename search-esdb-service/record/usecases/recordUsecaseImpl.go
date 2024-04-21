package usecases

import (
	"search-esdb-service/config"
	mlRepository "search-esdb-service/record/repositories/mlRepository"
	recordRepository "search-esdb-service/record/repositories/recordRepository"
)

type recordUsecaseImpl struct {
	recordRepository recordRepository.RecordRepository
	mlRepository     mlRepository.MLRepository
	cfg              *config.Config
}

func NewRecordUsecase(
	recordRepository recordRepository.RecordRepository,
	mlRepository mlRepository.MLRepository,
	cfg *config.Config,
) RecordUsecase {
	return &recordUsecaseImpl{
		recordRepository: recordRepository,
		mlRepository:     mlRepository,
		cfg:              cfg,
	}
}
