package usecases

import (
	"auth-service/users/models"

	"github.com/labstack/echo/v4"
)

type UsersUsecase interface {
	// RegisterUser
	// If new user role is "admin" or "super-admin"
	// then requester role must be "admin" or "super-admin"
	// Parameters (JSON) :
	// - requesterRole : string ; one of admin, super-admin, user
	// - models.RegisterDto
	//		- username : string ; 3 <= length <= 50, unique
	//      - password : string ; 8 <= length <= 50, unique
	// 		- email : string ; valid email, unique
	// 		- role : string ; one of admin, super-admin, user
	//
	// Response
	// - 201 and user id
	// - 400 bad request ; or input invalid
	//      - Email already exsits => message `Email already exists`
	//      - Username already exsits => message `Username already exists`
	// - 409 conflict ; no permission when requester is not super-admin/admin
	// - 500 internal server error

	RegisterUser(requesterRole string, in *models.RegisterDto) (string, error)

	// Authentication
	// Parameters (models.LoginDto) :
	// - username : string ; 3 <= length <= 50
	// - password : string ; 8 <= length <= 50
	//
	// Response
	// - 200 , role and token
	// - 400 bad request ; some field missing or input invalid
	// - 401 unauthorized ;  username or password incorrect
	// - 500 internal server error
	Authentication(in *models.LoginDto) (string, string, error)

	// Request the link to reset password
	// Link when sent to input email if valid
	// Parameter(models.ForgetPassword)
	// - email (string,email)

	// Response
	// - 200 OK & reset password token (also send to email)
	// - 400 bad request (invalid email)
	// - 404 User not found (email not exists)
	// - 500 internal server error
	ForgetPassword(in *models.ForgetPassword) (string, error)

	// Reset Password
	// Parameters(models.ResetPassword)
	// - token (string) ; reset password token
	// - password (string) ; new password ; 8 <= length <= 50
	//
	// Response
	// - 201 Created ; Update password success
	// - 400 bad request (invalid format password)
	// - 401 Unautorize ; invalid reset password
	// - 500 internal server error
	ResetPassword(in *models.ResetPassword) error

	// Remove user by id & requestor role must be higher
	// Parameters  :
	// - requester Role (string) ; one of admin, super-admin, user
	// - models.RemoveUserDto
	// 		- id (string)
	//
	// Response
	// - 200 OK
	// - 400 bad request (invalid/missing id)
	// - 401 Unauthorize ; missing token
	// - 403 Forbidden ; no permission
	// - 404 User not found (invalid id)
	// - 500 internal server error
	RemoveUser(requesterRole string, in *models.RemoveUserDto) error

	// Verify Reset Token to verify the time valid of token (15 minute)
	// Parameter
	// - token ; string
	//
	// Response
	// - 200 OK & result (true/false)
	// - 404 Not found ; token == "" or not attach token
	// - 500 internal server error
	VerifyResetToken(token string) (bool, error)

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
	ChangePassword(in *models.ChangePassword, username string) error

	// Verify Token to verify the time valid of auth token
	// Header - Authorization : <token>
	//
	// Response
	// - 200 OK & result (true/false)
	// - 400 Bad request ; missing token
	// - 401 Unauthorize ; invalid token
	// - 500 internal server error
	VerifyToken(c echo.Context) (bool, error)

	VerifyUsername(username string) (bool, error)

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
	Authorize(c echo.Context, requireRole string) (bool, error)

	GetAllUsersData() ([]*models.Users, error)
}
