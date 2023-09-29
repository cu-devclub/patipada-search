package main

import (
	"fmt"
	// "search-esdb-service/internal/csv"
	"search-esdb-service/internal/csv"
	"search-esdb-service/internal/es"
	"search-esdb-service/internal/router"
	"search-esdb-service/internal/util"

	"github.com/spf13/viper"
)

func main() {
	util.InitViper()
	es.InitESDB()
	csv.ConvertCSVFilesInDirectory(viper.GetString("static.data"))
	r := router.RouterEngine()

	r.Run(fmt.Sprintf(":%d", viper.GetInt("connection.appPort")))
}
