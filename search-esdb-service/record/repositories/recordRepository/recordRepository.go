package repositories

import (
	"search-esdb-service/record/entities"
)

type RecordRepository interface {

	// KeywordSearch searches for records in the specified Elasticsearch index based on the provided query.
	//
	// Parameters:
	// - indexName: The name of the Elasticsearch index to search in.
	// - query: The query string used to search for records.
	//
	// Returns:
	// - []*entities.Record: A slice of records found in the index that match the query.
	// - error: An error if any occurred during the search operation.
	KeywordSearch(keywordSearchEntity *entities.KeywordSearchStruct) ([]*entities.Record, int, error)

	VectorSearch(vectorSearchEntity *entities.VectorSearchStruct) ([]*entities.Record, int, error)

	HybridSearch(hybridSearchEntity *entities.HybridSearchStruct) ([]*entities.Record, int, error)

	Tokenize(query string) ([]string, error)

	SearchByRecordIndex(indexName, recordIndex string) (*entities.Record, bool, error)

	// BulkInsert inserts multiple records into the Elasticsearch index.
	//
	// qars: A slice of pointers to Record entities representing the records to be inserted.
	// Returns an error if there was an issue inserting the records.
	BulkInsert(qars []*entities.Record) error

	UpdateRecord(record *entities.UpdateRecord) error
}
