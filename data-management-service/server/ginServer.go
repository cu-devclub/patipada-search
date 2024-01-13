package server

import (
	"data-management/communication"
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
	comm      communication.Communication
}

func NewGinServer(cfg *config.Config, db *database.Database, v *validator.Validator, c communication.Communication) Server {
	return &ginServer{
		app:       gin.Default(),
		db:        db,
		cfg:       cfg,
		validator: v,
		comm:      c,
	}
}

func (g *ginServer) Start() {
	// Allow CORS from frontend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{g.cfg.App.FrontendURL, "http://localhost:5173"}
	g.app.Use(cors.New(config))
	g.initializeRequestHttpHandler()
	g.app.Run(fmt.Sprintf(":%d", g.cfg.App.Port))
}

func (g *ginServer) initializeRequestHttpHandler() {
	database := *g.db
	requestRepositories := repositories.NewRequestRepositories(database.GetDb(), g.comm)

	requestUsecase := usecases.NewRequestUsecase(requestRepositories, *g.validator)

	requestHandlers := handlers.NewRequestHandler(requestUsecase)

	userRoutes := g.app.Group("/")
	userRoutes.Use(g.AuthMiddleware("user"))

	// GET /requests route is used to retrieve requests based on the provided query parameters:
	// status, username, requestID, index, and approvedBy.
	// If a query parameter is an empty string, it will not be included in the filter.
	// The function responds with a JSON object that includes the matching requests.
	// If an error occurs during the operation, the function responds with a JSON object that includes the error message and status code.
	// Possible status codes are:
	//   200: The operation was successful. The response body contains the matching requests.
	//   400: Bad Request. The request was invalid or cannot be served. The exact error is provided in the response.
	//   500: Internal Server Error. The server encountered an unexpected condition which prevented it from fulfilling the request.
	userRoutes.GET("/requests", requestHandlers.GetRequest)

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
	userRoutes.POST("/requests", requestHandlers.InsertRequest)

	authRoutes := g.app.Group("/")
	authRoutes.Use(g.AuthMiddleware("admin"))

	// PUT /request route is used to update a request.
	//
	// It uses the UpdateRequest handler function which updates a request using the gin context.
	// The function first tries to bind the JSON body of the request to a models.Request struct.
	// The models.Request struct has the following fields:
	//
	//  ID : The ID of the request. It is a string and is required.
	//  RequestID:  The ID of the request. It is a string and is required.
	//	Index:      The index of the request. It is a string and is required.
	//	YoutubeURL: The URL of the YouTube video for the request. It is a string and is required.
	//	Question:   The question of the request. It is a string and is required.
	//	Answer:     The answer of the request. It is a string and is required.
	//	StartTime:  The start time of the request in the YouTube video. It is a string and is required.
	//	EndTime:    The end time of the request in the YouTube video. It is a string and is required.
	//	CreatedAt:  The creation time of the request. It is a time.Time and is optional.
	//	UpdatedAt:  The update time of the request. It is a time.Time and is optional.
	//	By: 	   The user who created the request. It is a string.
	//  ApprovedBy: The user who approved the request. It is a string.
	//  Status:     The status of the request. It is a string.
	//
	// Responses:
	// - 200 OK: The request was successfully updated. The response body contains the updated request in JSON format.
	// - 400 Bad Request: The request body could not be bound to a models.Request struct or some fields are invalid.
	// - 500 Internal Server Error: An internal server error occurred.
	//
	// Usage: authRoutes.PUT("/request", requestHandlers.UpdateRequest)
	authRoutes.PUT("/request", requestHandlers.UpdateRequest)
}
