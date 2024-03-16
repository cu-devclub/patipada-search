package usecases

import (
	"auth-service/config"
	"auth-service/errors"
	"auth-service/jwt"
	"auth-service/messages"
	"auth-service/users/helper"
	"auth-service/users/models"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

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
func (u *UsersUsecaseImpl) Authentication(in *models.LoginDto) (string, string, error) {
	// Validate data
	validator := validator.New()
	if err := validator.Struct(in); err != nil {
		return "", "", errors.CreateError(400, fmt.Sprintf("Error validating request body: %v", err))
	}

	user, err := u.usersRepository.GetUserByUsername(in.Username)
	if err != nil {
		return "", "", errors.CreateError(401, messages.WRONG_USERNAME_PASSWORD)
	}

	if err := helper.VerifyPassword(user.Password, in.Password+user.Salt); err != nil {
		return "", "", errors.CreateError(401, messages.WRONG_USERNAME_PASSWORD)
	}

	token, err := jwt.CreateToken(user.Username, user.Role)
	if err != nil {
		return "", "", errors.CreateError(500, messages.INTERNAL_SERVER_ERROR)
	}

	return token, user.Role, nil
}

//* Also called by GRPC
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
func (u *UsersUsecaseImpl) Authorize(c echo.Context, requireRole string) (bool, error) {
	cfg := config.GetConfig()
	role, err := jwt.GetRole(c)
	if err != nil {
		return false, err
	}
	// check if requireRole one of admin, super-admin, user
	if requireRole != cfg.User.Admins.Role && requireRole != cfg.User.Users.Role && requireRole != cfg.User.SuperAdmin.Role {
		return false, errors.CreateError(400, "Invalid requireRole")
	}

	ch := jwt.HasAuthorizeRole(role, requireRole, true)

	return ch, nil
}
