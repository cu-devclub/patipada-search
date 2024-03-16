package database

import (
	"fmt"
	"search-esdb-service/config"
	"search-esdb-service/errors"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/elastic/go-elasticsearch/v8"
)

type elasticDatabase struct {
	Db *elasticsearch.Client
}

// NewElasticDatabase creates a new instance of the ElasticDatabase struct which implements the Database interface.
//
// It takes a pointer to a config.Config struct as a parameter.
// The function initializes a backoff function and an Elasticsearch config struct based on the provided config.
// It creates an Elasticsearch client and checks the health of the Elasticsearch cluster.
// It also checks the plugins and returns an error if there is any.
// Finally, it returns a pointer to an elasticDatabase struct which contains the Elasticsearch client.
func NewElasticDatabase(cfg *config.Config) (Database, error) {
	retryBackoff := backoff.NewExponentialBackOff()
	esCfg := elasticsearch.Config{
		Addresses: []string{cfg.ESDB.URL}, // Elasticsearch cluster URL
		Username:  cfg.ESDB.Username,      // Elasticsearch username
		Password:  cfg.ESDB.Password,      // Elasticsearch password
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
	client, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		return nil, errors.CreateError(500, "Error creating Elasticsearch client")
	}

	es := &elasticDatabase{
		Db: client,
	}

	// Check the Elasticsearch cluster health
	es.CheckClusterHealth()

	//* Check plugins => icu analyzer to extract token
	if err = es.checkPlugins(); err != nil {
		return nil, errors.CreateError(500, fmt.Sprintf("Error checking Elasticsearch plugins: %v", err.Error()))
	}

	return &elasticDatabase{
		Db: client,
	}, nil
}

func (es elasticDatabase) GetDB() *elasticsearch.Client {
	return es.Db
}
