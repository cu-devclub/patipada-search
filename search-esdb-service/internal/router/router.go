package router

import (
	"search-esdb-service/internal/record"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	// "github.com/spf13/viper"
)

func RouterEngine() *gin.Engine {
	r := gin.Default()
	r.Use(CORS())
	r.GET("/search", record.Search)
	return r
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		targets := viper.GetString("cors.target")
		// targets := "http://localhost:5173"
		c.Writer.Header().Set("Access-Control-Allow-Origin", targets)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
