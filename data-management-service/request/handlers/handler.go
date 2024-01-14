package handlers

import "github.com/gin-gonic/gin"

type Handlers interface {
	// InsertRequest is a HTTP handler function that inserts a new Request into the database.
	// It takes a *gin.Context as an argument, which contains the HTTP request context.
	//
	// The function first tries to bind the JSON body of the request to a models.Request struct.
	// The models.Request struct has the following fields:
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
	// If the binding fails, it responds with a 400 Bad Request status and returns.
	//
	// Then, it calls the InsertRequest method of the requestUsecase to insert the request into the database.
	// If the insertion fails, it responds with the status code and error message from the error returned by InsertRecord, and returns.
	//
	// Error Status Codes:
	//
	//	400: ERR_REQUEST_INDEX_NOT_EXISTS
	//	500: INTERNAL_SERVER_ERROR
	//
	// If the insertion is successful,
	// Response
	//
	//	201: Created status and the inserted request.
	//
	// Usage:
	//
	//	router.POST("/requests", requestHandler.InsertRequest)
	//
	// This will create a new route that accepts POST requests at /requests.
	// The body of the request should be a JSON object that matches the structure of the models.Request struct.
	InsertRequest(c *gin.Context)

	// UpdateRequest updates a request using the gin context.
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
	// Success
	// - 200 OK if success and return the updated request in JSON format
	//
	// Error
	// - 400 Bad Request if the binding fails or some fields are error
	// - 500 Internal Server Error if internal server error
	UpdateRequest(c *gin.Context)

	// GetRequest is a handler function for the GET /request endpoint.
	// It retrieves requests based on the provided query parameters: status, username, requestID, index, and approvedBy.
	// If a query parameter is an empty string, it will not be included in the filter.
	// The function responds with a JSON object that includes the matching requests.
	// If an error occurs during the operation, the function responds with a JSON object that includes the error message and status code.
	//  error status codes
	// - 400 (Bad Request) and
	// - 500 (Internal Server Error).
	GetRequest(c *gin.Context)

	// GetLastestRequestOfRecord is a handler function for the GET /request/latest endpoint.
	// Query Parameters: 
	// 	- index: The index of the record.
	// It retrieves the latest request of a record based on the provided index query parameter.
	// The function responds with a JSON object that includes the latest request.
	// If an error occurs during the operation, the function responds with a JSON object that includes the error message and status code.
	// 
	// Possible error status codes are
	// 		400 (Bad Request) and
	// 		500 (Internal Server Error).
	GetLastestRequestOfRecord(c *gin.Context)
}
