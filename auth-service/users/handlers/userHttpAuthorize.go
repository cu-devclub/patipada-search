package handlers

import (
	"auth-service/messages"

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
func (h *usersHttpHandler) Authorize(c echo.Context) error {
	requiredRole := c.QueryParam("requiredRole")
	if requiredRole == "" {
		return verifyTokenResponse(c,400, messages.MISSING_REQUIRED_ROLE, false)
	}

	result, err := h.usersUsecase.Authorize(c,requiredRole)
	if err != nil {
		return verifyTokenResponse(c, err.StatusCode, err.Error(), result)
	}
	return verifyTokenResponse(c, 200, messages.SUCCESSFUL_AUTHORIZE, result)
}