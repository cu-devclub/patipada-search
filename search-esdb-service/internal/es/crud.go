package es

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"search-esdb-service/internal/dto"
	"search-esdb-service/internal/util"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

func BulkInsert(qars []*dto.QARecord) {
	es := GetESClient()
	var countSuccessful uint64
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:  "record", // The default index name
		Client: es,       // The Elasticsearch client
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}

	start := time.Now().UTC()

	// Loop over the collection
	for order, a := range qars {
		// Prepare the data payload: encode article to JSON
		//
		data, err := json.Marshal(a)
		if err != nil {
			log.Fatalf("Cannot encode record")
		}

		// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
		//
		// Add an item to the BulkIndexer
		//
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
			log.Fatalf("Unexpected error: %s", err)
		}
		// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	}

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// Close the indexer
	//
	if err := bi.Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}
	// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

	biStats := bi.Stats()

	// Report the results: number of indexed docs, number of errors, duration, indexing rate
	//
	log.Println(strings.Repeat("▔", 65))

	dur := time.Since(start)

	if biStats.NumFailed > 0 {
		log.Fatalf(
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
}

func InsertRecord(qar *dto.QARecord, documentID string) error {
	client := GetESClient()
	escapedYoutubeURL := util.EscapeText(qar.YoutubeURL)
	escapedQuestion := util.EscapeText(qar.Question)
	escapedAnswer := util.EscapeText(qar.Answer)
	escapedStartTime := util.EscapeText(qar.StartTime)
	escapedEndTime := util.EscapeText(qar.EndTime)

	document := fmt.Sprintf(`{
	    "youtubeURL": "%s",
	    "question": "%s",
	    "answer": "%s",
	    "startTime": "%s",
	    "endTime": "%s"
	}`, escapedYoutubeURL, escapedQuestion, escapedAnswer, escapedStartTime, escapedEndTime)

	req := esapi.IndexRequest{
		Index:      "record",
		DocumentID: documentID,
		Body:       strings.NewReader(document),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("Elasticsearch error: %s", res.Status())
	}

	return nil
}

func SearchInIndex(searchQuery string,indexName string,field string) ([]map[string]interface{}, error) {
	client := GetESClient()
	searchQuery = "*" + searchQuery + "*"
	// Create a search request
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query":  searchQuery,
				"fields": []string{field},
			},
		},
	}

	// Convert the query to JSON
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	// Perform the search request
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(indexName),
		client.Search.WithBody(strings.NewReader(string(queryJSON))),
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

	// Extract and return the matched documents
	var matchedDocuments []map[string]interface{}
	for _, hit := range hits {
		doc := hit.(map[string]interface{})["_source"]
		docID := hit.(map[string]interface{})["_id"].(string)
		// Unescape fields (e.g., "question" and "answer") individually before appending them
		unescapedDoc := make(map[string]interface{})
		for key, value := range doc.(map[string]interface{}) {
			if stringValue, isString := value.(string); isString {
				// Unescape the string value
				unescapedValue := util.UnescapeDoubleQuotes(stringValue)
				unescapedDoc[key] = unescapedValue
			} else {
				unescapedDoc[key] = value
			}
		}
		unescapedDoc["id"] = docID

		matchedDocuments = append(matchedDocuments, unescapedDoc)
	}
	return matchedDocuments, nil
}
