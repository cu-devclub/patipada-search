package elasticQuery

import (
	"encoding/json"
	"strings"
)

func BuildMatchAllQuery() (string, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		return "", err
	}

	return string(queryJSON), nil
}

func BuildElasticsearchQuery(query string) (string, error) {
	// Build the Elasticsearch query
	queryString := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"multi_match": map[string]interface{}{
							"query":  query,
							"fields": []string{"question", "answer"},
						},
					},
					{
						"multi_match": map[string]interface{}{
							"query":  query,
							"type":   "phrase_prefix",
							"fields": []string{"question", "answer"},
						},
					},
					{
						"term": map[string]interface{}{
							"_id": query,
						},
					},
				},
			},
		},
	}

	// Convert the query to JSON
	queryJSON, err := json.Marshal(queryString)
	if err != nil {
		return "", err
	}

	return string(queryJSON), nil
}

func BuildElasticsearchQueryByTokens(query string) (string, error) {
	tokens := strings.Split(query, " ")
	// Build the Elasticsearch query
	queryString := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{},
			},
		},
	}

	// Add a match query for each token
	for _, token := range tokens {
		// First multi_match query
		matchQuery := map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  token,
				"fields": []string{"question", "answer"},
			},
		}
		queryString["query"].(map[string]interface{})["bool"].(map[string]interface{})["should"] = append(queryString["query"].(map[string]interface{})["bool"].(map[string]interface{})["should"].([]map[string]interface{}), matchQuery)

		// Second multi_match query with type phrase_prefix
		matchQueryPhrasePrefix := map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  token,
				"type":   "phrase_prefix",
				"fields": []string{"question", "answer"},
			},
		}
		queryString["query"].(map[string]interface{})["bool"].(map[string]interface{})["should"] = append(queryString["query"].(map[string]interface{})["bool"].(map[string]interface{})["should"].([]map[string]interface{}), matchQueryPhrasePrefix)
	}

	// Convert the query to JSON
	queryJSON, err := json.Marshal(queryString)
	if err != nil {
		return "", err
	}

	return string(queryJSON), nil
}
