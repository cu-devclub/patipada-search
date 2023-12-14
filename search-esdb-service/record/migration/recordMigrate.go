package migration

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"search-esdb-service/config"
	"search-esdb-service/database"
	"search-esdb-service/record/entities"
	"search-esdb-service/record/helper"
	"search-esdb-service/record/es_query"
	"search-esdb-service/record/repositories"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

// Migration steps 
// Create index named `record`; if not exists else return
// Convert file in csv format to json format from data folder
// insert json to es
// ------------------------------

// RecordMigrate migrates records to Elasticsearch.
//
// Takes a *config.Config and a database.Database as parameters.
// Does not return anything.
func RecordMigrate(cfg *config.Config, es database.Database) {
	fmt.Println("RECORD MIGRATION---------")
	client := es.GetDB()
	indexName := "record"
	exists, err := indexExists(client, indexName)
	if err != nil {
		panic(err)
	}
	if exists {
		fmt.Println("---------DATA ALREADY EXISTS---------")
		return // index already exists
	}

	// Create the index
	res, err := client.Indices.Create(
		indexName,
		client.Indices.Create.WithBody(strings.NewReader(es_query.CREATE_INDEX_ICU_TOKENIZER)),
	)
	if err != nil {
		panic(err)
	}
	log.Print(res)

	// Convert csv file
	records, err := ConvertCSVFilesInDirectory(cfg.Static.DataPath)
	fmt.Println("CONVERTING CSV-----------")
	fmt.Println(records[0])
	if err != nil {
		panic(err)
	}

	// bulk insert records to es
	recordESRepository := repositories.NewRecordESRepository(es.GetDB())
	if err := recordESRepository.BulkInsert(records); err != nil {
		panic(err)
	}

	fmt.Printf("Successfully migrated %d records\n", len(records))
}

// indexExists checks if the index exists using the Indices.Exists API.
//
// It takes a client *elasticsearch.Client and an indexName string as parameters.
// It returns a bool indicating whether the index exists and an error if any.
func indexExists(client *elasticsearch.Client, indexName string) (bool, error) {
	// Check if the index exists using the Indices.Exists API
	res, err := client.Indices.Exists([]string{indexName})
	if err != nil {
		return false, err
	}

	return res.StatusCode != 404, nil
}

// ConvertCSVFilesInDirectory converts CSV files in the specified 
// directory into a slice of entities.Record structs.
//
// It takes a directory path as a parameter and returns a slice of 
// entities.Record structs and an error.
func ConvertCSVFilesInDirectory(directoryPath string) ([]*entities.Record, error) {
	// Find file
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dataDirPath := filepath.Join(cwd, directoryPath)

	dir, err := os.ReadDir(dataDirPath)
	if err != nil {
		return nil, err
	}

	var records []*entities.Record

	// Read Files in directory (in case more than 1 file)
	for _, entry := range dir {
		// Check if the entry is a regular file and has a .csv extension
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".csv") {
			continue
		}

		// Build the full path to the CSV file
		csvFilePath := filepath.Join(dataDirPath, entry.Name())
		fileName := strings.TrimSuffix(entry.Name(), ".csv")

		// Insert data from the CSV file
		r, err := generateDataFromCSV(csvFilePath, fileName)
		if err != nil {
			fmt.Printf("Error inserting data from CSV file %s: %s\n", csvFilePath, err)
			continue // Continue to the next file if there's an error
		}
		records = append(records, r...)
	}

	return records, nil
}

// generateDataFromCSV generates a slice of entities.Record structs from a CSV file.
//
// Parameters:
// - filePath: the path to the CSV file.
// - fileName: the name of the CSV file.
//
// Returns:
// - []*entities.Record: a slice of entities.Record structs representing the CSV records.
// - error: an error if there was a problem reading the CSV file.
func generateDataFromCSV(filePath string, fileName string) ([]*entities.Record, error) {
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
	var qaRecords []*entities.Record
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
			record[i] = helper.EscapeText(record[i])
		}

		// Escape . to : in record[2] and record[3] (starttime and endtime)
		record[2] = strings.ReplaceAll(record[2], ".", ":")
		record[3] = strings.ReplaceAll(record[3], ".", ":")

		// Assuming your CSV columns are in the order: Question, Answe``r, StartTime, EndTime
		qar := &entities.Record{
			Index:      record[4],
			YoutubeURL: record[5],
			Question:   record[0],
			Answer:     record[1],
			StartTime:  record[2],
			EndTime:    record[3],
		}

		qaRecords = append(qaRecords, qar) // FOR BULKING
	}

	return qaRecords, nil
}
