package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"search-esdb-service/errors"
	"search-esdb-service/record/entities"
	"search-esdb-service/record/helper"
	"search-esdb-service/record/repositories/elasticQuery"
	"strings"
)

func (r *RecordESRepository) SearchByRecordIndex(indexName, recordIndex string) (*entities.Record, bool, error) {
	client := r.es

	recordIndex = url.PathEscape(recordIndex)

	// Perform the search request
	res, err := client.Get(indexName, recordIndex)
	if err != nil {
		return nil, false, errors.CreateError(500, fmt.Sprintf("Error getting record: %s", err))
	}
	defer res.Body.Close()

	// Check the response status
	if res.IsError() && res.StatusCode != 405 {
		return nil, false, errors.CreateError(res.StatusCode, fmt.Sprintf("Elasticsearch error: %s", res.Status()))
	}

	// Decode the response
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, false, errors.CreateError(500, fmt.Sprintf("Error decoding response: %s", err))
	}

	doc := response["_source"]
	docID := response["_id"].(string)

	record := helper.UnescapeFieldsAndCreateRecord(doc, docID)
	return record, true, nil
}

func (r *RecordESRepository) GetAllRecords(indexName string) ([]*entities.Record, error) {
	return r.performSearch(indexName, 0, elasticQuery.BuildMatchAllQuery, nil)
}

func (r *RecordESRepository) Search(indexName string, query interface{}, amount int) ([]*entities.Record, error) {
	return r.performSearch(indexName, amount, elasticQuery.BuildElasticsearchQuery, query)
}

func (r *RecordESRepository) VectorSearch(indexName string, query interface{}, amount int) ([]*entities.Record, error) {
	return r.performSearch(indexName, amount, elasticQuery.BuildKNNQuery, query)
}

func (r *RecordESRepository) performSearch(indexName string, amount int, buildQueryFunc interface{}, query interface{}) ([]*entities.Record, error) {
	client := r.es

	var queryJSON string
	var err error

	switch q := query.(type) {
	case string:
		queryFunc, ok := buildQueryFunc.(func(string) (string, error))
		if !ok {
			return nil, errors.CreateError(500, "Invalid query builder function")
		}
		queryJSON, err = queryFunc(q)
	case []float64:
		queryFunc, ok := buildQueryFunc.(func([]float64, string) (string, error))
		if !ok {
			return nil, errors.CreateError(500, "Invalid query builder function")
		}
		queryJSON, err = queryFunc(q, "question_lda")
	case nil:
		queryFunc, ok := buildQueryFunc.(func() (string, error))
		if !ok {
			return nil, errors.CreateError(500, "Invalid query builder function")
		}
		queryJSON, err = queryFunc()
	default:
		return nil, errors.CreateError(500, "Invalid query type")
	}

	if err != nil {
		return nil, errors.CreateError(500, fmt.Sprintf("Error building query: %s", err))
	}

	// Perform the search request
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(indexName),
		client.Search.WithBody(strings.NewReader(string(queryJSON))),
		client.Search.WithSize(amount),
	)
	if err != nil {
		return nil, errors.CreateError(500, fmt.Sprintf("Error getting response: %s", err))
	}
	defer res.Body.Close()

	// Check the response status
	if res.IsError() {
		return nil, errors.CreateError(res.StatusCode, fmt.Sprintf("Elasticsearch error: %s", res.Status()))
	}

	// Decode the response
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, errors.CreateError(500, fmt.Sprintf("Error decoding response: %s", err))
	}

	// Extract and iterate through the hits (documents) in the response
	hits, found := response["hits"].(map[string]interface{})["hits"].([]interface{})
	if !found {
		return nil, errors.CreateError(500, "Invalid response format")
	}

	var records []*entities.Record
	for _, hit := range hits {
		doc := hit.(map[string]interface{})["_source"].(map[string]interface{})
		docID := hit.(map[string]interface{})["_id"].(string)
		record := helper.UnescapeFieldsAndCreateRecord(doc, docID)
		records = append(records, record)
	}

	return records, nil
}
