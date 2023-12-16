package handlers

import "github.com/gin-gonic/gin"

type RecordHandler interface {
	GetAllRecords(c *gin.Context)
	Search(c *gin.Context)
}
