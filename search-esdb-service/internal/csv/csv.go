package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"search-esdb-service/internal/dto"
	"search-esdb-service/internal/es"
	"search-esdb-service/internal/util"
	"strings"
)

func ConvertCSVFilesInDirectory(directoryPath string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	dataDirPath := filepath.Join(cwd, directoryPath)

	dir, err := os.ReadDir(dataDirPath)
	if err != nil {
		return err
	}
	var qaRecords []*dto.QARecord

	for _, entry := range dir {
		// Check if the entry is a regular file and has a .csv extension
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".csv") {
			continue
		}

		// Build the full path to the CSV file
		csvFilePath := filepath.Join(dataDirPath, entry.Name())
		fileName := strings.TrimSuffix(entry.Name(), ".csv")

		// Insert data from the CSV file
		qa, err := generateDataFromCSV(csvFilePath, fileName)
		if err != nil {
			fmt.Printf("Error inserting data from CSV file %s: %s\n", csvFilePath, err)
			continue // Continue to the next file if there's an error
		}
		qaRecords = append(qaRecords, qa...)
	}
	// fmt.Println(len(qaRecords))
	es.BulkInsertQARecords(qaRecords)

	return nil
}

func generateDataFromCSV(filePath string, fileName string) ([]*dto.QARecord, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read and discard the header line
	if _, err := reader.Read(); err != nil {
		return nil, err
	}
	// FOR BULKING
	var qaRecords []*dto.QARecord
	order := 1
	// Read CSV records and insert them into Elasticsearch
	for {
		record, err := reader.Read()
		if err != nil {
			// End of file
			break
		}
		//Empty Record
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

		// Remove newline characters from the fields
		for i := range record {
			record[i] = util.EscapeText(record[i])
		}

		// Escape . to : in record[2] and record[3] (starttime and endtime)
		record[2] = strings.ReplaceAll(record[2], ".", ":")
		record[3] = strings.ReplaceAll(record[3], ".", ":")

		// Assuming your CSV columns are in the order: Question, Answe``r, StartTime, EndTime
		qar := &dto.QARecord{
			YoutubeURL: fileName,
			Question:   record[0],
			Answer:     record[1],
			StartTime:  record[2],
			EndTime:    record[3],
		}

		qaRecords = append(qaRecords, qar) // FOR BULKING
		order += 1
	}

	return qaRecords, nil
}
