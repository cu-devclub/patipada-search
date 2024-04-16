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

func (r *RecordESRepository) GetAllRecords(indexName string) ([]*entities.Record, int, error) {
	return r.performSearch(indexName, 0, 0, elasticQuery.BuildMatchAllQuery, nil, false)
}

func (r *RecordESRepository) Search(indexName string, query interface{}, offset, amount int, countNeeded bool) ([]*entities.Record, int, error) {
	return r.performSearch(indexName, offset, amount, elasticQuery.BuildElasticsearchQuery, query, countNeeded)
}

func (r *RecordESRepository) VectorSearch(indexName string, query interface{}, offset, amount int, countNeeded bool) ([]*entities.Record, int, error) {
	return r.performSearch(indexName, offset, amount, elasticQuery.BuildKNNQuery, query, countNeeded)
}

func (r *RecordESRepository) performSearch(indexName string, offset, amount int, buildQueryFunc interface{}, query interface{}, countNeeded bool) ([]*entities.Record, int, error) {
	client := r.es

	var queryJSON string
	var countQueryJSON string
	var err error

	switch q := query.(type) {
	case string: // For keyword search
		queryFunc, ok := buildQueryFunc.(func(string, int, int) (string, string, error))
		if !ok {
			return nil, 0, errors.CreateError(500, "Invalid query builder function")
		}
		queryJSON, countQueryJSON, err = queryFunc(q, offset, amount)
	case []float64: // For vector search
		queryFunc, ok := buildQueryFunc.(func([]float64, string, int, int) (string, error))
		if !ok {
			return nil, 0, errors.CreateError(500, "Invalid query builder function")
		}
		queryJSON, err = queryFunc(q, "question_lda", offset, amount)
	case nil: // For match all query
		queryFunc, ok := buildQueryFunc.(func() (string, error))
		if !ok {
			return nil, 0, errors.CreateError(500, "Invalid query builder function")
		}
		queryJSON, err = queryFunc()
	default:
		return nil, 0, errors.CreateError(500, "Invalid query type")
	}

	if err != nil {
		return nil, 0, errors.CreateError(500, fmt.Sprintf("Error building query: %s", err))
	}

	// Perform the search request
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(indexName),
		client.Search.WithBody(strings.NewReader(string(queryJSON))),
	)
	if err != nil {
		return nil, 0, errors.CreateError(500, fmt.Sprintf("Error getting response: %s", err))
	}
	defer res.Body.Close()

	// Check the response status
	if res.IsError() {
		return nil, 0, errors.CreateError(res.StatusCode, fmt.Sprintf("Elasticsearch error: %s", res.Status()))
	}

	// Decode the response
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, 0, errors.CreateError(500, fmt.Sprintf("Error decoding response: %s", err))
	}

	// Extract and iterate through the hits (documents) in the response
	hits, found := response["hits"].(map[string]interface{})["hits"].([]interface{})
	if !found {
		return nil, 0, errors.CreateError(500, "Invalid response format")
	}

	var records []*entities.Record
	for _, hit := range hits {
		// log.Println("Hit:", hit.(map[string]interface{}))
		// log.Println("--------------------")
		doc := hit.(map[string]interface{})["_source"].(map[string]interface{})
		docID := hit.(map[string]interface{})["_id"].(string)
		record := helper.UnescapeFieldsAndCreateRecord(doc, docID)
		records = append(records, record)
	}

	//* Count
	recordCount := -1 // Don't count if not needed
	if countNeeded {
		recordCount, err = r.countRecordFromQuery(indexName, countQueryJSON)
		if err != nil {
			return nil, 0, errors.CreateError(500, fmt.Sprintf("Error counting records: %s", err))
		}
	}

	return records, recordCount, nil
}

func (r *RecordESRepository) countRecordFromQuery(indexName string, query string) (int, error) {
	client := r.es

	// Perform the count request
	countRes, err := client.Count(
		client.Count.WithContext(context.Background()),
		client.Count.WithIndex(indexName),
		client.Count.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return 0, errors.CreateError(500, fmt.Sprintf("Error getting count response: %s", err))
	}
	defer countRes.Body.Close()

	// Check the count response status
	if countRes.IsError() {
		return 0, errors.CreateError(countRes.StatusCode, fmt.Sprintf("Elasticsearch count error: %s", countRes.Status()))
	}

	// Decode the count response
	var countResponse map[string]interface{}
	if err := json.NewDecoder(countRes.Body).Decode(&countResponse); err != nil {
		return 0, errors.CreateError(500, fmt.Sprintf("Error decoding count response: %s", err))
	}

	// Extract the count from the count response
	count, found := countResponse["count"].(float64)
	if !found {
		return 0, errors.CreateError(500, "Invalid count response format")
	}

	return int(count), nil
}
