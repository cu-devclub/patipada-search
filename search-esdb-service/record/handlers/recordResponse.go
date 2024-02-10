package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Implement as success response for future improvement
// if return has more condition, can implement successResponse further
func successResponse(c *gin.Context, responseCode int, response interface{}) {
	log.Println("Success with",responseCode)
	c.JSON(responseCode, response)
}

func errorResponse(c *gin.Context, responseCode int, response interface{}, logMessage string) {
	log.Println("Error with",responseCode, ":", logMessage)
	c.JSON(responseCode, response)
}
