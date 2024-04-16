package elasticQuery

import (
	"encoding/json"
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

func BuildElasticsearchQuery(query string, offset, amount int) (string, string, error) {
	searchFields := []string{"question"}
	// Build the Elasticsearch query
	queryString := map[string]interface{}{
		"from": offset,
		"size": amount,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"multi_match": map[string]interface{}{
							"query":  query,
							"fields": searchFields,
						},
					},
					{
						"multi_match": map[string]interface{}{
							"query":  query,
							"type":   "phrase_prefix",
							"fields": searchFields,
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

	// Build the count query
	countQuery := map[string]interface{}{
		"query": queryString["query"],
	}

	// Convert the query to JSON
	queryJSON, err := json.Marshal(queryString)
	if err != nil {
		return "", "", err
	}

	countQueryJSON, err := json.Marshal(countQuery)
	if err != nil {
		return "", "", err
	}

	return string(queryJSON), string(countQueryJSON), nil
}

func BuildKNNQuery(queryVector []float64, field string, offset, amount int) (string, error) {
	knnQuery := map[string]interface{}{
		"query_vector":   queryVector, // The vector to find neighbors for
		"k":              10,          // Number of nearest neighbors to retrieve
		"field":          field,       // The field to compare with
		"num_candidates": 50,
	}

	//TODO : verify from and size when implementing hybrid search
	query := map[string]interface{}{
		"from": offset,
		"size": amount,
		"knn":  knnQuery,
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		return "", err
	}

	return string(queryJSON), nil
}
