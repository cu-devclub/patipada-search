package handlers

import (
	"data-management/messages"
	"data-management/request/models"
	"net/http"

	"github.com/gin-gonic/gin"
)
// UpdateRequest updates a request using the gin context.
// It first binds the request body to a models.Request struct.
// Success
// - 200 OK if success and return the updated request in JSON format
// 
// Error
// - 400 Bad Request if the binding fails or some fields are error
// - 500 Internal Server Error if internal server error
func (r *requestHandler) UpdateRequest(c *gin.Context) {
	var request models.Request
	if err := c.ShouldBind(&request); err != nil {
		responseJSON(c, http.StatusBadRequest, messages.BAD_REQUEST, nil)
		return
	}

	err := r.requestUsecase.UpdateRequest(&request)
	if err != nil {
		responseJSON(c, err.StatusCode, err.Error(), nil)
		return
	}

	responseJSON(c, http.StatusOK, messages.SUCCESS_UPDATE_REQUEST, request)
}
