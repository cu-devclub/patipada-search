package usecases

import (
	"log"
	"search-esdb-service/record/helper"
	"search-esdb-service/record/models"
	"search-esdb-service/util"
)

// Search searches for records in the specified index using the given query.
//
// Parameters:
// - indexName: The name of the index to search in.
// - query: The query string used to search for records.
//
// Returns:
// - *models.SearchRecordStruct: The search results containing the matching records.
// - error: An error if the search operation fails.
func (r *recordUsecaseImpl) Search(indexName, query string, amount int) (*models.SearchRecordStruct, error) {
	// search the record
	records, err := r.recordRepository.Search(indexName, query, amount)
	if err != nil {
		log.Println("Error searching records: ", err)
		return nil, err
	}

	// extract tokens from query
	tokens, err := r.recordRepository.AnalyzeQueryKeyword(query)
	if err != nil {
		log.Println("Error extracting tokens: ", err)
		return nil, err
	}

	responseRecords := make([]*models.Record, 0)

	for _, record := range records {
		responseRecords = append(responseRecords, helper.RecordEntityToModels(record))
	}

	response := &models.SearchRecordStruct{
		Results: responseRecords,
		Tokens:  tokens,
	}
	return response, nil
}

func (r *recordUsecaseImpl) SearchByRecordIndex(indexName, recordIndex string) (*models.Record, error) {
	str, err := util.DecreaseIndexForSearchByIndex(recordIndex)
	if err != nil {
		return nil, err
	}
	// search the record
	records, err := r.recordRepository.SearchByRecordIndex(indexName, str)
	if err != nil {
		if err.Error() == "Elasticsearch error: 404 Not Found" {
			return nil, nil
		} else if err.Error() != "Elasticsearch error: 405 Method Not Allowed" {
			// 405 is because gRPC we can ignore it
			return nil, err
		}
	}
	response := helper.RecordEntityToModels(records)
	return response, nil
}
