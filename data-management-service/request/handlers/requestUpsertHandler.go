package handlers

import (
	"data-management/errors"
	"data-management/messages"
	"data-management/request/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *requestHandler) InsertRequest(c *gin.Context) {
	var request models.Request

	handlerOpts := NewHandlerOpts(c)
	handlerOpts.Params = request

	if err := c.ShouldBindJSON(&request); err != nil {
		r.errorResponse(c, handlerOpts, http.StatusBadRequest, messages.MISSING_REQUEST_INDEX)
		return
	}

	handlerOpts.Params = request

	err := r.requestUsecase.InsertRequest(&request)
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
		Response: request,
	}

	r.successResponse(c, *handlerOpts, http.StatusCreated, resp)
}

func (r *requestHandler) UpdateRequest(c *gin.Context) {
	var request models.Request

	handlerOpts := NewHandlerOpts(c)
	handlerOpts.Params = request

	if err := c.ShouldBind(&request); err != nil {
		r.errorResponse(c, handlerOpts, http.StatusBadRequest, messages.MISSING_REQUEST_INDEX)
		return
	}

	handlerOpts.Params = request

	err := r.requestUsecase.UpdateRequest(&request)
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
		Response: request,
	}
	r.successResponse(c, *handlerOpts, http.StatusOK, resp)
}

func (r *requestHandler) SyncRequestRecord(c *gin.Context) {
	var request models.SyncRequestRecord

	handlerOpts := NewHandlerOpts(c)
	handlerOpts.Params = request

	if err := c.ShouldBind(&request); err != nil {
		r.errorResponse(c, handlerOpts, http.StatusBadRequest, messages.MISSING_REQUEST_INDEX)
		return
	}

	err := r.requestUsecase.SyncRequestRecord(&request)
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
		Response: request.RequestId,
	}
	r.successResponse(c, *handlerOpts, http.StatusOK, resp)
}

func (r *requestHandler) SyncAllRequestRecords(c *gin.Context) {
	handlerOpts := NewHandlerOpts(c)

	err := r.requestUsecase.SyncAllRequestRecords()
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
		Response: "All request records have been synced",
	}

	r.successResponse(c, *handlerOpts, http.StatusOK, resp)
}
