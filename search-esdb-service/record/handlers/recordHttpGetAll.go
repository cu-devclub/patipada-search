package handlers

import (
	"search-esdb-service/messages"

	"github.com/gin-gonic/gin"
)

// GetAllRecords retrieves all records from the elastic database
// and sends a response back to the client.
//
// Response:
// - 200 & A list of all records retrieved from the database.
// - 500: An internal server error occurred.
func (r *recordHttpHandler) GetAllRecords(c *gin.Context) {
	records, err := r.recordUsecase.GetAllRecords("record")
	if err != nil {
		baseResponse(c, 500, messages.INTERNAL_SERVER_ERROR)
		return
	}
	baseResponse(c, 200, records)
}
