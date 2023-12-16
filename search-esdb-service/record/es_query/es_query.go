// Use to store and generate query word for elastic db 

package es_query

import (
	"encoding/json"
	"fmt"
)

const (
	CREATE_INDEX_ICU_TOKENIZER = `
{
  "settings": {
    "index": {
      "analysis": {
        "analyzer": {
          "analyzer_shingle": {
            "tokenizer": "icu_tokenizer",
            "filter": ["filter_shingle"]
          }
        },
        "filter": {
          "filter_shingle": {
            "type": "shingle",
            "max_shingle_size": 3,
            "min_shingle_size": 2,
            "output_unigrams": "true"
          }
        }
      }
    }
  },
    "mappings": {
    "properties": {
      "youtubeURL": {
        "type": "text"
      },
      "question": {
        "type": "text",
        "analyzer": "analyzer_shingle"
      },
      "answer": {
        "type": "text",
        "analyzer": "analyzer_shingle"
      },
      "startTime": {
        "type": "text"
      },
      "endTime": {
        "type": "text"
      }
    }
  }
}
`
)

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

func BuildAnalyzeQuery(index, query string) string {
	return fmt.Sprintf(`{
        "tokenizer": "icu_tokenizer",
        "text": "%s"
    }`, query)
}

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