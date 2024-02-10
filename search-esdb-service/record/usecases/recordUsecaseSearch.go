package usecases

import (
	"search-esdb-service/constant"
	"search-esdb-service/errors"
	"search-esdb-service/record/entities"
	"search-esdb-service/record/helper"
	"search-esdb-service/record/models"
	"search-esdb-service/util"
)

func (r *recordUsecaseImpl) GetAllRecords(indexName string) ([]*models.Record, *errors.RequestError) {
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

func (r *recordUsecaseImpl) Search(indexName, query, searchType string, amount int) (*models.SearchRecordStruct, *errors.RequestError) {
	var records []*entities.Record

	// extract tokens from query
	tokens, err := r.recordRepository.AnalyzeQueryKeyword(query)
	if err != nil {
		return nil, err
	}

	switch searchType {
	case constant.SEARCH_BY_TOKENS:
		stopWords := r.dataI.GetStopWord()
		tokens = helper.RemoveStopWordsFromTokensArray(stopWords.PythaiNLP, tokens)
		records, err = r.recordRepository.SearchByTokens(indexName, tokens, amount)
	default:
		records, err = r.recordRepository.Search(indexName, query, amount)
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

func (r *recordUsecaseImpl) SearchByRecordIndex(indexName, recordIndex string) (*models.Record, *errors.RequestError) {
	str, err := util.DecreaseIndexForSearchByIndex(recordIndex)
	if err != nil {
		return nil, errors.CreateError(400, err.Error())
	}
	// search the record
	records, err := r.recordRepository.SearchByRecordIndex(indexName, str)
	if err != nil {
		if err.Error() == "Elasticsearch error: 404 Not Found" {
			return nil, nil
		} else if err.Error() != "Elasticsearch error: 405 Method Not Allowed" {
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

// Query -> word tokenize -> remove stop word -> Bag of words -> Search : Done 
// Query -> word tokenize -> remove stop word -> Bag of words -> TF-IDF -> Search
// Query -> word tokenize -> remove stop word -> Bag of words -> LDA -> Topic -> Search
