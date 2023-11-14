package es

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/spf13/viper"
)

var es_Client *elasticsearch.Client

func GetESClient() *elasticsearch.Client { return es_Client }

func InitESDB() {
	retryBackoff := backoff.NewExponentialBackOff()

	//* for production
	address := viper.GetString("connection.esdbURL")
	username := viper.GetString("esdb.username")
	password := viper.GetString("esdb.password")
	//* for local development
	// address := "http://localhost:9200"
	// username := "elastic"
	// password := "admin"
	cfg := elasticsearch.Config{
		Addresses: []string{address}, // Elasticsearch cluster URL
		Username:  username,          // Elasticsearch username
		Password:  password,          // Elasticsearch password
		// Retry on 429 TooManyRequests statuses
		//
		RetryOnStatus: []int{502, 503, 504, 429},

		// Configure the backoff function
		//
		RetryBackoff: func(i int) time.Duration {
			if i == 1 {
				retryBackoff.Reset()
			}
			return retryBackoff.NextBackOff()
		},

		// Retry up to 5 attempts
		//
		MaxRetries: 5,
	}

	// Create an Elasticsearch client
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Printf("Error creating Elasticsearch client: %s", err)
		return
	}

	// Check the Elasticsearch cluster health
	checkClusterHealth(client)

	//* Check plugins
	if err = checkPlugins(client); err != nil {
		fmt.Printf("Error checking Plugins: %s", err)
		return
	}

	//* Creating Record Index
	if err = createRecordIndex(client, "record"); err != nil {
		fmt.Printf("Error creating index: %s", err)
		return
	}

	es_Client = client
}

func checkClusterHealth(client *elasticsearch.Client) {
	// Create a request to check the cluster health
	req := esapi.ClusterHealthRequest{
		Pretty: true,
	}

	// Perform the request
	res, err := req.Do(context.Background(), client)
	if err != nil {
		fmt.Printf("Error checking cluster health: %s", err)
		return
	}
	defer res.Body.Close()

	// Check the response status
	if res.IsError() {
		fmt.Printf("Error: %s", res.Status())
		return
	}

	// Print the cluster health information
	fmt.Println("Elastic Cluster Health:")
	fmt.Println("---------------")
	fmt.Printf("Status: %s\n", res.Status())

}

func checkPlugins(client *elasticsearch.Client) error {
	// Create a request to check the installed plugins
	req := esapi.CatPluginsRequest{
		Format: "json", // Use JSON format for the response
	}

	// Perform the request
	res, err := req.Do(context.Background(), client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check the response status
	if res.IsError() {
		return fmt.Errorf("Elasticsearch error: %s", res.Status())
	}

	// Decode the JSON response
	var plugins []map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&plugins); err != nil {
		return err
	}

	if len(plugins) == 0 {
		return fmt.Errorf("No plugins installed")
	}
	// Print the list of installed plugins
	for _, plugin := range plugins {
		fmt.Printf("Name: %s, Component: %s, Version: %s\n", plugin["name"], plugin["component"], plugin["version"])
	}

	return nil
}
