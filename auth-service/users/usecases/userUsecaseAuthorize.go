package usecases

import (
	"auth-service/config"
	"auth-service/errors"
	"auth-service/jwt"

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
