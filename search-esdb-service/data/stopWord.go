package data

import (
	"os"
	"path/filepath"
	"search-esdb-service/config"
	"strings"
)

var stopWord []string

func ReadStopWord(cfg *config.Config) error {
	// read .txt file in /datasource/stopWord
	// and store it in stopWord
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	directoryPath := cfg.Static.DataPath + cfg.Static.StopwordPath

	dataDirPath := filepath.Join(cwd, directoryPath)

	dir, err := os.ReadDir(dataDirPath)
	if err != nil {
		return err
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(dataDirPath, file.Name())
		// read file
		f, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		// split by \n
		words := strings.Split(string(f), "\n")
		for _, word := range words {
			// append to stopWord
			stopWord = append(stopWord, string(word))
		}
	}

	return nil
}

func GetStopWord() []string {
	return stopWord
}