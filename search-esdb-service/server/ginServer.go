package server

import (
	"fmt"
	"search-esdb-service/config"
	recordHandlers "search-esdb-service/record/handlers"
	recordRepository "search-esdb-service/record/repositories"
	recordUsecases "search-esdb-service/record/usecases"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type ginServer struct {
	app *gin.Engine
	db  *elasticsearch.Client
	cfg *config.Config
}

func NewGinServer(cfg *config.Config, db *elasticsearch.Client) Server {
	return &ginServer{
		app: gin.Default(),
		db:  db,
		cfg: cfg,
	}
}

func (g *ginServer) Start() {
	
	// Allow CORS from frontend 
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{g.cfg.App.FrontendURL}
	g.app.Use(cors.New(config))
	
	g.initializeRecordHttpHandler()

	g.app.Run(fmt.Sprintf(":%d", g.cfg.App.Port))
}

// initializeRecordHttpHandler initializes the HTTP handlers for the record API.
//
// It creates the necessary dependencies, such as the record repository,
// usecase, and HTTP handler. It registers the handlers for the different
// endpoints of the record API, such as "/displayAllRecords" and "/search".
//
// The "/displayAllRecords" endpoint handles GET requests and returns a
// status code 200 with records found if any, or an empty list if none are
// found. It returns a status code 500 for internal server errors.
//
// The "/search" endpoint handles GET requests and returns a
// status code 200 with records found if any, or an empty list if none are
// found. And returns a status code 400 (Bad Request) if the query word is 
// not attached.
func (g *ginServer) initializeRecordHttpHandler() {
	recordESRepository := recordRepository.NewRecordESRepository(g.db)

	recordUsecase := recordUsecases.NewRecordUsecase(recordESRepository)

	recordHttpHandler := recordHandlers.NewRecordHttpHandler(recordUsecase)

	// GET Request return 200 with records found (get/search)
	// if not found still return 200 with empty list 
	// 500 for internal server error 
	g.app.GET("/displayAllRecords", recordHttpHandler.GetAllRecords)
	
	// If query word is not attach, return 400 Bad request 
	g.app.GET("/search", recordHttpHandler.Search)

}
