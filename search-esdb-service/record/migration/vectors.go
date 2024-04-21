package migration

import (
	"encoding/csv"
	"log"
	"os"
	"search-esdb-service/config"
	"search-esdb-service/record/entities"
	"search-esdb-service/util"
)

func VectorMigration(records []*entities.Record, cfg *config.Config) error {
	for _, api := range cfg.MlConfig.APIs {
		// READ CSV FILE
		file, err := os.Open(api.CsvFilePath)
		if err != nil {
			return err
		}
		defer file.Close()

		reader := csv.NewReader(file)
		if _, err := reader.Read(); err != nil {
			return err
		}

		// GENERATE VECTORS
		recordVectors := generateRecordVectors(api.Name, reader)

		// UPDATE RECORDS
		for _, recordVec := range recordVectors {
			for _, record := range records {
				if recordVec.Vectors.RecordIndex == record.Index {
					record.Vectors = append(record.Vectors, recordVec)
					break
				}
			}
		}

	}

	return nil
}

func generateRecordVectors(modelName string, reader *csv.Reader) []*entities.RecordVectors {
	recordVecs := make([]*entities.RecordVectors, 0)
	for {
		recordVecReading, err := reader.Read()
		if err != nil {
			log.Println("Error reading record vectors", err)
			break
		}

		questionVec,err := util.ConvertStringToFloat32Arrays(recordVecReading[1])
		if err != nil {
			log.Println("Error converting question vector", err)
			break
		}
		answerVec,err  := util.ConvertStringToFloat32Arrays(recordVecReading[2])
		if err != nil {
			log.Println("Error converting answer vector", err)
			break
		}

		vec := &entities.Vectors{
			RecordIndex:    recordVecReading[0],
			QuestionVector: questionVec,
			AnswerVector:   answerVec,
		}

		recordVec := &entities.RecordVectors{
			Model:   modelName,
			Vectors: vec,
		}

		recordVecs = append(recordVecs, recordVec)
	}

	return recordVecs
}
