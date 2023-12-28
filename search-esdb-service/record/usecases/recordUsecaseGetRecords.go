package usecases

import "search-esdb-service/record/models"

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
