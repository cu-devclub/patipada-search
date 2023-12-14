package handlers

import "github.com/gin-gonic/gin"

// Implement as base response for future improvement
// if return has more condition, can implement baseResponse further 
func baseResponse(c *gin.Context, responseCode int, response interface{}) {
	c.JSON(responseCode, response)
}

