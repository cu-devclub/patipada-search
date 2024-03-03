package usecases

import (
	"search-esdb-service/constant"
	"search-esdb-service/errors"
	"search-esdb-service/messages"
	"search-esdb-service/record/entities"
	"search-esdb-service/record/helper"
	"search-esdb-service/record/models"
	"strings"
)

func (r *recordUsecaseImpl) GetAllRecords(indexName string) ([]*models.Record, error) {
	records, err := r.recordRepository.GetAllRecords(indexName)
	if err != nil {
		return nil, err
	}

	responseRecords := make([]*models.Record, 0)
	for _, r := range records {
		responseRecords = append(responseRecords, helper.RecordEntityToModels(r))
	}

	return responseRecords, nil
}

func (r *recordUsecaseImpl) Search(indexName, query, searchType string, amount int) (*models.SearchRecordStruct, error) {
	var records []*entities.Record

	var searchQuery interface{}
	// extract tokens from query
	tokens, err := r.mlRepository.TokenizeQuery(query)
	if err != nil {
		return nil, err
	}

	pureQuery, err := r.mlRepository.RemoveStopWordFromTokensArrays(tokens)
	if err != nil {
		return nil, err
	}

	switch searchType {
	case constant.SEARCH_BY_TF_IDF:
		q := strings.Join(pureQuery, "")
		searchQuery = q
		records, err = r.recordRepository.Search(indexName, searchQuery, amount)
		if err != nil {
			return nil, err
		}
	case constant.SEARCH_BY_LDA:
		searchQuery, err = r.mlRepository.PerformLDATopicModelling(tokens)
		if err != nil {
			return nil, err
		}
		records, err = r.recordRepository.VectorSearch(indexName, searchQuery, amount)
		if err != nil {
			return nil, err
		}
	default:
		// DO nothing
		searchQuery = query
		records, err = r.recordRepository.Search(indexName, searchQuery, amount)
		if err != nil {
			return nil, err
		}
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

// TODO : Preparation Query function for multiple
// e.g. 1 : Keyword search with remove stop word
// e.g. 2 : TF-IDF search with remove stop word
// e.g. 3 : LDA

// Query -> [ word tokenize -> remove stop word ] -> Bag of words -> Search : Duplicate
// Query -> [ word tokenize -> remove stop word  ]-> Bag of words -> TF-IDF -> Search : Done
// Query -> [ word tokenize -> remove stop word -> Bag of words -> LDA -> Topic ] -> Search
// RAW data -> [ remove stop word -> Bag of word -> LDA -> [vector] ] -> PROCESSED CSV
// Migrate start service
// QUERY -> [ remove stop word -> Bag of word -> LDA -> [vector] ] -> cosine similarity -> ELASTIC
// [....] => external
