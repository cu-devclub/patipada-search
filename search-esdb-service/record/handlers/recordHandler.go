package handlers

import "github.com/gin-gonic/gin"

type RecordHandler interface {
	// GetAllRecords retrieves all records from the elastic database 
	// and sends a response back to the client.
	//
	// Response:
	// - 200 & A list of all records retrieved from the database.
	// - 500: An internal server error occurred.
	GetAllRecords(c *gin.Context)

	// Search searches for records based on the provided query.
	//
	// It takes a gin.Context object as a parameter.
	// It returns the search results as a slice of records.
	//
	// Query :
	// - query (*required): The query string used to search for records.
	// - amount : The number of results to return. default is 20
	//
	// Response :
	// - 200: The search results.
	// - 400: Bad request. (query not attached) or invalid amount
	// - 500: An internal server error occurred.
	Search(c *gin.Context)

	SearchByRecordIndex(c *gin.Context)
}
