package server

import (
	"fmt"

	"auth-service/config"
	usersHandlers "auth-service/users/handlers"
	usersRepositories "auth-service/users/repositories"
	usersUsecases "auth-service/users/usecases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	App *echo.Echo
	db  *gorm.DB
	cfg *config.Config
}

func NewEchoServer(cfg *config.Config, db *gorm.DB) Server {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{cfg.App.FrontendURL, "http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	s := &echoServer{
		App: e,
		db:  db,
		cfg: cfg,
	}

	s.initializeUsersHttpHandler()

	return s
}

func (s *echoServer) Start() {
	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.App.Logger.Fatal(s.App.Start(serverUrl))
}

func (s *echoServer) GetHandler() *echo.Echo {
	return s.App
}

func (s *echoServer) GetDB() *gorm.DB {
	return s.db
}

// initializeUsersHttpHandler initializes the users HTTP handler.
//
// No parameters.
// No return values.
func (s *echoServer) initializeUsersHttpHandler() {
	// Initialize all layers
	usersPostgresRepository := usersRepositories.NewUsersPostgresRepository(s.db)
	userEmailRepository := usersRepositories.NewUserJordanWrightEmailing(s.cfg.Email.SenderName, s.cfg.Email.SenderEmail, s.cfg.Email.SenderPassword)
	usersUsecase := usersUsecases.NewUsersUsecaseImpl(
		usersPostgresRepository,
		userEmailRepository,
	)

	usersHttpHandler := usersHandlers.NewUsersHttpHandler(usersUsecase)

	// Initialize all routes
	s.initializedUnAuthRoutes(usersHttpHandler)
	s.initializedUserRoutes(usersHttpHandler)
	s.initializedAdminRoutes(usersHttpHandler)
}

func (s *echoServer) initializedUnAuthRoutes(handler usersHandlers.UsersHandler) {
	// Login Request
	// Parameters (JSON) :
	// - username : string ; 3 <= length <= 50
	// - password : string ; 8 <= length <= 50
	//
	// Response
	// - 200 , role and token
	// - 400 bad request ; some field missing or input invalid
	// - 401 unauthorized ;  username or password incorrect
	// - 500 internal server error
	s.App.POST("/login", handler.Login)

	// RegisterUser handles the HTTP request to register users.
	// If new user role is "admin" or "super-admin"
	// then requester role must be "admin" or "super-admin"
	// Header - Authorization : <token>
	// Parameters (JSON) :
	// - username : string ; 3 <= length <= 50, unique
	// - password : string ; 8 <= length <= 50, unique
	// - email : string ; valid email, unique
	// - role : string ; one of admin, super-admin, user
	//
	// Response
	// - 201 and user id
	// - 400 bad request ; or input invalid
	//      - Email already exsits => message `Email already exists`
	//      - Username already exsits => message `Username already exists`
	// - 409 conflict ; no permission when requester is not super-admin/admin
	// - 500 internal server error
	s.App.POST("/register", handler.RegisterUser)

	// Verify Reset Token to verify the time valid of token (15 minute)
	// Route Params - `token`
	//
	// Response
	// - 200 OK & result (true/false)
	// - 404 Not found ; token == "" or not attach token
	// - 500 internal server error
	s.App.GET("/verify-reset-token/:token", handler.VerifyResetToken)

	// Verify Token to verify the time valid of auth token
	// Header - Authorization : <token>
	//
	// Response
	// - 200 OK & result (true/false)
	// - 400 Bad request ; missing token
	// - 401 Unauthorize ; invalid token
	// - 500 internal server error
	s.App.GET("/verify-token", handler.VerifyToken)

	// Request the link to reset password
	// Link when sent to input email if valid
	// Route Parameter
	// - email (string,email)

	// Response
	// - 200 OK & reset password token (also send to email)
	// - 400 bad request (invalid email)
	// - 404 User not found (email not exists)
	// - 500 internal server error
	s.App.POST("forget-password/:email", handler.ForgetPassword)

	// Reset Password : change from reset password link
	// Parameters(JSON)
	// - token (string) ; reset password token
	// - password (string) ; new password ; 8 <= length <= 50
	//
	// Response
	// - 201 Created ; Update password success
	// - 400 bad request (invalid format password)
	// - 401 Unauthorize ; invalid reset password token
	// - 500 internal server error
	s.App.POST("reset-password", handler.ResetPassword)
}

func (s *echoServer) initializedUserRoutes(handler usersHandlers.UsersHandler) {
	// Routers
	userRoutes := s.App.Group("/")
	userRoutes.Use(s.AuthMiddleware("user"))

	// Authorize to verify the user authorization
	// Header - Authorization : <token>
	//
	// Query Params
	// - requiredRole (string) ; one of admin, super-admin, user
	//
	// Response
	// - 200 OK & result (true/false)
	// - 400 Bad request ; missing token or invalid requires role
	// - 401 Unauthorize ; invalid token
	// - 500 internal server error
	userRoutes.GET("authorize", handler.Authorize)

	// Change Password : manual change
	// Header Authorization - token
	// Parameter(JSON)
	// - oldPassword (string) ; old password ; 8 <= length <= 50
	// - newPassword (string) ; new password ; 8 <= length <= 50
	//
	// Response
	// - 200 OK ; Update password success
	// - 400 bad request (invalid format password)
	// - 401 Unauthorize ; invalid old password
	// - 500 internal server error
	userRoutes.POST("change-password", handler.ChangePassword)
}

func (s *echoServer) initializedAdminRoutes(handler usersHandlers.UsersHandler) {
	// Routers
	adminRoutes := s.App.Group("/")
	adminRoutes.Use(s.AuthMiddleware("admin"))

	// Remove user by id & requestor role must be higher
	// Header - Authorization : <token>
	// Parameters (Route Param) :
	// - id (string)
	//
	// Response
	// - 200 OK
	// - 400 bad request (invalid/missing id)
	// - 401 Unauthorize ; missing token
	// - 403 Forbidden ; no permission
	// - 404 User not found (invalid id)
	// - 500 internal server error
	adminRoutes.DELETE("user/:id", handler.RemoveUser)

	// Get all users
	// Header - Authorization : <token>
	//
	// Response
	// - 200 OK & users
	// - 401 Unauthorize ; missing token
	// - 403 Forbidden ; no permission
	// - 500 internal server error
	adminRoutes.GET("users", handler.GetAllUsers)
}
