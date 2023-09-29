package es

import (
	"context"
	"fmt"
	"log"
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
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	// Check the Elasticsearch cluster health
	checkClusterHealth(client)

	if err = createIndex(client, "record"); err != nil {
		log.Fatalf("Error creating index: %s", err)
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
		log.Fatalf("Error checking cluster health: %s", err)
	}
	defer res.Body.Close()

	// Check the response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.Status())
	}

	// Print the cluster health information
	fmt.Println("Cluster Health:")
	fmt.Println("---------------")
	fmt.Printf("Status: %s\n", res.Status())

}
