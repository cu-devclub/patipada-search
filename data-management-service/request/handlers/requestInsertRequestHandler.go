package handlers

import (
	"data-management/errors"
	"data-management/messages"
	"data-management/request/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *requestHandler) InsertRequest(c *gin.Context) {
	log.Println("InsertRequest handler : starting handler .....")
	var request models.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		responseJSON(c, http.StatusBadRequest, messages.BAD_REQUEST, nil)
		return
	}
	log.Println("With request: ", request.ToString(), " .....")

	err := r.requestUsecase.InsertRequest(&request)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			responseJSON(c, er.StatusCode, er.Error(), nil)
			return
		} else {
			responseJSON(c, 500, er.Error(), nil)
			return
		}
	}

	responseJSON(c, http.StatusCreated, messages.SUCCESS_INSERT_REQUEST, request)
}
