// TEMPORARY
package helper

import (
	"io"
	"net/http"
	"search-esdb-service/errors"
)

type StringResponse struct {
	Result []string `json:"result"`
}

type FloatResponse struct {
	Result []float64 `json:"result"`
}

func MakeGETRequest(fullUrl string) ([]byte, error) {
	// Make the request
	resp, err := http.Get(fullUrl)
	if err != nil {
		return nil, errors.CreateError(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.CreateError(http.StatusInternalServerError, err.Error())
	}

	return body, nil
}

func MakePOSTRequest(fullUrl string, body io.Reader) ([]byte, error) {
	// Make the request
	resp, err := http.Post(fullUrl, "application/json", body)
	if err != nil {
		return nil, errors.CreateError(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	// Read the response body
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.CreateError(http.StatusInternalServerError, err.Error())
	}
	return resBody, nil
}
