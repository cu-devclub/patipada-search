package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func responseJSON(c *gin.Context, status int, message string, data interface{}) {
	log.Println("response with status: ", status, " message: ", message, " data: ", data)
	c.JSON(status, gin.H{"message": message, "data": data})
}
