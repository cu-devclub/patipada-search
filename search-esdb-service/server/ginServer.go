package server

import (
	"fmt"
	"search-esdb-service/config"
	recordHandlers "search-esdb-service/record/handlers"
	recordRepository "search-esdb-service/record/repositories"
	recordUsecases "search-esdb-service/record/usecases"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

func (g *ginServer) GetDB() *elasticsearch.Client {
	return g.db
}

func (g *ginServer) Start() {

	// Allow CORS from frontend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{g.cfg.App.FrontendURL,"http://localhost:5173"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	g.app.Use(cors.New(config))

	g.initializeRecordHttpHandler()

	g.app.Run(fmt.Sprintf(":%d", g.cfg.App.Port))
}

// initializeRecordHttpHandler initializes the HTTP handlers for the record API.
//
// It creates the necessary dependencies, such as the record repository,
// usecase, and HTTP handler. It registers the handlers for the different
// endpoints of the record API, such as "/displayAllRecords" and "/search".
func (g *ginServer) initializeRecordHttpHandler() {
	recordESRepository := recordRepository.NewRecordESRepository(g.db)

	recordUsecase := recordUsecases.NewRecordUsecase(recordESRepository)

	recordHttpHandler := recordHandlers.NewRecordHttpHandler(recordUsecase)

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
	// - amount : The number of results to return. default is 20
	//
	// Response :
	// - 200 & The search results.
	// - 400: Bad request. (query not attached) or invalid amount
	// - 500: An internal server error occurred.
	g.app.GET("/search", recordHttpHandler.Search)

	g.app.GET("/search/:recordIndex", recordHttpHandler.SearchByRecordIndex)

}
