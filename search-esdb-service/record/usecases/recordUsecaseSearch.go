package usecases

import (
	"fmt"
	"search-esdb-service/constant"
	"search-esdb-service/data"
	"search-esdb-service/errors"
	"search-esdb-service/record/entities"
	"search-esdb-service/record/helper"
	"search-esdb-service/record/models"
	"search-esdb-service/util"
	"strings"
)

func (r *recordUsecaseImpl) Search(indexName, query, searchType string, offset, amount int, countNeeded bool) (*models.SearchRecordStruct, error) {

	var records []*entities.Record
	var count int

	tokens, err := r.recordRepository.Tokenize(query)
	if err != nil {
		return nil, err
	}
	pureTokens := util.RemoveSliceFromArrays(tokens, data.GetStopWord())

	config := &entities.SearchConfig{
		IndexName:   indexName,
		Offset:      offset,
		Amount:      amount,
		CountNeeded: countNeeded,
	}

	switch searchType {
	case constant.SEARCH_BY_TF_IDF:
		records, count, err = r.keywordSearch(indexName, pureTokens, config)
	case constant.SEARCH_BY_VECTOR:
		records, count, err = r.vectorSearch(indexName, query, config)
	default:
		records, count, err = r.hybridSearch(indexName, query, pureTokens, config)
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
		Tokens:  pureTokens,
		Amount:  count,
	}
	return response, nil
}

func (r *recordUsecaseImpl) keywordSearch(indexName string, tokens []string, config *entities.SearchConfig) ([]*entities.Record, int, error) {
	searchQuery := strings.Join(tokens, "")

	keywordSearchStruct := &entities.KeywordSearchStruct{
		Query:               searchQuery,
		KeywordSearchFields: []string{"question"},
		Config:              config,
	}

	records, count, err := r.recordRepository.KeywordSearch(keywordSearchStruct)
	if err != nil {
		return nil, 0, err
	}

	return records, count, nil
}

func (r *recordUsecaseImpl) vectorSearch(indexName, query string, config *entities.SearchConfig) ([]*entities.Record, int, error) {
	vectorResponses, err := r.mlRepository.Text2VecGateway(query)
	if err != nil {
		return nil, 0, err
	}

	vectorSearchStruct := &entities.VectorSearchStruct{
		VectorFields: vectorResponses,
		Config:       config,
	}

	records, count, err := r.recordRepository.VectorSearch(vectorSearchStruct)
	if err != nil {
		return nil, 0, errors.CreateError(500, fmt.Sprintf("Error performing vector search: %v", err))
	}

	return records, count, nil
}

func (r *recordUsecaseImpl) hybridSearch(indexName, query string, tokens []string, config *entities.SearchConfig) ([]*entities.Record, int, error) {
	keywordSearchQuery := strings.Join(tokens, "")

	vectorResponses, err := r.mlRepository.Text2VecGateway(query)
	if err != nil {
		return nil, 0, err
	}

	hybridSearchStruct := &entities.HybridSearchStruct{
		Query:                    keywordSearchQuery,
		KeywordSearchFields:      []string{"question"},
		KeywordSearchScoreWeight: r.cfg.MlConfig.TfIDFScoreWeight,
		VectorFields:             vectorResponses,
		Config:                   config,
	}

	records, count, err := r.recordRepository.HybridSearch(hybridSearchStruct)
	if err != nil {
		return nil, 0, err
	}

	return records, count, nil
}
