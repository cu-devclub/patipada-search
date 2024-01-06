package handlers

import (
	"auth-service/messages"

	"github.com/labstack/echo/v4"
)

// Verify Token to verify the time valid of auth token
// Header - Authorization : <token>
//
// Response
// - 200 OK & result (true/false)
// - 400 Bad request ; missing token
// - 401 Unauthorize ; invalid token
// - 500 internal server error
func (h *usersHttpHandler) VerifyToken(c echo.Context) error {
	result, err := h.usersUsecase.VerifyToken(c)
	if err != nil {
		return verifyTokenResponse(c, err.StatusCode, err.Error(), result)
	}
	return verifyTokenResponse(c, 200, messages.SUCCESSFUL_VERIFIY_AUTH_TOKEN, result)
}
