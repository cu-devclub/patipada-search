package handlers

import (
	"data-management/errors"
	"data-management/messages"
	"data-management/request/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *requestHandler) UpdateRequest(c *gin.Context) {
	log.Println("Update request handler : starting handler .....")
	var request models.Request
	if err := c.ShouldBind(&request); err != nil {
		responseJSON(c, http.StatusBadRequest, messages.BAD_REQUEST, nil)
		return
	}

	log.Println("With Request: ", request.ToString())

	err := r.requestUsecase.UpdateRequest(&request)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			responseJSON(c, er.StatusCode, er.Error(), nil)
			return
		} else {
			responseJSON(c, 500, er.Error(), nil)
			return
		}
	}

	responseJSON(c, http.StatusOK, messages.SUCCESS_UPDATE_REQUEST, request)
}
