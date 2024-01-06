package server

import (
	"data-management/config"
	"data-management/database"
	"data-management/request/handlers"
	"data-management/request/repositories"
	"data-management/request/usecases"
	validator "data-management/structValidator"

	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ginServer struct {
	app       *gin.Engine
	db        *database.Database
	cfg       *config.Config
	validator *validator.Validator
}

func NewGinServer(cfg *config.Config, db *database.Database, v *validator.Validator) Server {
	return &ginServer{
		app:       gin.Default(),
		db:        db,
		cfg:       cfg,
		validator: v,
	}
}

func (g *ginServer) Start() {
	// Allow CORS from frontend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{g.cfg.App.FrontendURL}
	g.app.Use(cors.New(config))

	g.initializeRequestHttpHandler()

	g.app.Run(fmt.Sprintf(":%d", g.cfg.App.Port))
}

func (g *ginServer) initializeRequestHttpHandler() {
	database := *g.db
	requestRepositories := repositories.NewRequestRepositories(database.GetDb())

	requestUsecase := usecases.NewRequestUsecase(requestRepositories, *g.validator)

	requestHandlers := handlers.NewRequestHandler(requestUsecase)

	// POST /requests is a route that inserts a new request into the database.
	// It expects a JSON body that matches the structure of the models.Request struct.
	//
	//	Index:      The index of the request. It is a string and is required.
	//	YoutubeURL: The URL of the YouTube video for the request. It is a string and is required.
	//	Question:   The question of the request. It is a string and is required.
	//	Answer:     The answer of the request. It is a string and is required.
	//	StartTime:  The start time of the request in the YouTube video. It is a string and is required.
	//	EndTime:    The end time of the request in the YouTube video. It is a string and is required.
	//	CreatedAt:  The creation time of the request. It is a time.Time and is optional.
	//	UpdatedAt:  The update time of the request. It is a time.Time and is optional.
	// 	By: 	   The user who created the request. It is a string.
	//
	// Responses:
	//     201: The request was successfully created. The response body contains the created request.
	//     400: The request body could not be bound to a models.Request struct, or the request index does not exist.
	//     500: An internal server error occurred.
	g.app.POST("/requests", requestHandlers.InsertRequest)

	// The GET /requests route handles the HTTP request for retrieving all requests.
	// It uses the GetAllRequests method of the requestHandlers to get all requests and sends the response in JSON format.
	//
	// Response :
	// - 200 OK if successful, returning all requests in JSON format.
	// - 500 Internal Server Error if an internal server error occurs.
	g.app.GET("/requests", requestHandlers.GetAllRequests)

	// The GET /requests/:requestID route handles the HTTP request for retrieving a request by its request ID.
	// It uses the GetRequestByRequestID method of the requestHandlers to get the request and sends the response in JSON format.
	//
	// Response :
	// - 200 OK if successful, returning the request in JSON format.
	// - 404 Not Found if no request with the provided request ID is found.
	// - 500 Internal Server Error if an internal server error occurs.
	g.app.GET("/requests/:requestID", requestHandlers.GetRequestByRequestID)

	// PUT /request route is used to update a request.
	//
	// It uses the UpdateRequest handler function which updates a request using the gin context.
	// It first binds the request body to a models.Request struct.
	//
	// Responses:
	// - 200 OK: The request was successfully updated. The response body contains the updated request in JSON format.
	// - 400 Bad Request: The request body could not be bound to a models.Request struct or some fields are invalid.
	// - 500 Internal Server Error: An internal server error occurred.
	//
	// Usage: g.app.PUT("/request", requestHandlers.UpdateRequest)
	g.app.PUT("/request", requestHandlers.UpdateRequest)
}
