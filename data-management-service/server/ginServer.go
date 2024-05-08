package server

import (
	"data-management/communication"
	"data-management/config"
	"data-management/database"
	requestHandlers "data-management/request/handlers"
	requestRepositories "data-management/request/repositories"
	requestUsecases "data-management/request/usecases"

	ratingHandlers "data-management/ratings/handlers"
	ratingRepositories "data-management/ratings/repositories"
	ratingUsecases "data-management/ratings/usecases"
	validator "data-management/structValidator"

	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ginServer struct {
	app         *gin.Engine
	db          *database.Database
	cfg         *config.Config
	validator   *validator.Validator
	comm        communication.Communication
	requestArch *RequestArch
}

type RequestArch struct {
	Repo    requestRepositories.Repositories
	Usecase requestUsecases.UseCase
	Handler requestHandlers.Handlers
}

func NewGinServer(cfg *config.Config, db *database.Database, v *validator.Validator, c *communication.Communication) Server {
	g := gin.New()
	// Allow CORS from frontend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{cfg.App.FrontendURL, "http://localhost:5173"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	g.Use(cors.New(config))

	serv := &ginServer{
		app:       g,
		db:        db,
		cfg:       cfg,
		validator: v,
		comm:      *c,
	}

	serv.initializeRequestHttpHandler()

	return serv
}

func (g *ginServer) GetRequestArch() *RequestArch {
	return g.requestArch
}

func (g *ginServer) Start() {
	g.app.Run(fmt.Sprintf(":%d", g.cfg.App.Port))
}

func (g *ginServer) initializeRequestHttpHandler() {
	database := *g.db
	requestRepository := requestRepositories.NewRequestRepositories(database.GetDb(), &g.comm)

	requestUsecase := requestUsecases.NewRequestUsecase(&requestRepository, g.validator)

	requestHandler := requestHandlers.NewRequestHandler(&requestUsecase)

	g.requestArch = &RequestArch{
		Repo:    requestRepository,
		Usecase: requestUsecase,
		Handler: requestHandler,
	}

	g.initializedRequestUserRoutes(requestHandler)
	g.initializedRequestAdminRoutes(requestHandler)

	ratingRepository := ratingRepositories.NewRatingRepository(database.GetDb())
	ratingUsecase := ratingUsecases.NewRatingUsecase(&ratingRepository)
	ratingHandler := ratingHandlers.NewRatingHandler(&ratingUsecase)

	g.initializedRatingRoutes(ratingHandler)
}

func (g *ginServer) initializedRequestUserRoutes(handler requestHandlers.Handlers) {
	userRoutes := g.app.Group("/")
	userRoutes.Use(g.AuthMiddleware("user"))

	// GET /requests route is used to retrieve requests based on the provided
	// Query parameters:
	// status, username, requestID, index, and approvedBy.
	// If a query parameter is an empty string, it will not be included in the filter.
	// The function responds with a JSON object that includes the matching requests.
	// If an error occurs during the operation, the function responds with a JSON object that includes the error message and status code.
	//
	// Response , Possible status codes are:
	//   200: The operation was successful. The response body contains the matching requests.
	//   400: Bad Request. The request was invalid or cannot be served. The exact error is provided in the response.
	//   500: Internal Server Error. The server encountered an unexpected condition which prevented it from fulfilling the request.
	userRoutes.GET("/requests", handler.GetRequest)

	// GetLastestRequestOfRecord is a handler function for the GET /request/latest endpoint.
	// Query Parameters:
	// 	- index: The index of the record.
	// It retrieves the latest request of a record based on the provided index query parameter.
	// The function responds with status 200 and a JSON object that includes the latest request.
	// If an error occurs during the operation, the function responds with a JSON object that includes the error message and status code.
	//
	// Possible error status codes are
	// 		400 (Bad Request) and
	// 		500 (Internal Server Error).
	userRoutes.GET("/request/latestRecord", handler.GetLastestRequestOfRecord)

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
	userRoutes.POST("/requests", handler.InsertRequest)
}

func (g *ginServer) initializedRequestAdminRoutes(handler requestHandlers.Handlers) {
	adminRoutes := g.app.Group("/")
	adminRoutes.Use(g.AuthMiddleware("admin"))

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
	//  approvedBy: The user who approved the request. It is a string.
	//  Status:     The status of the request. It is a string.
	//
	// Responses:
	// - 200 OK: The request was successfully updated. The response body contains the updated request in JSON format.
	// - 400 Bad Request: The request body could not be bound to a models.Request struct or some fields are invalid.
	// - 500 Internal Server Error: An internal server error occurred.
	//
	// Usage: adminRoutes.PUT("/request", requestHandlers.UpdateRequest)
	adminRoutes.PUT("/request", handler.UpdateRequest)

	// GET /summary route is used to retrieve the summary of the requests and records.
	adminRoutes.GET("/summary", handler.GetSummary)

	// POST /sync route is used to sync the record in the request to the search.
	adminRoutes.POST("/sync", handler.SyncRequestRecord)

	// POST /sync-all route is used to sync all records in the request to the search.
	adminRoutes.POST("/sync-all", handler.SyncAllRequestRecords)
}

func (g *ginServer) initializedRatingRoutes(handler ratingHandlers.Handlers) {
	ratingRoutes := g.app.Group("/ratings")

	// POST /ratings route is used to insert a new rating into the database.
	// It expects a JSON body that matches the structure of the models.Rating struct.
	//
	//  RatingID: The ID of the rating. It is a string and is optional.
	//  Stars:    The number of stars given in the rating. It is an integer and is required.
	//  Comment:  The comment given in the rating. It is a string and is optional.
	//
	// Responses:
	// - 201 Created: The rating was successfully created. The response body contains the created rating.
	// - 400 Bad Request: The request body could not be bound to a models.Rating struct.
	// - 500 Internal Server Error: An internal server error occurred.
	ratingRoutes.POST("", handler.InsertRating)

	// GET /ratings route is used to retrieve all ratings from the database.
	// The function responds with a JSON object that includes all ratings.
	// If an error occurs during the operation, the function responds with a JSON object that includes the error message and status code.
	//
	// Responses:
	// - 200 OK: The operation was successful. The response body contains all ratings.
	// - 500 Internal Server Error: The server encountered an unexpected condition which prevented it from fulfilling the request.
	ratingRoutes.GET("", handler.GetRatings)

	// GET /ratings/average route is used to retrieve the average rating from the database.
	// The function responds with a JSON object that includes the average rating.
	// If an error occurs during the operation, the function responds with a JSON object that includes the error message and status code.
	//
	// Responses:
	// - 200 OK: The operation was successful. The response body contains the average rating.
	// - 500 Internal Server Error: The server encountered an unexpected condition which prevented it from fulfilling the request.
	ratingRoutes.GET("/average", handler.GetSummaryRatings)
}
