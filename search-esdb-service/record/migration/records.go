package migration

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"search-esdb-service/config"
	"search-esdb-service/data"
	"search-esdb-service/record/entities"
	"search-esdb-service/record/helper"
	"strings"
)

func ConvertRawCSVDatatoRecords(cfg *config.Config) ([]*entities.Record, error) {
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
		// fileName := strings.TrimSuffix(entry.Name(), ".csv")

		r, err := generateRecordsFromCSV(csvFilePath)
		if err != nil {
			continue
		}
		records = append(records, r...)
	}

	return records, nil
}

func generateRecordsFromCSV(filePath string) ([]*entities.Record, error) {
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
			Index:      record[4],
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

// func UpdateRecordsWithLDA(cfg *config.Config, records []*entities.Record) ([]*entities.Record, error) {
// 	ldaDirPath := cfg.Static.DataPath + cfg.Static.LDAPath
// 	dir, err := os.ReadDir(ldaDirPath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	recordMap := make(map[string]*entities.Record)
// 	for _, record := range records {
// 		recordMap[record.Index] = record
// 	}

// 	for _, entry := range dir {
// 		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".csv") {
// 			continue
// 		}

// 		csvFilePath := filepath.Join(ldaDirPath, entry.Name())
// 		file, err := os.Open(csvFilePath)
// 		if err != nil {
// 			continue
// 		}
// 		defer file.Close()

// 		reader := csv.NewReader(file)
// 		for {
// 			record, err := reader.Read()
// 			if err != nil {
// 				break
// 			}

// 			if recordEntity, ok := recordMap[record[0]]; ok {
// 				questionLDA, err := util.ConvertStringToFloat64Arrays(record[1])
// 				if err != nil {
// 					return nil, err
// 				}
// 				if questionLDA != nil {
// 					recordEntity.QuestionLDA = questionLDA
// 				}

// 				answerLDA, err := util.ConvertStringToFloat64Arrays(record[2])
// 				if err != nil {
// 					return nil, err
// 				}
// 				if answerLDA != nil {
// 					recordEntity.AnswerLDA = answerLDA
// 				}

// 			}
// 		}
// 	}

// 	return records, nil
// }
