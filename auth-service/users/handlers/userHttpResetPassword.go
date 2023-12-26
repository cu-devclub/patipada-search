package handlers

import (
	"auth-service/errors"
	"auth-service/messages"
	"auth-service/users/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Reset Password
// Parameters(JSON)
// - token (string) ; reset password token
// - password (string) ; new password ; 8 <= length <= 50
//
// Response
// - 201 Created ; Update password success
// - 400 bad request (invalid format password)
// - 401 Unautorize ; invalid reset password
// - 422 ; New password == Old password
// - 500 internal server error
func (h *usersHttpHandler) ResetPassword(c echo.Context) error {
	reqBody := new(models.ResetPassword)
	if err := c.Bind(reqBody); err != nil {
		return baseResponse(c, http.StatusBadRequest, messages.BAD_REQUEST)
	}

	if err := h.usersUsecase.ResetPassword(reqBody); err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return baseResponse(c, er.StatusCode, er.Error())
		} else {
			return baseResponse(c, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}
	}
	return baseResponse(c, http.StatusCreated, messages.SUCCESSFUL_RESET_PASSWORD)

}
