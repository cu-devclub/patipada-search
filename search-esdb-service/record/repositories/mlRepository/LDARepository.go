package repositories

import (
	"bytes"
	"encoding/json"
	"net/http"
	"search-esdb-service/errors"
	"search-esdb-service/record/helper"
)

func (r *MLServiceRepository) PerformLDATopicModelling(tokens []string) ([]float64, error) {
	baseUrl := "http://localhost:8084/lda"
	data := map[string][]string{
		"tokens": tokens,
	}

	// Convert the data to JSON
	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.CreateError(http.StatusInternalServerError, err.Error())
	}

	// Create a reader from the JSON body
	bodyReader := bytes.NewReader(body)

	// Make the POST request
	result, err := helper.MakePOSTRequest(baseUrl, bodyReader)
	if err != nil {
		return nil, errors.CreateError(http.StatusInternalServerError, err.Error())
	}

	floatResult, err := helper.UnMarshalFloatResponse(result)
	if err != nil {
		return nil, errors.CreateError(http.StatusInternalServerError, err.Error())
	}

	return floatResult, nil
}
