package elasticQuery

import (
	"encoding/json"
	"search-esdb-service/record/entities"
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

func BuildKeywordSearchQuery(keywordSearchEntity *entities.KeywordSearchStruct) (string, string, error) {
	// Build the Elasticsearch query
	queryString := map[string]interface{}{
		"from": keywordSearchEntity.Config.Offset,
		"size": keywordSearchEntity.Config.Amount,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"multi_match": map[string]interface{}{
							"query":  keywordSearchEntity.Query,
							"fields": keywordSearchEntity.KeywordSearchFields,
						},
					},
					{
						"multi_match": map[string]interface{}{
							"query":  keywordSearchEntity.Query,
							"type":   "phrase_prefix",
							"fields": keywordSearchEntity.KeywordSearchFields,
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

func BuildKNNQuery(vectorSearchEntity *entities.VectorSearchStruct) (string, string, error) {
	knnQuery := []map[string]interface{}{}
	for _, model := range vectorSearchEntity.VectorFields {
		knnQuery = append(knnQuery, map[string]interface{}{
			"field":          model.Name + "-question",
			"query_vector":   model.Embedding,
			"k":              5,
			"num_candidates": 50,
			"boost":          model.ScoreWeight,
		})
		knnQuery = append(knnQuery, map[string]interface{}{
			"field":          model.Name + "-answer",
			"query_vector":   model.Embedding,
			"k":              10,
			"num_candidates": 10,
			"boost":          model.ScoreWeight,
		})
	}

	query := map[string]interface{}{
		"from": vectorSearchEntity.Config.Offset,
		"size": vectorSearchEntity.Config.Amount,
		"knn":  knnQuery,
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		return "", "", err
	}

	// Build the count query
	countQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"function_score": map[string]interface{}{
				"query": map[string]interface{}{
					"match_all": map[string]interface{}{}, // Or any other query if needed
				},
				"functions": knnQuery,
			},
		},
	}

	countQueryJSON, err := json.Marshal(countQuery)
	if err != nil {
		return "", "", err
	}

	return string(queryJSON), string(countQueryJSON), nil
}

func BuildHybridSearchQuery(hybridSearchEntity *entities.HybridSearchStruct) (string, string, error) {
	keywordQuery := map[string]interface{}{
		"bool": map[string]interface{}{
			"should": []map[string]interface{}{
				{
					"multi_match": map[string]interface{}{
						"query":  hybridSearchEntity.Query,
						"fields": hybridSearchEntity.KeywordSearchFields,
						"boost":  hybridSearchEntity.KeywordSearchScoreWeight,
					},
				},
				{
					"multi_match": map[string]interface{}{
						"query":  hybridSearchEntity.Query,
						"type":   "phrase_prefix",
						"fields": hybridSearchEntity.KeywordSearchFields,
						"boost":  hybridSearchEntity.KeywordSearchScoreWeight,
					},
				},
			},
		},
	}

	knnQuery := []map[string]interface{}{}
	for _, model := range hybridSearchEntity.VectorFields {
		knnQuery = append(knnQuery, map[string]interface{}{
			"field":          model.Name + "-question",
			"query_vector":   model.Embedding,
			"k":              5,
			"num_candidates": 50,
			"boost":          model.ScoreWeight,
		})
		knnQuery = append(knnQuery, map[string]interface{}{
			"field":          model.Name + "-answer",
			"query_vector":   model.Embedding,
			"k":              10,
			"num_candidates": 10,
			"boost":          model.ScoreWeight,
		})
	}

	queryString := map[string]interface{}{
		"from":  hybridSearchEntity.Config.Offset,
		"size":  hybridSearchEntity.Config.Amount,
		"query": keywordQuery,
		"knn":   knnQuery,
	}

	queryJSON, err := json.Marshal(queryString)
	if err != nil {
		return "", "", err
	}

	// Build the count query
	countQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					keywordQuery,
					{
						"knn": knnQuery,
					},
				},
			},
		},
	}

	countQueryJSON, err := json.Marshal(countQuery)
	if err != nil {
		return "", "", err
	}

	return string(queryJSON), string(countQueryJSON), nil

}
