package handlers

import (
	"auth-service/errors"
	"auth-service/messages"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Verify Reset Token to verify the time valid of token (15 minute)
// Route Params - `token`
//
// Response
// - 200 OK & result (true/false)
// - 404 Not found ; token == "" or not attach token
// - 500 internal server error
func (h *usersHttpHandler) VerifyResetToken(c echo.Context) error {
	token := c.Param("token")
	ch, err := h.usersUsecase.VerifyResetToken(token)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return verifyTokenResponse(c, er.StatusCode, er.Error(), false)
		} else {
			return verifyTokenResponse(c, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR, false)
		}
	}
	return verifyTokenResponse(c, http.StatusOK, messages.SUCCESSFUL_VERIFY_RESET_TOKEN, ch)
}
