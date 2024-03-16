package handlers

import (
	"net/http"
	"search-esdb-service/constant"
	"search-esdb-service/errors"
	"search-esdb-service/messages"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *recordHttpHandler) Search(c *gin.Context) {
	handlerOpts := &HandlerOpts{
		Name:   c.Request.URL.Path,
		Method: c.Request.Method,
		Params: c.Request.URL.Query(),
	}

	// retrieve query
	query := c.Query("query")
	if query == "" {
		r.errorResponse(c, handlerOpts, http.StatusBadRequest,
			messages.BAD_REQUEST, messages.QUERY_PARAMETER_EMPTY,
		)
		return
	}

	// retrieve amount
	sAmount := c.Query("amount")
	amount := 50 // default to 50 results
	if sAmount != "" {
		var err error
		amount, err = strconv.Atoi(sAmount)
		if err != nil {
			r.errorResponse(c, handlerOpts, http.StatusBadRequest,
				messages.BAD_REQUEST, messages.AMOUNT_INSUFFICENT,
			)
			return
		}
	}

	// retreive search type
	searchType := c.Query("searchType")
	if searchType == "" {
		searchType = constant.SEARCH_BY_DEFAULT
	}

	// search for records
	records, err := r.recordUsecase.Search("record", query, searchType, amount)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			r.errorResponse(c, handlerOpts, er.StatusCode, er.Message, er.Error())
			return
		} else {
			r.errorResponse(c, handlerOpts, http.StatusInternalServerError,
				messages.INTERNAL_SERVER_ERROR, err.Error(),
			)
			return
		}
	}

	resp := ResponseOptions{
		Response: records,
		OptionalResponse: &SearchRecordLogResponse{
			Length: len(records.Results),
		},
	}

	r.successResponse(c, handlerOpts, http.StatusOK, resp)
}

func (r *recordHttpHandler) SearchByRecordIndex(c *gin.Context) {
	handlerOpts := &HandlerOpts{
		Name:   c.Request.URL.Path,
		Method: c.Request.Method,
	}
	// retrieve query
	recordIndex := c.Param("recordIndex")
	if recordIndex == "" {
		r.errorResponse(c, handlerOpts, http.StatusBadRequest,
			messages.BAD_REQUEST, messages.QUERY_PARAMETER_EMPTY,
		)
		return
	}

	record, err := r.recordUsecase.SearchByRecordIndex("record", recordIndex)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			r.errorResponse(c, handlerOpts, er.StatusCode, er.Message, er.Error())
			return
		} else {
			r.errorResponse(c, handlerOpts, http.StatusInternalServerError,
				messages.INTERNAL_SERVER_ERROR, err.Error(),
			)
			return
		}
	}
	if record == nil {
		r.errorResponse(c, handlerOpts, http.StatusNotFound, messages.NOT_FOUND, messages.RECORD_INDEX_NOT_FOUND)
		return
	}

	res := ResponseOptions{
		Response: record,
		OptionalResponse: &RecordIndexLogResponse{
			Index: recordIndex,
		},
	}
	r.successResponse(c, handlerOpts, http.StatusOK, res)
}

// GetAllRecords retrieves all records from the elastic database
// and sends a response back to the client.
//
// Response:
// - 200 & A list of all records retrieved from the database.
// - 500: An internal server error occurred.
func (r *recordHttpHandler) GetAllRecords(c *gin.Context) {
	handlerOpts := &HandlerOpts{
		Name:   c.Request.URL.Path,
		Method: c.Request.Method,
		Params: "",
	}

	records, err := r.recordUsecase.GetAllRecords("record")
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			r.errorResponse(c, handlerOpts, er.StatusCode, er.Message, er.Error())
			return
		} else {
			r.errorResponse(c, handlerOpts, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR, err.Error())
			return
		}
	}

	res := ResponseOptions{
		Response: records,
		OptionalResponse: &SearchRecordLogResponse{
			Length: len(records),
		},
	}
	r.successResponse(c, handlerOpts, http.StatusOK, res)
}
