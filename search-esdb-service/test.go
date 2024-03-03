package main

// import (
// 	"log"
// 	"search-esdb-service/config"
// 	"search-esdb-service/record/migration"
// )

// func main() {
// 	log.Println("Initializing config...")
// 	config.InitializeViper("./")

// 	cfg := config.GetConfig()
// 	log.Println("Config initialized")

// 	records,err := migration.ConvertCSVToRecords(&cfg)
// 	if err != nil {
// 		panic(err)
// 	}

// 	res, err := migration.UpdateRecordsWithLDA(&cfg,records)
// 	if err != nil {
// 		panic(err)
// 	}

// 	log.Println("Records updated with LDA:",res[0].ToString())
// }
