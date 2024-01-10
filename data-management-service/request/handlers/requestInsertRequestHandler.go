package handlers

import (
	"data-management/messages"
	"data-management/request/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InsertRequest is a HTTP handler function that inserts a new Request into the database.
// It takes a *gin.Context as an argument, which contains the HTTP request context.
//
// The function first tries to bind the JSON body of the request to a models.Request struct.
// The models.Request struct has the following fields:
//
//	Index:      The index of the request. It is a string and is required.
//	YoutubeURL: The URL of the YouTube video for the request. It is a string and is required.
//	Question:   The question of the request. It is a string and is required.
//	Answer:     The answer of the request. It is a string and is required.
//	StartTime:  The start time of the request in the YouTube video. It is a string and is required.
//	EndTime:    The end time of the request in the YouTube video. It is a string and is required.
//	CreatedAt:  The creation time of the request. It is a time.Time and is optional.
//	UpdatedAt:  The update time of the request. It is a time.Time and is optional.
//	By: 	   The user who created the request. It is a string.
//
// If the binding fails, it responds with a 400 Bad Request status and returns.
//
// Then, it calls the InsertRequest method of the requestUsecase to insert the request into the database.
// If the insertion fails, it responds with the status code and error message from the error returned by InsertRecord, and returns.
//
// Error Status Codes:
//
//	400: ERR_REQUEST_INDEX_NOT_EXISTS
//	500: INTERNAL_SERVER_ERROR
//
// If the insertion is successful,
// Response
//
//	201: Created status and the inserted request.
//
// Usage:
//
//	router.POST("/requests", requestHandler.InsertRequest)
//
// This will create a new route that accepts POST requests at /requests.
// The body of the request should be a JSON object that matches the structure of the models.Request struct.
func (r *requestHandler) InsertRequest(c *gin.Context) {
	var request models.Request

	if err := c.ShouldBindJSON(&request); err != nil {
		responseJSON(c, http.StatusBadRequest, messages.BAD_REQUEST, nil)
		return
	}
	er := r.requestUsecase.InsertRequest(&request)
	if er != nil {
		responseJSON(c, er.StatusCode, er.Error(), nil)
		return
	}

	responseJSON(c, http.StatusCreated, messages.SUCCESS_INSERT_REQUEST, request)
}
