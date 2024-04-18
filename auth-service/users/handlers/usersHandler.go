package handlers

import (
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
)

type HandlerOpts struct {
	Name      string `json:"name"`
	Method    string `json:"method"`
	Params    any    `json:"params"`
	Time      string `json:"time"`
	RemoteIP  string `json:"remote_ip"`
	Host      string `json:"host"`
	UserAgent string `json:"user_agent"`
}

func NewHandlerOpts(c echo.Context) *HandlerOpts {
	start := time.Now()
	return &HandlerOpts{
		Name:      c.Request().URL.Path,
		Method:    c.Request().Method,
		Params:    nil,
		Time:      start.Format(time.RFC3339Nano),
		RemoteIP:  c.Request().RemoteAddr,
		Host:      c.Request().Host,
		UserAgent: c.Request().UserAgent(),
	}
}

func (h HandlerOpts) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("name", h.Name),
		slog.String("method", h.Method),
		slog.Any("params", h.Params),
		slog.String("time", h.Time),
		slog.String("remote_ip", h.RemoteIP),
		slog.String("host", h.Host),
		slog.String("user_agent", h.UserAgent),
	)
}

type UsersHandler interface {
	// RegisterUser handles the HTTP request to register users.
	// If new user role is "admin" or "super-admin"
	// then requester role must be "admin" or "super-admin"
	//
	// It takes in a `c` parameter of type `echo.Context`
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
	RegisterUser(c echo.Context) error

	// Login handles the login request.
	//
	// It takes in a `c` parameter of type `echo.Context`
	// Parameters (JSON) :
	// - username : string ; 3 <= length <= 50
	// - password : string ; 8 <= length <= 50
	//
	// Response
	// - 200 , role and token
	// - 400 bad request ; some field missing or input invalid
	// - 401 unauthorized ;  username or password incorrect
	// - 500 internal server error
	Login(c echo.Context) error

	// Request the link to reset password
	// Link when sent to input email if valid
	// Route Parameter
	// - email (string,email)

	// Response
	// - 200 OK & reset password token (also send to email)
	// - 400 bad request (invalid email)
	// - 404 User not found (email not exists)
	// - 500 internal server error
	ForgetPassword(c echo.Context) error

	// Reset Password
	// Parameters(JSON)
	// - token (string) ; reset password token
	// - password (string) ; new password ; 8 <= length <= 50
	//
	// Response
	// - 201 Created ; Update password success
	// - 400 bad request (invalid format password)
	// - 401 Unautorize ; invalid reset password
	// - 500 internal server error
	ResetPassword(c echo.Context) error

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
	RemoveUser(c echo.Context) error

	// Verify Reset Token to verify the time valid of token (15 minute)
	// Route Params - `token`
	//
	// Response
	// - 200 OK & result (true/false)
	// - 404 Not found ; token == "" or not attach token
	// - 500 internal server error
	VerifyResetToken(c echo.Context) error

	// Change Password
	// Parameter(JSON)
	// - oldPassword (string) ; old password ; 8 <= length <= 50
	// - newPassword (string) ; new password ; 8 <= length <= 50
	//
	// Response
	// - 200 OK ; Update password success
	// - 400 bad request (invalid format password)
	// - 401 Unautorize ; invalid old password
	// - 500 internal server error
	ChangePassword(c echo.Context) error

	// Verify Token to verify the time valid of auth token
	// Header - Authorization : <token>
	//
	// Response
	// - 200 OK & result (true/false)
	// - 400 Bad request ; missing token
	// - 401 Unauthorize ; invalid token
	// - 500 internal server error
	VerifyToken(c echo.Context) error

	// Authorize to verify the user authorization
	// Header - Authorization : <token>
	//
	// Query Params
	// - requiredRole (string) ; one of admin, super-admin, user
	//
	// Response
	// - 200 OK & result (true/false)
	// - 400 Bad request ; missing token
	// - 401 Unauthorize ; invalid token
	// - 500 internal server error
	Authorize(c echo.Context) error

	GetAllUsers(c echo.Context) error
}
