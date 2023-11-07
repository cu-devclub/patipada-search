package record

import (
	"net/http"
	"search-esdb-service/internal/es"
	"github.com/gin-gonic/gin"
)

func DisplayAllRecords(c *gin.Context) {
	documents,err := es.GetAllDocumentsFromIndex("record")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"Length":len(documents),"Documents":documents})
}

func Search(c *gin.Context) {
	query := c.Query("query")

	matchedDocuments, err := es.SearchInIndex(query,"record")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"results": matchedDocuments})
}
