package es

import (
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func createRecordIndex(client *elasticsearch.Client, indexName string) error {
	// Check if the index already exists
	exists, err := indexExists(client, indexName)
	if err != nil {
		return err
	}

	if !exists {
		index := indexName
    // TODO : change tokenizer 
        mapping := `
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


		res, err := client.Indices.Create(
			index,
			client.Indices.Create.WithBody(strings.NewReader(mapping)),
		)
		if err != nil {
			return err
		}
        log.Print(res)
	}

	return nil
}



func indexExists(client *elasticsearch.Client, indexName string) (bool, error) {
	// Check if the index exists using the Indices.Exists API
	res, err := client.Indices.Exists([]string{indexName})
	if err != nil {
		return false, err
	}

	return res.StatusCode != 404, nil
}
