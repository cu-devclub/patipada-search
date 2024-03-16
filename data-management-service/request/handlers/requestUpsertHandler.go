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
	handlerOpts := &HandlerOpts{
		Name:   c.Request.URL.Path,
		Method: c.Request.Method,
		Params: request,
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		r.errorResponse(c, handlerOpts, http.StatusBadRequest, messages.MISSING_REQUEST_INDEX)
		return
	}

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
		Response: request.ToString(),
	}

	r.successResponse(c, *handlerOpts, http.StatusCreated, resp)
}

func (r *requestHandler) UpdateRequest(c *gin.Context) {
	var request models.Request
	handlerOpts := &HandlerOpts{
		Name:   c.Request.URL.Path,
		Method: c.Request.Method,
		Params: request,
	}
	if err := c.ShouldBind(&request); err != nil {
		r.errorResponse(c, handlerOpts, http.StatusBadRequest, messages.MISSING_REQUEST_INDEX)
		return
	}

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
		Response: request.ToString(),
	}
	r.successResponse(c, *handlerOpts, http.StatusOK, resp)
}
