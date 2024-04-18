package server

import (
	"fmt"
	"ml-gateway-service/config"
	"ml-gateway-service/gateway/handlers"
	"ml-gateway-service/gateway/repositories"
	"ml-gateway-service/gateway/usecases"

	"github.com/gin-gonic/gin"
)

type ginServer struct {
	app         *gin.Engine
	cfg         *config.Config
	gatewayArch *GatewayArch
}

type GatewayArch struct {
	Repo    repositories.Repository
	Usecase usecases.Usecase
	Handler handlers.Handler
}

func NewGinServer(cfg *config.Config) Server {
	serv := gin.New()

	g := &ginServer{
		app: serv,
		cfg: cfg,
	}

	return g
}

func (g *ginServer) GetGatewayArch() *GatewayArch {
	return g.gatewayArch
}

func (g *ginServer) Start() {
	g.initializeRoutes()
	g.app.Run(fmt.Sprintf(":%d", g.cfg.App.Port))
}

func (g *ginServer) initializeRoutes() {
	repo := repositories.NewGatewayRepository()

	usecase := usecases.NewUsecase(repo, &g.cfg.MlConfig)

	handler := handlers.NewGatewayHandler(&usecase)

	g.gatewayArch = &GatewayArch{
		Repo:    repo,
		Usecase: usecase,
		Handler: handler,
	}

	g.app.GET("/text2vec", handler.Text2Vec)
}
