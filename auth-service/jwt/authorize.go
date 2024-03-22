package jwt

import (
	"auth-service/config"

	"github.com/labstack/echo/v4"
)

// HasMatchRole checks if the given role matches the role in the claims extracted from the context.
//
// Parameters:
// - requesterRole: the role of the person who requset
// - requiredRole: the role required to access the resource
//
// Returns:
// - bool: true if the given role matches the role in the claims, false otherwise.
func HasAuthorizeRole(requesterRole string, requiredRole string, allowEqualRole bool) bool {
	cfg := config.GetConfig()
	if allowEqualRole {
		if cfg.App.RolesMap[requiredRole] > cfg.App.RolesMap[requesterRole] {
			// Require role > requester -> false
			return false
		}
	} else {
		if cfg.App.RolesMap[requiredRole] >= cfg.App.RolesMap[requesterRole] {
			// Require role >= requester -> true
			return false
		}
	}

	return true
}

func GetRole(c echo.Context) (string, error) {
	claims, err := ValidateAndExtractClaims(c)
	if err != nil {
		return "", err
	}
	return claims.Role, nil
}
