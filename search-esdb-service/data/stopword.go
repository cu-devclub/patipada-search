package data

import (
	"bufio"
	"os"
	"search-esdb-service/config"
)

type StopWords struct {
	PythaiNLP []string
}

// read pythainlp-corpus-stopwords_th.txt and return the content
func retrieveStopWordFromFile(cfg *config.Config) (*StopWords, error) {
	stopWordPath := cfg.Static.DataPath + cfg.Static.StopWordPath
	file, err := os.Open(stopWordPath + "/pythainlp-corpus-stopwords_th.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	stopWords := &StopWords{
		PythaiNLP: lines,
	}
	return stopWords, nil
}

func (d *DataImpl) GetStopWord() StopWords {
	return *d.StopWords
}
