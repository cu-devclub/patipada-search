package usecases

import (
	"search-esdb-service/errors"
	"search-esdb-service/record/models"
)

type RecordUsecase interface {
	// GetAllRecords retrieves all records from the specified index.
	//
	// Parameters:
	// - indexName: The name of the index to retrieve records from.
	//
	// Returns:
	// - []*models.Record: An array of record objects.
	// - error: Any error that occurred during the retrieval process.
	GetAllRecords(indexName string) ([]*models.Record, error)

	// Search searches for records in the specified index using the given query.
	//
	// Parameters:
	// - indexName: The name of the index to search in.
	// - query: The query string used to search for records.
	//
	// Returns:
	// - *models.SearchRecordStruct: The search results containing the matching records.
	// - error: An error if the search operation fails.
	Search(indexName, query, searchType string, amount int) (*models.SearchRecordStruct, error)

	SearchByRecordIndex(indexName, recordIndex string) (*models.Record, error)

	UpdateRecord(record *models.UpdateRecord) *errors.RequestError
}
