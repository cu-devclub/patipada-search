package handlers

import (
	"auth-service/errors"
	"auth-service/jwt"
	"auth-service/messages"
	"auth-service/users/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Change Password : Manual change password
// Header Authorization - token
// Parameter(JSON)
// - oldPassword (string) ; old password ; 8 <= length <= 50
// - newPassword (string) ; new password ; 8 <= length <= 50
//
// Response
// - 200 OK ; Update password success
// - 400 bad request (invalid format password)
// - 401 Unauthorize ; invalid old password
// - 500 internal server error
func (h *usersHttpHandler) ChangePassword(c echo.Context) error {
	reqBody := new(models.ChangePassword)
	handlerOpts := NewHandlerOpts(c)

	if err := c.Bind(reqBody); err != nil {
		return h.errorResponse(c, handlerOpts, http.StatusBadRequest, messages.BAD_REQUEST)

	}

	claims, err := jwt.ValidateAndExtractClaims(c)
	if err != nil {
		return h.errorResponse(c, handlerOpts, http.StatusUnauthorized, messages.UNAUTHORIZED)

	}
	username := claims.Username
	handlerOpts.Params = map[string]string{
		"username": username,
	}

	if err := h.usersUsecase.ChangePassword(reqBody, username); err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return h.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
		} else {
			return h.errorResponse(c, handlerOpts, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}

	}

	resp := ResponseOptions{
		Response: &baseResponseStruct{
			Message: messages.SUCCESSFUL_CHANGE_PASSWORD,
		},
	}

	return h.successResponse(c, handlerOpts, http.StatusOK, resp)
}

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
	in := new(models.ForgetPassword)
	handlerOpts := NewHandlerOpts(c)
	handlerOpts.Params = in

	email := c.Param("email")
	if email == "" {
		return h.errorResponse(c, handlerOpts, http.StatusBadRequest, messages.BAD_REQUEST)
	}

	in.Email = email

	token, err := h.usersUsecase.ForgetPassword(in)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return h.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
		} else {
			return h.errorResponse(c, handlerOpts, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}
	}

	resp := ResponseOptions{
		Response: &forgetPasswordStruct{
			Message: messages.SUCCESSFUL_SEND_EMAIL_FORGET_PASSWORD,
			Token:   token,
		},
		OptionalResponse: &forgetPasswordResLogStruct{
			Email:   email,
			Message: messages.SUCCESSFUL_SEND_EMAIL_FORGET_PASSWORD,
		},
	}

	return h.successResponse(c, handlerOpts, http.StatusOK, resp)
}

// Reset Password : change password when forget password from reset link
// Parameters(JSON)
// - token (string) ; reset password token
// - password (string) ; new password ; 8 <= length <= 50
//
// Response
// - 201 Created ; Update password success
// - 400 bad request (invalid format password)
// - 401 Unautorize ; invalid reset password
// - 500 internal server error
func (h *usersHttpHandler) ResetPassword(c echo.Context) error {
	handlerOpts := &HandlerOpts{
		Name:   c.Request().URL.Path,
		Method: c.Request().Method,
		Params: &models.ResetPasswordLog{},
	}

	reqBody := new(models.ResetPassword)
	if err := c.Bind(reqBody); err != nil {
		return h.errorResponse(c, handlerOpts, http.StatusBadRequest, messages.BAD_REQUEST)
	}

	handlerOpts.Params = &models.ResetPasswordLog{
		Token: reqBody.Token,
	}

	if err := h.usersUsecase.ResetPassword(reqBody); err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return h.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
		} else {
			return h.errorResponse(c, handlerOpts, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}
	}

	res := ResponseOptions{
		Response: &baseResponseStruct{
			Message: messages.SUCCESSFUL_RESET_PASSWORD,
		},
	}

	return h.successResponse(c, handlerOpts, http.StatusCreated, res)

}
