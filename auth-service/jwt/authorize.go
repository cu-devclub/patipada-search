package jwt

import (
	"auth-service/config"
	"auth-service/errors"
	"auth-service/messages"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Authorize authorizes the request by validating and extracting claims from the provided echo.Context.
//
// It returns an error if the validation fails, otherwise it returns the role of the claims in a JSON response.
func Authorize(c echo.Context) error {
	claims, err := ValidateAndExtractClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{
			"error": messages.UNAUTHORIZED,
		})
		return err
	}
	c.JSON(http.StatusOK, map[string]string{
		"role": claims.Role,
	})
	return nil
}

// HasMatchRole checks if the given role matches the role in the claims extracted from the context.
//
// Parameters:
// - c: the echo.Context object containing the request context.
// - role: the role to match against the role in the claims.
//
// Returns:
// - bool: true if the given role matches the role in the claims, false otherwise.
func HasAuthorizeRole(requesterRole string, requiredRole string, allowEqualRole bool) (bool, error) {
	cfg := config.GetConfig()

	if allowEqualRole {
		if cfg.App.RolesMap[requiredRole] > cfg.App.RolesMap[requesterRole] {
			// Require role > requester -> false
			return false, errors.CreateError(403, messages.NO_PERMISSION)
		}
	} else {
		if cfg.App.RolesMap[requiredRole] >= cfg.App.RolesMap[requesterRole] {
			// Require role >= requester -> true
			return false, errors.CreateError(403, messages.NO_PERMISSION)
		}
	}

	return true, nil
}

func GetRole(c echo.Context) (string, error) {
	claims, err := ValidateAndExtractClaims(c)
	if err != nil {
		return "", err
	}
	return claims.Role, nil
}
