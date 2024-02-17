package handlers

import (
	"log"
	"search-esdb-service/constant"
	"search-esdb-service/messages"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *recordHttpHandler) Search(c *gin.Context) {
	log.Println("Search Handler ...")
	// retrieve query
	query := c.Query("query")
	if query == "" {
		errorResponse(c, 400, messages.BAD_REQUEST, messages.QUERY_PARAMETER_EMPTY)
		return
	}

	// retrieve amount
	sAmount := c.Query("amount")
	amount := 50 // default to 50 results
	if sAmount != "" {
		var err error
		amount, err = strconv.Atoi(sAmount)
		if err != nil {
			errorResponse(c, 400, messages.BAD_REQUEST, messages.AMOUNT_INSUFFICENT)
			return
		}
	}

	// retreive search type
	searchType := c.Query("searchType")
	if searchType == "" {
		searchType = constant.SEARCH_BY_DEFAULT
	}

	log.Println("With query:", query, "amount:", amount, "searchType:", searchType)

	// search for records
	records, err := r.recordUsecase.Search("record", query, searchType, amount)
	if err != nil {
		errorResponse(c, 500, messages.INTERNAL_SERVER_ERROR, err.Error())
		return
	}
	successResponse(c, 200, records)

}

func (r *recordHttpHandler) SearchByRecordIndex(c *gin.Context) {
	log.Println("SearchByRecordIndex Handler ...")
	// retrieve query
	recordIndex := c.Param("recordIndex")
	if recordIndex == "" {
		errorResponse(c, 400, messages.BAD_REQUEST, messages.QUERY_PARAMETER_EMPTY)
		return
	}

	log.Println("With recordIndex:", recordIndex)

	record, err := r.recordUsecase.SearchByRecordIndex("record", recordIndex)
	if err != nil {
		errorResponse(c, 500, messages.INTERNAL_SERVER_ERROR, err.Error())
		return
	}
	if record == nil {
		errorResponse(c, 404, messages.NOT_FOUND, messages.RECORD_INDEX_NOT_FOUND)
		return
	}
	successResponse(c, 200, record)
}

// GetAllRecords retrieves all records from the elastic database
// and sends a response back to the client.
//
// Response:
// - 200 & A list of all records retrieved from the database.
// - 500: An internal server error occurred.
func (r *recordHttpHandler) GetAllRecords(c *gin.Context) {
	records, err := r.recordUsecase.GetAllRecords("record")
	if err != nil {
		errorResponse(c, 500, messages.INTERNAL_SERVER_ERROR, err.Error())
		return
	}
	successResponse(c, 200, records)
}
