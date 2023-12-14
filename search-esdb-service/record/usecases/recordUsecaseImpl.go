package usecases

import (
	"search-esdb-service/record/models"
	"search-esdb-service/record/repositories"
)

type recordUsecaseImpl struct {
	recordRepository repositories.RecordRepository
}

func NewRecordUsecase(recordRepository repositories.RecordRepository) RecordUsecase {
	return &recordUsecaseImpl{
		recordRepository: recordRepository,
	}
}

// GetAllRecords retrieves all records from the specified index.
//
// Parameters:
// - indexName: The name of the index to retrieve records from.
//
// Returns:
// - []*models.Record: An array of record objects.
// - error: Any error that occurred during the retrieval process.
func (r *recordUsecaseImpl) GetAllRecords(indexName string) ([]*models.Record, error) {
	records, err := r.recordRepository.GetAllRecords(indexName)
	if err != nil {
		return nil, err
	}

	responseRecords := make([]*models.Record, 0)
	for _, r := range records {
		responseRecords = append(responseRecords, &models.Record{
			Index:      r.Index,
			YoutubeURL: r.YoutubeURL,
			Question:   r.Question,
			Answer:     r.Answer,
			StartTime:  r.StartTime,
			EndTime:    r.EndTime,
		})
	}

	return responseRecords, nil
}

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
	records, err := r.recordRepository.Search(indexName, query,amount)
	if err != nil {
		return nil, err
	}

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
