package migration

import (
	"log"
	"log/slog"
	"search-esdb-service/config"
	"search-esdb-service/database"
	recordRepository "search-esdb-service/record/repositories/recordRepository"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

// const (
// 	IndexCreationBody = `
// {
//   "settings": {
// 	"index": {
// 	  "analysis": {
// 		"analyzer": {
// 		  "analyzer_shingle": {
// 			"tokenizer": "icu_tokenizer",
// 			"filter": ["filter_shingle"]
// 		  }
// 		},
// 		"filter": {
// 		  "filter_shingle": {
// 			"type": "shingle",
// 			"max_shingle_size": 3,
// 			"min_shingle_size": 2,
// 			"output_unigrams": true
// 		  }
// 		}
// 	  }
// 	}
//   },
// 	"mappings": {
// 	"properties": {
// 	  "youtubeURL": {
// 		"type": "text"
// 	  },
// 	  "question": {
// 		"type": "text",
// 		"analyzer": "analyzer_shingle"
// 	  },
// 	  "question_lda":{
// 		"type": "dense_vector",
// 		"dims": 5
// 	  },
// 	  "answer": {
// 		"type": "text",
// 		"analyzer": "analyzer_shingle"
// 	  },
// 	  "answer_lda":{
// 		"type": "dense_vector",
// 		"dims": 5
// 	  },
// 	  "startTime": {
// 		"type": "text"
// 	  },
// 	  "endTime": {
// 		"type": "text"
// 	  }
// 	}
//   }
// }
// `
// )

func MigrateRecords(cfg *config.Config, es database.Database) error {
	client := es.GetDB()
	indexName := "record"
	exists, err := doesIndexExist(client, indexName)
	if err != nil {
		panic(err)
	}
	if exists {
		slog.Info("-----Index already exists, no need to migrate ------")
		return nil
	}

	indexCreationQuery, err := CreateIndexCreationBody(&cfg.MlConfig)
	if err != nil {
		return err
	}

	log.Println("Creating index with query: ", indexCreationQuery)

	_, err = client.Indices.Create(
		indexName,
		client.Indices.Create.WithBody(strings.NewReader(indexCreationQuery)),
	)
	if err != nil {
		return err
	}

	records, err := ConvertRawCSVDatatoRecords(cfg)
	if err != nil {
		return err
	}

	err = VectorMigration(records, cfg)
	if err != nil {
		return err
	}

	recordESRepository := recordRepository.NewRecordESRepository(es.GetDB())
	if err := recordESRepository.BulkInsert(records); err != nil {
		return err
	}

	return nil
}

func doesIndexExist(client *elasticsearch.Client, indexName string) (bool, error) {
	res, err := client.Indices.Exists([]string{indexName})
	if err != nil {
		return false, err
	}
	return res.StatusCode != 404, nil
}
