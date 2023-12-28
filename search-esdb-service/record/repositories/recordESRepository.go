package repositories

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type RecordESRepository struct {
	es *elasticsearch.Client
}

func NewRecordESRepository(es *elasticsearch.Client) RecordRepository {
	return &RecordESRepository{
		es: es,
	}
}
