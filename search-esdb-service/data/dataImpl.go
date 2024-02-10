package data

import (
	"log"
	"search-esdb-service/config"
)

type DataImpl struct {
	StopWords *StopWords
}

func NewData(cfg *config.Config) Data {
	sw, err := retrieveStopWordFromFile(cfg)
	if err != nil {
		log.Fatal("Error getting stopWord:", err)
	}

	return &DataImpl{
		StopWords: sw,
	}
}
