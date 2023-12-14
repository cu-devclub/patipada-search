package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"search-esdb-service/record/entities"
	"search-esdb-service/record/es_query"
	"search-esdb-service/record/helper"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

type RecordESRepository struct {
	es *elasticsearch.Client
}

func NewRecordESRepository(es *elasticsearch.Client) RecordRepository {
	return &RecordESRepository{
		es: es,
	}
}

// GetAllRecords retrieves all records from the specified index in Elasticsearch.
//
// indexName: The name of the Elasticsearch index.
// []*entities.Record: An array of Record objects representing the retrieved documents.
// error: An error object if there was an issue retrieving the records.
func (r *RecordESRepository) GetAllRecords(indexName string) ([]*entities.Record, error) {
	client := r.es
	// Create a search request to retrieve all documents
	queryJSON, err := es_query.BuildMatchAllQuery()
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

// AnalyzeQueryKeyword analyzes the given query keyword.
//
// query: the query keyword to be analyzed.
// []string: a list of analyzed tokens.
// error: an error if the analysis fails.
func (r *RecordESRepository) AnalyzeQueryKeyword(query string) ([]string, error) {
	client := r.es

	analyzeQuery := es_query.BuildAnalyzeQuery("record", query)
	request := esapi.IndicesAnalyzeRequest{
		Index: "record",
		Body:  strings.NewReader(analyzeQuery),
	}

	// Perform the request
	response, err := request.Do(context.Background(), client)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	result, err := helper.ExtractTokens(responseBody)

	return result, err
}

// Search searches for records in the specified Elasticsearch index based on the provided query.
//
// Parameters:
// - indexName: The name of the Elasticsearch index to search in.
// - query: The query string used to search for records.
//
// Returns:
// - []*entities.Record: A slice of records found in the index that match the query.
// - error: An error if any occurred during the search operation.
func (r *RecordESRepository) Search(indexName, query string,amount int) ([]*entities.Record, error) {
	client := r.es

	// Build the Elasticsearch query
	queryJSON, err := es_query.BuildElasticsearchQuery(query)
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

// BulkInsert inserts multiple records into the Elasticsearch index.
//
// qars: A slice of pointers to Record entities representing the records to be inserted.
// Returns an error if there was an issue inserting the records.
func (r *RecordESRepository) BulkInsert(qars []*entities.Record) error {
	es := r.es
	var countSuccessful uint64
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  "record",
		Client: es, // The Elasticsearch client
	})
	if err != nil {
		return fmt.Errorf("Error creating the indexer: %s", err)
	}

	start := time.Now().UTC()

	// Loop over the collection
	for order, a := range qars {
		data, err := json.Marshal(a)

		if err != nil {
			return fmt.Errorf("Cannot encode data %v: %s", a.Question, err)
		}

		// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

		// Add an item to the BulkIndexer

		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				// Action field configures the operation to perform (index, create, delete, update)
				Action: "index",

				// DocumentID is the (optional) document ID
				DocumentID: a.YoutubeURL + "-" + strconv.Itoa(order),

				// Body is an `io.Reader` with the payload
				Body: bytes.NewReader(data),

				// OnSuccess is called for each successful operation
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},

				// OnFailure is called for each failed operation
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			return fmt.Errorf("Unexpected error: %s", err)
		}
		// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	}

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// Close the indexer
	//
	if err := bi.Close(context.Background()); err != nil {
		return fmt.Errorf("Unexpected error: %s", err)
	}
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	biStats := bi.Stats()

	// Report the results: number of indexed docs, number of errors, duration, indexing rate
	//
	log.Println(strings.Repeat("▔", 65))

	dur := time.Since(start)

	if biStats.NumFailed > 0 {
		return fmt.Errorf(
			"Indexed [%s] documents with [%s] errors in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			humanize.Comma(int64(biStats.NumFailed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	} else {
		log.Printf(
			"Sucessfuly indexed [%s] documents in %s (%s docs/sec)",
			humanize.Comma(int64(biStats.NumFlushed)),
			dur.Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed))),
		)
	}

	return nil
}
