package es

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

func createIndex(client *elasticsearch.Client, indexName string) error {
	// Check if the index already exists
	exists, err := indexExists(client, indexName)
	if err != nil {
		return err
	}

	if !exists {
		// Create the index using the Indices.Create API
		res, err := client.Indices.Create(indexName)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		// Check the response status
		if res.IsError() {
			return fmt.Errorf("Elasticsearch error: %s", res.Status())
		}
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
