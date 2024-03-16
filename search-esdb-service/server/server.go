package server

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type Server interface {
	Start()
	GetDB() *elasticsearch.Client
	GetRecordArch() *RecordArch
}
