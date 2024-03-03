package repositories

import (
	"context"
	"io"
	"search-esdb-service/errors"
	"search-esdb-service/record/helper"
	"search-esdb-service/record/repositories/elasticQuery"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func (r *RecordESRepository) Tokenize(query string) ([]string, error) {
	client := r.es

	analyzeQuery := elasticQuery.BuildAnalyzeQuery("record", query)
	request := esapi.IndicesAnalyzeRequest{
		Index: "record",
		Body:  strings.NewReader(analyzeQuery),
	}

	// Perform the request
	response, err := request.Do(context.Background(), client)
	if err != nil {
		return nil, errors.CreateError(500, err.Error())
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.CreateError(500, err.Error())
	}

	result, err := helper.ExtractTokens(responseBody)
	if err != nil {
		return nil, errors.CreateError(500, err.Error())
	}

	return result, nil
}
