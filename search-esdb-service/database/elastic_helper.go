package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// checkClusterHealth is a function that checks the health of the Elasticsearch cluster.
//
// It takes a client of type *elasticsearch.Client as a parameter.
// The function does not return anything.
func (es *elasticDatabase) CheckClusterHealth() {
	// Create a request to check the cluster health
	req := esapi.ClusterHealthRequest{
		Pretty: true,
	}

	// Perform the request
	res, err := req.Do(context.Background(), es.GetDB())
	if err != nil {
		log.Printf("Error checking cluster health: %s", err)
		return
	}
	defer res.Body.Close()

	// Check the response status
	if res.IsError() {
		log.Printf("Error: %s", res.Status())
		return
	}
}

// checkPlugins checks the installed plugins in Elasticsearch.
//
// It takes a `client` parameter of type `*elasticsearch.Client` which is used to
// perform the request to check the installed plugins.
//
// It returns an error if there is an error performing the request or decoding the
// JSON response. It also returns an error if no plugins are installed.
func (es *elasticDatabase) checkPlugins() error {
	// Create a request to check the installed plugins
	req := esapi.CatPluginsRequest{
		Format: "json", // Use JSON format for the response
	}

	// Perform the request
	res, err := req.Do(context.Background(), es.GetDB())
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check the response status
	if res.IsError() {
		return fmt.Errorf("elasticsearch error: %s", res.Status())
	}

	// Decode the JSON response
	var plugins []map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&plugins); err != nil {
		return err
	}

	if len(plugins) == 0 {
		return fmt.Errorf("no plugins installed")
	}
	// Print the list of installed plugins
	for _, plugin := range plugins {
		fmt.Printf("Name: %s, Component: %s, Version: %s\n", plugin["name"], plugin["component"], plugin["version"])
	}

	return nil
}
