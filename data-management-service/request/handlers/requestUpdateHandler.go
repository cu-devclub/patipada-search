package handlers

import (
	"data-management/messages"
	"data-management/request/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateRequest updates a request using the gin context.
// The function first tries to bind the JSON body of the request to a models.Request struct.
// The models.Request struct has the following fields:
//
//  ID : The ID of the request. It is a string and is required.
//  RequestID:  The ID of the request. It is a string and is required.
//	Index:      The index of the request. It is a string and is required.
//	YoutubeURL: The URL of the YouTube video for the request. It is a string and is required.
//	Question:   The question of the request. It is a string and is required.
//	Answer:     The answer of the request. It is a string and is required.
//	StartTime:  The start time of the request in the YouTube video. It is a string and is required.
//	EndTime:    The end time of the request in the YouTube video. It is a string and is required.
//	CreatedAt:  The creation time of the request. It is a time.Time and is optional.
//	UpdatedAt:  The update time of the request. It is a time.Time and is optional.
//	By: 	   The user who created the request. It is a string.
//  ApprovedBy: The user who approved the request. It is a string.
//  Status:     The status of the request. It is a string.
//
// If the binding fails, it responds with a 400 Bad Request status and returns.

// Success
// - 200 OK if success and return the updated request in JSON format
//
// Error
// - 400 Bad Request if the binding fails or some fields are error
// - 500 Internal Server Error if internal server error
func (r *requestHandler) UpdateRequest(c *gin.Context) {
	var request models.Request
	if err := c.ShouldBind(&request); err != nil {
		log.Println("Error binding JSON Handler; Error: ", err)
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
