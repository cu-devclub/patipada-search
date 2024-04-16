package usecases

import (
	"search-esdb-service/constant"
	"search-esdb-service/data"
	"search-esdb-service/record/entities"
	"search-esdb-service/record/helper"
	"search-esdb-service/record/models"
	"search-esdb-service/util"
	"strings"
)

func (r *recordUsecaseImpl) Search(indexName, query, searchType string, offset, amount int) (*models.SearchRecordStruct, error) {

	var records []*entities.Record
	var tokens []string
	var err error

	switch searchType {
	case constant.SEARCH_BY_TF_IDF:
		records, tokens, err = r.internalSearch(indexName, query, offset, amount)
	case constant.SEARCH_BY_LDA:
		records, tokens, err = r.externalSearch(indexName, query, offset, amount)
	default:
		records, tokens, err = r.internalSearch(indexName, query, offset, amount)
	}

	if err != nil {
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

func (r *recordUsecaseImpl) internalSearch(indexName, query string, offset, amount int) ([]*entities.Record, []string, error) {
	// working (tokenize, tf-idf) with only elastic
	tokens, err := r.recordRepository.Tokenize(query)
	if err != nil {
		return nil, nil, err
	}

	pureTokens := util.RemoveSliceFromArrays(tokens, data.GetStopWord())
	searchQuery := strings.Join(pureTokens, "")

	records, err := r.recordRepository.Search(indexName, searchQuery, offset, amount)
	if err != nil {
		return nil, nil, err
	}

	return records, pureTokens, nil
}

func (r *recordUsecaseImpl) externalSearch(indexName, query string, offset, amount int) ([]*entities.Record, []string, error) {
	// working (tokenize, lda,...) with external service
	// TODO : make 1 service call for all tokenize, remove keyword and lda
	tokens, err := r.mlRepository.TokenizeQuery(query)
	if err != nil {
		return nil, nil, err
	}

	pureTokens := util.RemoveSliceFromArrays(tokens, data.GetStopWord())

	searchQuery, err := r.mlRepository.PerformLDATopicModelling(pureTokens)
	if err != nil {
		return nil, nil, err
	}

	records, err := r.recordRepository.VectorSearch(indexName, searchQuery, offset, amount)
	if err != nil {
		return nil, nil, err
	}

	return records, pureTokens, nil
}
