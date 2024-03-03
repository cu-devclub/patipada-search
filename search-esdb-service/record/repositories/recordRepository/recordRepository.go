package repositories

import (
	"search-esdb-service/record/entities"
)

type RecordRepository interface {

	// GetAllRecords retrieves all records from the specified index in Elasticsearch.
	//
	// indexName: The name of the Elasticsearch index.
	// []*entities.Record: An array of Record objects representing the retrieved documents.
	// error: An error object if there was an issue retrieving the records.
	GetAllRecords(indexName string) ([]*entities.Record, error)

	// TODO : Tokenize service connection with ML service

	// Search searches for records in the specified Elasticsearch index based on the provided query.
	//
	// Parameters:
	// - indexName: The name of the Elasticsearch index to search in.
	// - query: The query string used to search for records.
	//
	// Returns:
	// - []*entities.Record: A slice of records found in the index that match the query.
	// - error: An error if any occurred during the search operation.
	Search(indexName string, query interface{}, amount int) ([]*entities.Record, error)

	VectorSearch(indexName string, query interface{}, amount int) ([]*entities.Record, error)

	Tokenize(query string) ([]string, error)

	SearchByRecordIndex(indexName, recordIndex string) (*entities.Record, bool, error)

	// BulkInsert inserts multiple records into the Elasticsearch index.
	//
	// qars: A slice of pointers to Record entities representing the records to be inserted.
	// Returns an error if there was an issue inserting the records.
	BulkInsert(qars []*entities.Record) error

	UpdateRecord(record *entities.UpdateRecord) error
}
