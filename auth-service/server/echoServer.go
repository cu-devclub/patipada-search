package server

import (
	"fmt"

	"auth-service/config"
	"auth-service/jwt"
	usersHandlers "auth-service/users/handlers"
	usersRepositories "auth-service/users/repositories"
	usersUsecases "auth-service/users/usecases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	app *echo.Echo
	db  *gorm.DB
	cfg *config.Config
}

func NewEchoServer(cfg *config.Config, db *gorm.DB) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *echoServer) Start() {
	s.initializeUsersHttpHandler()

	s.app.Use(middleware.Logger())

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

// initializeUsersHttpHandler initializes the users HTTP handler.
//
// No parameters.
// No return values.
func (s *echoServer) initializeUsersHttpHandler() {
	// Initialize all layers
	usersPostgresRepository := usersRepositories.NewUsersPostgresRepository(s.db)

	usersUsecase := usersUsecases.NewUsersUsecaseImpl(
		usersPostgresRepository,
	)

	usersHttpHandler := usersHandlers.NewUsersHttpHandler(usersUsecase)

	// Routers
	// JSON Params - username (string) and password (string)
	s.app.POST("/login", usersHttpHandler.Login)

	// JSON Params - token to verify inside "Authorization" header 
	// TODO : Improve in future => list of action with roles 
	s.app.POST("/authorize", jwt.Authorize)

	usersRouters := s.app.Group("users")

	//* validate token and role; need to be admin to configure users
	usersRouters.Use(jwt.ValidateToken)
	usersRouters.Use(RoleBasedMiddleware("admin"))

	// JSON Params - username (string) password (string) email(string,email) role(string,oneof=super-admin admin user)
	usersRouters.POST("", usersHttpHandler.InsertUsers)
}
