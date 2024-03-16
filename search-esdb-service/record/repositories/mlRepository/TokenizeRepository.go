package repositories

import (
	"fmt"
	"net/url"
	"search-esdb-service/record/helper"
)

func (r *MLServiceRepository) TokenizeQuery(query string) ([]string, error) {
	//TODO : change to gRPC
	baseUrl := "http://localhost:8084/tokenize"
	params := url.Values{}
	params.Add("query", query)
	fullUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())

	s, err := helper.MakeGETRequest(fullUrl)
	if err != nil {
		return nil, err
	}

	return helper.UnMarshalStringResponse(s)
}
