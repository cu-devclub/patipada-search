package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"search-esdb-service/record/entities"
	"search-esdb-service/record/helper"
	"strings"
)

// GetAllRecords retrieves all records from the specified index in Elasticsearch.
//
// indexName: The name of the Elasticsearch index.
// []*entities.Record: An array of Record objects representing the retrieved documents.
// error: An error object if there was an issue retrieving the records.
func (r *RecordESRepository) GetAllRecords(indexName string) ([]*entities.Record, error) {
	client := r.es
	// Create a search request to retrieve all documents
	queryJSON, err := buildMatchAllQuery()
	if err != nil {
		return nil, err
	}

	// Create a search request
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(indexName),
		client.Search.WithBody(strings.NewReader(string(queryJSON))))
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
		return nil, fmt.Errorf("No hits found in the response")
	}

	var records []*entities.Record
	for _, hit := range hits {
		doc := hit.(map[string]interface{})["_source"].(map[string]interface{})

		// Map Elasticsearch document fields to Record struct fields
		record := &entities.Record{
			Index:      indexName,
			YoutubeURL: helper.GetStringField(doc, "youtubeURL"),
			Question:   helper.GetStringField(doc, "question"),
			Answer:     helper.GetStringField(doc, "answer"),
			StartTime:  helper.GetStringField(doc, "startTime"),
			EndTime:    helper.GetStringField(doc, "endTime"),
		}

		records = append(records, record)
	}

	return records, nil
}

func buildMatchAllQuery() (string, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		return "", err
	}

	return string(queryJSON), nil
}
