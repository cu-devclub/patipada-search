package repositories

import (
	"fmt"
	"net/url"
	"search-esdb-service/record/helper"
)

func (r *MLServiceRepository) RemoveStopWordFromQuery(query string) ([]string, error) {
	//TODO : change to gRPC
	baseUrl := "http://localhost:8084/remove-stopWords-from-text"
	params := url.Values{}
	params.Add("query", query)
	fullUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())

	s, err := helper.MakeGETRequest(fullUrl)
	if err != nil {
		return nil, err
	}

	return helper.UnMarshalStringResponse(s)
}

func (r *MLServiceRepository) RemoveStopWordFromTokensArrays(tokens []string) ([]string, error) {
	//TODO : change to gRPC
	baseUrl := "http://localhost:8084/remove-stopWords-from-list"
	params := url.Values{}
	queryString := ""
	for _, token := range tokens {
		queryString += token + ","
	}
	params.Add("query", queryString)
	fullUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())

	s, err := helper.MakeGETRequest(fullUrl)
	if err != nil {
		return nil, err
	}

	return helper.UnMarshalStringResponse(s)
}
