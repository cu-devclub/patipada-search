package handlers

import (
	"auth-service/errors"
	"auth-service/messages"
	"net/http"

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
		if er, ok := err.(*errors.RequestError); ok {
			return verifyTokenResponse(c, er.StatusCode, er.Error(), false)
		} else {
			return verifyTokenResponse(c, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR, false)
		}
	}
	return verifyTokenResponse(c, 200, messages.SUCCESSFUL_VERIFIY_AUTH_TOKEN, result)
}
