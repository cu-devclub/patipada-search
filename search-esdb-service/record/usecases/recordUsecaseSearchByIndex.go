package usecases

import (
	"search-esdb-service/errors"
	"search-esdb-service/messages"
	"search-esdb-service/record/helper"
	"search-esdb-service/record/models"
)

func (r *recordUsecaseImpl) SearchByRecordIndex(indexName, recordIndex string) (*models.Record, error) {
	// search the record
	records, isFound, err := r.recordRepository.SearchByRecordIndex(indexName, recordIndex)
	if !isFound && err != nil {
		if err.Error() == messages.ELASTIC_404_ERROR {
			return nil, nil
		} else if err.Error() != messages.ELASTIC_405_ERROR {
			// 405 is because gRPC we can ignore it
			return nil, errors.CreateError(500, err.Error())
		}
	}

	response := helper.RecordEntityToModels(records)
	return response, nil
}
