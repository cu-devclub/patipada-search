package handlers

import (
	"data-management/errors"
	"data-management/messages"
	"log"

	"github.com/gin-gonic/gin"
)

func (r *requestHandler) GetRequest(c *gin.Context) {
	status := c.Query("status")
	username := c.Query("username")
	requestID := c.Query("requestID")
	index := c.Query("index")
	approvedBy := c.Query("approvedBy")
	log.Println("GetRequest handler : starting handler ..... with status: ", status,
		" username: ", username,
		" requestID: ", requestID,
		" index: ", index,
		" approvedBy: ", approvedBy,
	)

	modelsRequest, err := r.requestUsecase.GetRequest(status, username, requestID, index, approvedBy)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			responseJSON(c, er.StatusCode, er.Error(), nil)
			return
		} else {
			responseJSON(c, 500, messages.INTERNAL_SERVER_ERROR, nil)
			return
		}
	}

	responseJSON(c, 200, messages.SUCCESS_GET_REQUEST, modelsRequest)
}

func (r *requestHandler) GetLastestRequestOfRecord(c *gin.Context) {
	index := c.Query("index")
	log.Println("GetLastestRequestOfRecord handler : starting handler with index ",index," .....")
	if index == "" {
		responseJSON(c, 400, messages.BAD_REQUEST, nil)
		return
	}

	modelsRequest, err := r.requestUsecase.GetLastestRequestOfRecord(index)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			responseJSON(c, er.StatusCode, er.Error(), nil)
			return
		} else {
			responseJSON(c, 500, messages.INTERNAL_SERVER_ERROR, nil)
			return
		}
	}

	responseJSON(c, 200, messages.SUCCESS_GET_REQUEST, modelsRequest)
}
