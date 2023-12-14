package helper

import (
	"encoding/json"
	"search-esdb-service/record/entities"
)

// ExtractTokens extracts tokens from the analyze response JSON
// and returns them as an array of strings.
func ExtractTokens(responseJSON []byte) ([]string, error) {
	var analyzeResponse struct {
		Tokens []entities.Token `json:"tokens"`
	}

	// Unmarshal the JSON response into the analyzeResponse struct
	if err := json.Unmarshal(responseJSON, &analyzeResponse); err != nil {
		return nil, err
	}

	// Extract tokens from the struct
	var tokens []string
	for _, token := range analyzeResponse.Tokens {
		tokens = append(tokens, token.Token)
	}

	return tokens, nil
}
