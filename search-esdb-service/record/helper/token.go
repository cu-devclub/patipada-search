package helper

import (
	"encoding/json"
	"log"
	"search-esdb-service/record/entities"
	"search-esdb-service/util"
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

func RemoveStopWordsFromTokensArray(stopWords []string, tokens []string) []string {
	var result []string
	for _, token := range tokens {
		if !util.Contains(stopWords, token) {
			result = append(result, token)
		} else {
			log.Println("Stop word removed: ", token)
		}
	}

	// if the tokens contain only stop words, do not remove any token
	if len(result) == 0 {
		return tokens
	}

	return result
}
