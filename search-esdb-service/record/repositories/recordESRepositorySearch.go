package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"search-esdb-service/record/entities"
	"search-esdb-service/record/helper"
	"strings"
)

// Search searches for records in the specified Elasticsearch index based on the provided query.
//
// Parameters:
// - indexName: The name of the Elasticsearch index to search in.
// - query: The query string used to search for records.
//
// Returns:
// - []*entities.Record: A slice of records found in the index that match the query.
// - error: An error if any occurred during the search operation.
func (r *RecordESRepository) Search(indexName, query string, amount int) ([]*entities.Record, error) {
	client := r.es

	// Build the Elasticsearch query
	queryJSON, err := buildElasticsearchQuery(query)
	if err != nil {
		return nil, err
	}

	// Perform the search request
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(indexName),
		client.Search.WithBody(strings.NewReader(string(queryJSON))),
		client.Search.WithSize(amount),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Check the response status
	if res.IsError() {
		return nil, fmt.Errorf("Elasticsearch error: %s", res.Status())
	}

	// Decode the response
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	// Extract and iterate through the hits (documents) in the response
	hits, found := response["hits"].(map[string]interface{})["hits"].([]interface{})
	if !found {
		return nil, errors.New("No hits found in the response")
	}

	var records []*entities.Record
	for _, hit := range hits {
		doc := hit.(map[string]interface{})["_source"]
		docID := hit.(map[string]interface{})["_id"].(string)
		// Unescape fields (e.g., "question" and "answer") individually before appending them
		unescapedDoc := make(map[string]interface{})
		for key, value := range doc.(map[string]interface{}) {
			if stringValue, isString := value.(string); isString {
				// Unescape the string value
				unescapedValue := helper.UnescapeDoubleQuotes(stringValue)
				unescapedDoc[key] = unescapedValue
			} else {
				unescapedDoc[key] = value
			}
		}
		unescapedDoc["id"] = docID

		record := &entities.Record{
			Index:      docID,
			YoutubeURL: unescapedDoc["youtubeURL"].(string),
			Question:   unescapedDoc["question"].(string),
			Answer:     unescapedDoc["answer"].(string),
			StartTime:  unescapedDoc["startTime"].(string),
			EndTime:    unescapedDoc["endTime"].(string),
		}
		records = append(records, record)
	}

	return records, nil
}

func buildElasticsearchQuery(query string) (string, error) {
	// Build the Elasticsearch query
	queryString := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"multi_match": map[string]interface{}{
							"query":  query,
							"fields": []string{"question", "answer"},
						},
					},
					{
						"multi_match": map[string]interface{}{
							"query":  query,
							"type":   "phrase_prefix",
							"fields": []string{"question", "answer"},
						},
					},
					{
						"term": map[string]interface{}{
							"_id": query,
						},
					},
				},
			},
		},
	}

	// Convert the query to JSON
	queryJSON, err := json.Marshal(queryString)
	if err != nil {
		return "", err
	}

	return string(queryJSON), nil
}


func (r *RecordESRepository) SearchByRecordIndex(indexName, recordIndex string) (*entities.Record, error) {
	client := r.es

	// Perform the search request
	res, err := client.Get(indexName, recordIndex)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Check the response status
	if res.IsError() {
		return nil, fmt.Errorf("Elasticsearch error: %s", res.Status())
	}

	// Decode the response
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	doc := response["_source"]
	docID := response["_id"].(string)
	// Unescape fields (e.g., "question" and "answer") individually before appending them
	unescapedDoc := make(map[string]interface{})
	for key, value := range doc.(map[string]interface{}) {
		if stringValue, isString := value.(string); isString {
			// Unescape the string value
			unescapedValue := helper.UnescapeDoubleQuotes(stringValue)
			unescapedDoc[key] = unescapedValue
		} else {
			unescapedDoc[key] = value
		}
	}
	unescapedDoc["id"] = docID

	record := &entities.Record{
		Index:      docID,
		YoutubeURL: unescapedDoc["youtubeURL"].(string),
		Question:   unescapedDoc["question"].(string),
		Answer:     unescapedDoc["answer"].(string),
		StartTime:  unescapedDoc["startTime"].(string),
		EndTime:    unescapedDoc["endTime"].(string),
	}

	return record, nil
}