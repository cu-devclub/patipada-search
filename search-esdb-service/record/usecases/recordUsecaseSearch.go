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

func (r *recordUsecaseImpl) Search(indexName, query, searchType string, offset, amount int, countNeeded bool) (*models.SearchRecordStruct, error) {

	var records []*entities.Record
	var tokens []string
	var count int
	var err error

	switch searchType {
	case constant.SEARCH_BY_TF_IDF:
		records, tokens, count, err = r.internalSearch(indexName, query, offset, amount, countNeeded)
	case constant.SEARCH_BY_LDA:
		records, tokens, err = r.externalSearch(indexName, query, offset, amount)
	default:
		records, tokens, count, err = r.internalSearch(indexName, query, offset, amount, countNeeded)
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
		Amount:  count,
	}
	return response, nil
}

func (r *recordUsecaseImpl) internalSearch(indexName, query string, offset, amount int, countNeeded bool) ([]*entities.Record, []string, int, error) {
	// working (tokenize, tf-idf) with only elastic
	tokens, err := r.recordRepository.Tokenize(query)
	if err != nil {
		return nil, nil, 0, err
	}

	pureTokens := util.RemoveSliceFromArrays(tokens, data.GetStopWord())
	searchQuery := strings.Join(pureTokens, "")

	records, count, err := r.recordRepository.Search(indexName, searchQuery, offset, amount, countNeeded)
	if err != nil {
		return nil, nil, 0, err
	}

	return records, pureTokens, count, nil
}

func (r *recordUsecaseImpl) externalSearch(indexName, query string, offset, amount int) ([]*entities.Record, []string, error) {
	// TODO : implementing calling text 2 vec gateway	

	return nil, nil, nil
}
