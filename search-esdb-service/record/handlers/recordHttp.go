package handlers

import (
	"search-esdb-service/messages"
	"search-esdb-service/record/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

type recordHttpHandler struct {
	recordUsecase usecases.RecordUsecase
}

func NewRecordHttpHandler(recordUsecase usecases.RecordUsecase) RecordHandler {
	return &recordHttpHandler{
		recordUsecase: recordUsecase,
	}
}

// GetAllRecords retrieves all records from the database and sends a response back to the client.
//
// Response:
// - 200: A list of all records retrieved from the database.
// - 500: An internal server error occurred.
func (r *recordHttpHandler) GetAllRecords(c *gin.Context) {
	records, err := r.recordUsecase.GetAllRecords("record")
	if err != nil {
		baseResponse(c, 500, messages.INTERNAL_SERVER_ERROR)
		return
	}
	baseResponse(c, 200, records)
}

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
func (r *recordHttpHandler) Search(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		baseResponse(c, 400, messages.BAD_REQUEST)
		return
	}

	sAmount := c.Query("amount")
	amount := 20 // default to 20 results
	if sAmount != "" {
		var err error 
		amount,err = strconv.Atoi(sAmount)
		if err != nil {
			baseResponse(c, 400, messages.BAD_REQUEST)
			return
		}
	}

	records, err := r.recordUsecase.Search("record", query,amount)
	if err != nil {
		baseResponse(c, 500, messages.INTERNAL_SERVER_ERROR)
		return
	}
	baseResponse(c, 200, records)

}
