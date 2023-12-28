package usecases

import "search-esdb-service/record/models"

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
		return nil, err
	}

	// extract tokens from query
	tokens, err := r.recordRepository.AnalyzeQueryKeyword(query)
	if err != nil {
		return nil, err
	}

	responseRecords := make([]*models.Record, 0)

	for _, record := range records {
		responseRecords = append(responseRecords, &models.Record{
			Index:      record.Index,
			YoutubeURL: record.YoutubeURL,
			Question:   record.Question,
			Answer:     record.Answer,
			StartTime:  record.StartTime,
			EndTime:    record.EndTime,
		})
	}

	response := &models.SearchRecordStruct{
		Results: responseRecords,
		Tokens:  tokens,
	}
	return response, nil
}
