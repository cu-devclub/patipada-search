package migration

import (
	"data-management/config"
	"data-management/database"
	"data-management/request/entities"
	"data-management/server"
	"encoding/csv"
	"os"
)

func Migration(cfg *config.Config, db *database.Database, serv *server.Server) error {
	// TODO : Implement the migration of request
	// Make a queue message to search for the request that need to be updated

	// ---- Migrate Record data ----
	s := *serv
	requestArch := s.GetRequestArch()

	rawDataPath := cfg.App.DataSourcePath + "/raw-data.csv"
	recordsAmount, youtubeAmount, err := processRawRecordFile(rawDataPath)
	if err != nil {
		return err
	}

	recordCounter := &entities.RecordCounter{
		RecordAmount:      recordsAmount,
		YoutubeClipAmount: youtubeAmount,
	}
	if err := requestArch.Repo.UpsertRecordCounter(recordCounter); err != nil {
		return err
	}

	return nil
}

func processRawRecordFile(filePath string) (int, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return 0, 0, err
	}

	// Create a map to store unique YouTube URLs
	urls := make(map[string]bool)

	// Iterate over the lines, starting from the second line (ignoring the header)
	for i := 1; i < len(lines); i++ {
		// The YouTube URL is in the 6th column (index 5)
		url := lines[i][5]
		urls[url] = true
	}

	// The number of lines is the length of the lines slice minus 1 (for the header)
	numLines := len(lines) - 1

	// The number of unique URLs is the size of the urls map
	numUniqueURLs := len(urls)

	return numLines, numUniqueURLs, nil
}
