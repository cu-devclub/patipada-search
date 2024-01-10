package handlers

import (
	"data-management/messages"

	"github.com/gin-gonic/gin"
)

// GetAllRequests handles the HTTP request for retrieving all requests.
// It uses the requestUsecase to get all requests and sends the response in JSON format.
//
// Response :
// - 200 OK if success and return all requests in JSON format
// - 500 Internal Server Error if internal server error
//
// Usage :
//
//	router.GET("/requests", requestHandler.GetAllRequests)
func (r *requestHandler) GetAllRequests(c *gin.Context) {
	modelsRequests, err := r.requestUsecase.GetAllRequests()
	if err != nil {
		responseJSON(c, err.StatusCode, err.Error(), nil)
		return
	}

	responseJSON(c, 200, messages.SUCCESS_GET_REQUEST, modelsRequests)
}

// GetRequestByRequestID handles the HTTP request for retrieving a request by its RequestID.
// It uses the requestUsecase to get the request and sends the response in JSON format.
//
// Response :
// - 200 OK if success and return the request in JSON format
// - 404 Not Found if no request found
// - 500 Internal Server Error if internal server error
//
// Usage :
//
//	router.GET("/requests/:requestID", requestHandler.GetRequestByRequestID)
func (r *requestHandler) GetRequestByRequestID(c *gin.Context) {
	requestID := c.Param("requestID")
	modelsRequest, err := r.requestUsecase.GetRequestByRequestID(requestID)
	if err != nil {
		responseJSON(c, err.StatusCode, err.Error(), nil)
		return
	}

	responseJSON(c, 200, "OK", modelsRequest)
}


func (r *requestHandler) GetRequestByRecordIndex(c *gin.Context) {
	index := c.Param("index")
	modelsRequest, err := r.requestUsecase.GetRequestByRecordIndex(index)
	if err != nil {
		responseJSON(c, err.StatusCode, err.Error(), nil)
		return
	}

	responseJSON(c, 200, "OK", modelsRequest)
}