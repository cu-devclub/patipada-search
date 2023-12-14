package main

import (
	"fmt"
	"search-esdb-service/config"
	"search-esdb-service/database"
	recordMigrator "search-esdb-service/record/migration"
	"search-esdb-service/server"
)

func main() {
	cfg := config.GetConfig()
	fmt.Println("CONFIG--------",cfg)
	db := database.NewElasticDatabase(&cfg)
	recordMigrator.RecordMigrate(&cfg, db)
	server.NewGinServer(&cfg, db.GetDB()).Start()
}
