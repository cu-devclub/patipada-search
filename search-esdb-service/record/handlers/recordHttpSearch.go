package handlers

import (
	"search-esdb-service/messages"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Search searches for records based on the provided query.
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
	// retrieve query
	query := c.Query("query")
	if query == "" {
		baseResponse(c, 400, messages.BAD_REQUEST)
		return
	}

	// retrieve amount
	sAmount := c.Query("amount")
	amount := 20 // default to 20 results
	if sAmount != "" {
		var err error
		amount, err = strconv.Atoi(sAmount)
		if err != nil {
			baseResponse(c, 400, messages.BAD_REQUEST)
			return
		}
	}

	records, err := r.recordUsecase.Search("record", query, amount)
	if err != nil {
		baseResponse(c, 500, messages.INTERNAL_SERVER_ERROR)
		return
	}
	baseResponse(c, 200, records)

}
