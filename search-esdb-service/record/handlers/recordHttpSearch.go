package handlers

import (
	"net/http"
	"search-esdb-service/config"
	"search-esdb-service/constant"
	"search-esdb-service/errors"
	"search-esdb-service/logging"
	"search-esdb-service/messages"
	"search-esdb-service/monitoring"
	"search-esdb-service/record/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *recordHttpHandler) Search(c *gin.Context) {
	handlerOpts := NewHandlerOpts(c)
	handlerOpts.Params = c.Request.URL.Query()

	// retrieve query
	query := c.Query("query")
	if query == "" {
		r.errorResponse(c, handlerOpts, http.StatusBadRequest,
			messages.BAD_REQUEST, messages.QUERY_PARAMETER_EMPTY,
		)
		return
	}

	// retrieve offset
	sOffSet := c.Query("offset")
	offset := 0 // default to 0
	if sOffSet != "" {
		var err error
		offset, err = strconv.Atoi(sOffSet)
		if err != nil {
			r.errorResponse(c, handlerOpts, http.StatusBadRequest,
				messages.BAD_REQUEST, messages.OFFSET_INSUFFICENT,
			)
			return
		}
	}

	// retrieve amount
	sAmount := c.Query("amount")
	amount := 8 // default to 8 results (per page)
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

	// retreive search status
	searchStatus := c.Query("searchStatus")
	if searchStatus == "" {
		searchStatus = constant.SEARCH_STATUS_DRAFTING
	}

	countNeeded := c.Query("countNeeded") == "true"

	cfg := config.GetConfig()

	if searchStatus == constant.SEARCH_STATUS_CONFIRM {
		searchLogsPath := cfg.Static.SearchLogsConfirmPath
		logging.WriteLogsToFile(cfg.Static.LogsPath, searchLogsPath, helper.SearchLogsMessage(handlerOpts.Time, query, offset, amount))
	}

	// monitor search
	monitoring.MonitoringSearch(searchStatus)

	// search for records
	records, err := r.recordUsecase.Search("record", query, searchType, offset, amount, countNeeded)
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
			Amount: records.Amount,
			Status: searchStatus,
		},
	}

	r.successResponse(c, handlerOpts, http.StatusOK, resp)
}

func (r *recordHttpHandler) SearchByRecordIndex(c *gin.Context) {
	handlerOpts := NewHandlerOpts(c)
	handlerOpts.Params = map[string]string{"recordIndex": c.Param("recordIndex")}
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
	handlerOpts := NewHandlerOpts(c)

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
