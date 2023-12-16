package usecases

import "search-esdb-service/record/models"

type RecordUsecase interface {
	GetAllRecords(indexName string) ([]*models.Record, error)
	Search(indexName, query string,amount int) (*models.SearchRecordStruct, error)
}