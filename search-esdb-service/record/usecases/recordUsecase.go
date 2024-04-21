package usecases

import (
	"search-esdb-service/record/models"
)

type RecordUsecase interface {

	// Search searches for records in the specified index using the given query.
	//
	// Parameters:
	// - indexName: The name of the index to search in.
	// - query: The query string used to search for records.
	//
	// Returns:
	// - *models.SearchRecordStruct: The search results containing the matching records.
	// - error: An error if the search operation fails.
	Search(indexName, query, searchType string, offset, amount int, countNeeded bool) (*models.SearchRecordStruct, error)

	SearchByRecordIndex(indexName, recordIndex string) (*models.Record, error)

	UpdateRecord(record *models.UpdateRecord) error
}
