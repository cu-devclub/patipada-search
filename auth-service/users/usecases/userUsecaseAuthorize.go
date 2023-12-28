package usecases

import (
	"auth-service/config"
	"auth-service/errors"
	"auth-service/jwt"
	"fmt"

	"github.com/labstack/echo/v4"
)

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
func (u *UsersUsecaseImpl) Authorize(c echo.Context, requireRole string) (bool, *errors.RequestError) {
	cfg := config.GetConfig()
	role, err := jwt.GetRole(c)
	if err != nil {
		return false, errors.CreateError(err.StatusCode, err.Error())
	}
	// check if requireRole one of admin, super-admin, user
	fmt.Println("require role", requireRole)
	fmt.Println(requireRole != cfg.User.Admins.Role && requireRole != cfg.User.Users.Role && requireRole != cfg.User.SuperAdmin.Role )
	if requireRole != cfg.User.Admins.Role && requireRole != cfg.User.Users.Role && requireRole != cfg.User.SuperAdmin.Role {
		return false, errors.CreateError(400, "Invalid requireRole")
	}

	ch := jwt.HasAuthorizeRole(role, requireRole, true)
	if err != nil {
		return false, errors.CreateError(err.StatusCode, err.Error())
	}

	return ch, nil
}
