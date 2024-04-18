package server

import (
	"fmt"
	"search-esdb-service/config"
	recordHandlers "search-esdb-service/record/handlers"
	mlRepository "search-esdb-service/record/repositories/mlRepository"
	recordRepository "search-esdb-service/record/repositories/recordRepository"
	recordUsecases "search-esdb-service/record/usecases"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ginServer struct {
	app        *gin.Engine
	db         *elasticsearch.Client
	cfg        *config.Config
	recordArch *RecordArch
}

type RecordArch struct {
	Repo    recordRepository.RecordRepository
	Mlrepo  mlRepository.MLRepository
	Usecase recordUsecases.RecordUsecase
	Handler recordHandlers.RecordHandler
}

func NewGinServer(cfg *config.Config, db *elasticsearch.Client) Server {
	serv := gin.New()

	// Allow CORS from frontend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{cfg.App.FrontendURL, "http://localhost:5173"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	serv.Use(cors.New(config))

	g := &ginServer{
		app: serv,
		db:  db,
		cfg: cfg,
	}

	g.initializeRecordHttpHandler()

	return g
}

func (g *ginServer) GetDB() *elasticsearch.Client {
	return g.db
}

func (g *ginServer) GetRecordArch() *RecordArch {
	return g.recordArch
}

func (g *ginServer) Start() {
	g.app.Run(fmt.Sprintf(":%d", g.cfg.App.Port))
}

// initializeRecordHttpHandler initializes the HTTP handlers for the record API.
//
// It creates the necessary dependencies, such as the record repository,
// usecase, and HTTP handler. It registers the handlers for the different
// endpoints of the record API, such as "/displayAllRecords" and "/search".
func (g *ginServer) initializeRecordHttpHandler() {

	recordESRepository := recordRepository.NewRecordESRepository(g.db)
	mlRepository := mlRepository.NewMLServiceRepository()
	recordUsecase := recordUsecases.NewRecordUsecase(recordESRepository, mlRepository)

	recordHttpHandler := recordHandlers.NewRecordHttpHandler(recordUsecase)

	g.recordArch = &RecordArch{
		Repo:    recordESRepository,
		Mlrepo:  mlRepository,
		Usecase: recordUsecase,
		Handler: recordHttpHandler,
	}

	// GetAllRecords retrieves all records from the elastic database
	// and sends a response back to the client.
	//
	// Response:
	// - 200 & A list of all records retrieved from the database.
	// - 500: An internal server error occurred.
	g.app.GET("/displayAllRecords", recordHttpHandler.GetAllRecords)

	// Search searches for records based on the provided query.
	//
	// Query :
	// - query (*required): The query string used to search for records.
	// - amount : The number of results to return. default is 50
	// - searchType : The type of search to perform. one of "tf-idf" or "default"
	//
	// Response :
	// - 200 & The search results.
	// - 400: Bad request. (query not attached) or invalid amount
	// - 500: An internal server error occurred.
	g.app.GET("/search", recordHttpHandler.Search)

	g.app.GET("/search/:recordIndex", recordHttpHandler.SearchByRecordIndex)

	// Prometheus metrics
	g.app.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
