package server

import (
	"auth-service/config"
	"auth-service/errors"
	"auth-service/jwt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

var unAuthroizedMessage = map[string]string{"message": "Unauthorized"}

// AuthMiddleware is an Echo middleware for checking authorization
func (s *echoServer) AuthMiddleware(requiredRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			token := c.Request().Header.Get("Authorization")
			if token == "" {
				slog.Error("AuthMiddleware", slog.Any("Error", "No token"))
				return c.JSON(http.StatusUnauthorized, unAuthroizedMessage)
			}

			result, err := authorize(c, requiredRole)
			if err != nil || !result {
				slog.Error("AuthMiddleware", slog.Any("AuthorizeFailed", err))
				return c.JSON(http.StatusUnauthorized, unAuthroizedMessage)
			}

			return next(c)
		}
	}
}

func authorize(c echo.Context, requireRole string) (bool, error) {
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
