package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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

func (r *RecordESRepository) KeywordSearch(keywordSearchEntity *entities.KeywordSearchStruct) ([]*entities.Record, int, error) {
	queryString, countQueryString, err := elasticQuery.BuildKeywordSearchQuery(keywordSearchEntity)
	if err != nil {
		return nil, 0, errors.CreateError(500, fmt.Sprintf("Error building keyword search query: %s", err))
	}

	records, err := r.performSearch(keywordSearchEntity.Config.IndexName, queryString)
	if err != nil {
		return nil, 0, errors.CreateError(500, fmt.Sprintf("Error performing keyword search: %s", err))
	}

	recordCount := 0
	if keywordSearchEntity.Config.CountNeeded {
		count, err := r.countRecordFromQuery(keywordSearchEntity.Config.IndexName, countQueryString)
		if err != nil {
			return nil, 0, errors.CreateError(500, fmt.Sprintf("Error counting records: %s", err))
		}
		recordCount = count
	}

	return records, recordCount, nil
}

func (r *RecordESRepository) VectorSearch(vectorSearchEntity *entities.VectorSearchStruct) ([]*entities.Record, int, error) {
	queryString, countQueryString, err := elasticQuery.BuildKNNQuery(vectorSearchEntity)
	if err != nil {
		return nil, 0, errors.CreateError(500, fmt.Sprintf("Error building knn search query: %s", err))
	}

	records, err := r.performSearch(vectorSearchEntity.Config.IndexName, queryString)
	if err != nil {
		return nil, 0, errors.CreateError(500, fmt.Sprintf("Error performing keyword search: %s", err))
	}

	recordCount := 0
	if vectorSearchEntity.Config.CountNeeded {
		count, err := r.countRecordFromQuery(vectorSearchEntity.Config.IndexName, countQueryString)
		if err != nil {
			return nil, 0, errors.CreateError(500, fmt.Sprintf("Error counting records: %s", err))
		}
		recordCount = count
	}

	return records, recordCount, nil
}

func (r *RecordESRepository) HybridSearch(hybridSearchEntity *entities.HybridSearchStruct) ([]*entities.Record, int, error) {
	queryString, countQueryString, err := elasticQuery.BuildHybridSearchQuery(hybridSearchEntity)
	if err != nil {
		return nil, 0, err
	}

	log.Println("Performing hybrid search query with:", queryString)
	records, err := r.performSearch(hybridSearchEntity.Config.IndexName, queryString)
	if err != nil {
		return nil, 0, errors.CreateError(500, fmt.Sprintf("Error performing keyword search: %s", err))
	}

	recordCount := 0
	if hybridSearchEntity.Config.CountNeeded {
		count, err := r.countRecordFromQuery(hybridSearchEntity.Config.IndexName, countQueryString)
		if err != nil {
			return nil, 0, errors.CreateError(500, fmt.Sprintf("Error counting records: %s", err))
		}
		recordCount = count
	}

	return records, recordCount, nil

}

func (r *RecordESRepository) performSearch(indexName string, queryString string) ([]*entities.Record, error) {
	client := r.es

	// Perform the search request
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(indexName),
		client.Search.WithBody(strings.NewReader(queryString)),
	)
	if err != nil {
		return nil, errors.CreateError(500, fmt.Sprintf("Error getting response: %s", err))
	}
	defer res.Body.Close()

	// Check the response status
	if res.IsError() {
		return nil, errors.CreateError(res.StatusCode, fmt.Sprintf("Elasticsearch error: %s", res.Status()))
	}

	// log.Println("RESPONSE BODY",res.Body)
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
		// log.Println("HIT DOC",doc["index"],hit.(map[string]interface{})["_score"],doc["lda-answer"])
		record := helper.UnescapeFieldsAndCreateRecord(doc, docID)
		records = append(records, record)
	}

	return records, nil
}

func (r *RecordESRepository) countRecordFromQuery(indexName string, query string) (int, error) {
	log.Println("count with query:", query)
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
		log.Println("error", countRes)
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
