package handlers

import (
	"data-management/errors"
	"data-management/messages"

	"github.com/gin-gonic/gin"
)

func (r *requestHandler) GetRequest(c *gin.Context) {
	handlerOpts := &HandlerOpts{
		Name:   c.Request.URL.Path,
		Method: c.Request.Method,
		Params: c.Request.URL.Query(),
	}

	status := c.Query("status")
	username := c.Query("username")
	requestID := c.Query("requestID")
	index := c.Query("index")
	approvedBy := c.Query("approvedBy")

	modelsRequest, err := r.requestUsecase.GetRequest(status, username, requestID, index, approvedBy)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			r.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
			return
		} else {
			r.errorResponse(c, handlerOpts, 500, messages.INTERNAL_SERVER_ERROR)
			return
		}
	}

	resp := ResponseOptions{
		Response: modelsRequest,
		OptionalResponse: &ArrayRequestsLog{
			Length: len(modelsRequest),
		},
	}

	r.successResponse(c, *handlerOpts, 200, resp)
}

func (r *requestHandler) GetLastestRequestOfRecord(c *gin.Context) {
	handlerOpts := &HandlerOpts{
		Name:   c.Request.URL.Path,
		Method: c.Request.Method,
		Params: c.Request.URL.Query(),
	}

	index := c.Query("index")
	if index == "" {
		r.errorResponse(c, handlerOpts, 400, messages.MISSING_REQUEST_INDEX)
		return
	}

	modelsRequest, err := r.requestUsecase.GetLastestRequestOfRecord(index)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			r.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
			return
		} else {
			r.errorResponse(c, handlerOpts, 500, messages.INTERNAL_SERVER_ERROR)
			return
		}
	}

	resp := ResponseOptions{
		Response: modelsRequest,
		OptionalResponse: &RequestLog{
			RequestID: modelsRequest.RequestID,
			Status:    modelsRequest.Status,
		},
	}

	r.successResponse(c, *handlerOpts, 200, resp)
}
