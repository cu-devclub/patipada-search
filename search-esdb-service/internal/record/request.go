package record

import (
	"net/http"
	"search-esdb-service/internal/es"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	query := c.Query("query")

	matchedDocumentsQuestions, err := es.SearchInIndex(query,"record","question")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	matchedDocumentsAnswer, err := es.SearchInIndex(query,"record","answer")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mergeDocuments := append(matchedDocumentsQuestions, matchedDocumentsAnswer...)
	uniqueDocuments := GenerateUniqueDocuments(mergeDocuments)
	c.JSON(http.StatusOK, gin.H{"results": uniqueDocuments})
}
