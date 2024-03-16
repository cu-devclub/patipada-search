package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a Gin middleware for checking authorization
func (g *ginServer) AuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		result, err := g.comm.Authorization(token, requiredRole)
		if err != nil || !result {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
