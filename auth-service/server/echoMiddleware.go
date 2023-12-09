package server

import (
	"auth-service/config"
	"auth-service/jwt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RoleBasedMiddleware is a middleware function that checks if the user has the required role.
//
// It takes a string parameter 'requiredRole' which represents the role that the user must have.
// The middleware function returns a 'echo.MiddlewareFunc' which is a function that takes a 'echo.HandlerFunc'
// and returns a 'echo.HandlerFunc'. The inner function takes a 'echo.Context' and returns an error.
// If the user's role is lower than the required role, it returns a JSON response with status code 403 (Forbidden)
// and an error message. Otherwise, it calls the next handler in the chain.
func RoleBasedMiddleware(requiredRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, err := jwt.ClaimToken(c)
			if err != nil {
				return err
			}
			cfg := config.GetConfig()

			if cfg.App.RolesMap[claims.Role] < cfg.App.RolesMap[requiredRole] {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error": "Unauthorized access",
				})
			}

			return next(c)
		}
	}
}
