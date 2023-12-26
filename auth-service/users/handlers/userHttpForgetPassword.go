package handlers

import (
	"auth-service/errors"
	"auth-service/messages"
	"auth-service/users/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Request the link to reset password
// Link when sent to input email if valid
// Route Parameter
// - email (string,email)

// Response
// - 200 OK & reset password token (also send to email)
// - 400 bad request (invalid email)
// - 404 User not found (email not exists)
// - 500 internal server error
func (h *usersHttpHandler) ForgetPassword(c echo.Context) error {
	email := c.Param("email")
	if email == "" {
		return forgetPasswordResponse(c, http.StatusBadRequest, messages.BAD_REQUEST, "")
	}

	in := &models.ForgetPassword{
		Email: email,
	}

	token, err := h.usersUsecase.ForgetPassword(in)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return forgetPasswordResponse(c, er.StatusCode, er.Error(), "")
		} else {
			return forgetPasswordResponse(c, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR, "")
		}
	}
	return forgetPasswordResponse(c, http.StatusOK, messages.SUCCESSFUL_SEND_EMAIL_FORGET_PASSWORD, token)
}
