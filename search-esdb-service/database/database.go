package database 

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type Database interface {
	GetDB() *elasticsearch.Client
}