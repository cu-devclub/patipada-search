package handlers

import "github.com/gin-gonic/gin"

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func responseJSON(c *gin.Context, status int,message string, data any) {
	c.JSON(status, gin.H{"message": message, "data": data})
}
