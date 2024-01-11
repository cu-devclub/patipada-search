package server

import (
	"context"
	"data-management/auth_proto"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// AuthMiddleware is a Gin middleware for checking authorization
func (g *ginServer) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		log.Println("Connecting to auth service via gRPC")
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", g.cfg.App.AuthService, g.cfg.App.GRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		log.Println("Connection success...")
		client := auth_proto.NewAuthServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		log.Println("Checking authorization ...")
		result, err := client.Authorization(ctx, &auth_proto.AuthorizationRequest{Token: token, RequiredRole: "admin"})
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if result.IsAuthorized != true {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
