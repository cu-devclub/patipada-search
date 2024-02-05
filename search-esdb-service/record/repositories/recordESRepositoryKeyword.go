package repositories

import (
	"context"
	"fmt"
	"io"
	"search-esdb-service/record/helper"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)
// TODO : Preparation Query function for multiple
// e.g. 1 : Keyword search with remove stop word 
// e.g. 2 : TF-IDF search with remove stop word
// e.g. 3 : LDA 
// AnalyzeQueryKeyword analyzes the given query keyword.
//
// query: the query keyword to be analyzed.
// []string: a list of analyzed tokens.
// error: an error if the analysis fails.
func (r *RecordESRepository) AnalyzeQueryKeyword(query string) ([]string, error) {
	client := r.es

	analyzeQuery := buildAnalyzeQuery("record", query)
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

func buildAnalyzeQuery(index, query string) string {
	return fmt.Sprintf(`{
        "tokenizer": "icu_tokenizer",
        "text": "%s"
    }`, query)
}


// Query -> word tokenize -> remove stop word -> Bag of words -> LDA -> Topic -> Search
// Query -> word tokenize -> remove stop word -> Bag of words -> TF-IDF -> Search