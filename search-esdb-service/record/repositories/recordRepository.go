package repositories

import "search-esdb-service/record/entities"

type RecordRepository interface {
	GetAllRecords(indexName string) ([]*entities.Record, error)
	AnalyzeQueryKeyword(query string) ([]string, error)
	Search(indexName, query string,amount int) ([]*entities.Record, error)
	BulkInsert(qars []*entities.Record) error
}
