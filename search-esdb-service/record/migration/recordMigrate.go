package migration

import (
	"encoding/csv"
	"log/slog"
	"os"
	"path/filepath"
	"search-esdb-service/config"
	"search-esdb-service/data"
	"search-esdb-service/database"
	"search-esdb-service/record/entities"
	"search-esdb-service/record/helper"
	recordRepository "search-esdb-service/record/repositories/recordRepository"
	"search-esdb-service/util"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

const (
	IndexCreationBody = `
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
			"output_unigrams": true
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
	  "question_lda":{
		"type": "dense_vector",
		"dims": 5
	  },
	  "answer": {
		"type": "text",
		"analyzer": "analyzer_shingle"
	  },
	  "answer_lda":{
		"type": "dense_vector",
		"dims": 5
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

	_, err = client.Indices.Create(
		indexName,
		client.Indices.Create.WithBody(strings.NewReader(IndexCreationBody)),
	)
	if err != nil {
		return err
	}

	records, err := ConvertCSVToRecords(cfg)
	if err != nil {
		return err
	}

	records, err = UpdateRecordsWithLDA(cfg, records)
	if err != nil {
		return err
	}

	es.CheckClusterHealth()

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

func ConvertCSVToRecords(cfg *config.Config) ([]*entities.Record, error) {
	dataDirPath := cfg.Static.DataPath + cfg.Static.RecordPath
	dir, err := data.GetRecordCSVFilesEntry(cfg)
	if err != nil {
		return nil, err
	}

	var records []*entities.Record
	for _, entry := range dir {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".csv") {
			continue
		}

		csvFilePath := filepath.Join(dataDirPath, entry.Name())
		fileName := strings.TrimSuffix(entry.Name(), ".csv")

		r, err := generateRecordsFromCSV(csvFilePath, fileName)
		if err != nil {
			continue
		}
		records = append(records, r...)
	}

	return records, nil
}

func generateRecordsFromCSV(filePath string, fileName string) ([]*entities.Record, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	var qaRecords []*entities.Record
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		ch := false
		for i := range record {
			if record[i] == "" {
				ch = true
				break
			}
		}
		if ch {
			continue
		}

		for i := range record {
			record[i] = helper.EscapeText(record[i])
		}

		record[2] = strings.ReplaceAll(record[2], ".", ":")
		record[3] = strings.ReplaceAll(record[3], ".", ":")

		qar := &entities.Record{
			Index:      record[5] + "-" + record[4],
			YoutubeURL: record[5],
			Question:   record[0],
			Answer:     record[1],
			StartTime:  record[2],
			EndTime:    record[3],
		}

		qaRecords = append(qaRecords, qar)
	}

	return qaRecords, nil
}

func UpdateRecordsWithLDA(cfg *config.Config, records []*entities.Record) ([]*entities.Record, error) {
	ldaDirPath := cfg.Static.DataPath + cfg.Static.LDAPath
	dir, err := os.ReadDir(ldaDirPath)
	if err != nil {
		return nil, err
	}

	recordMap := make(map[string]*entities.Record)
	for _, record := range records {
		recordMap[record.Index] = record
	}

	for _, entry := range dir {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".csv") {
			continue
		}

		csvFilePath := filepath.Join(ldaDirPath, entry.Name())
		file, err := os.Open(csvFilePath)
		if err != nil {
			continue
		}
		defer file.Close()

		reader := csv.NewReader(file)
		for {
			record, err := reader.Read()
			if err != nil {
				break
			}

			if recordEntity, ok := recordMap[record[0]]; ok {
				questionLDA, err := util.ConvertStringToFloat64Arrays(record[1])
				if err != nil {
					return nil, err
				}
				if questionLDA != nil {
					recordEntity.QuestionLDA = questionLDA
				}

				answerLDA, err := util.ConvertStringToFloat64Arrays(record[2])
				if err != nil {
					return nil, err
				}
				if answerLDA != nil {
					recordEntity.AnswerLDA = answerLDA
				}

			}
		}
	}

	return records, nil
}
