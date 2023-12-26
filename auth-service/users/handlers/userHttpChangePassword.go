package handlers

import (
	"auth-service/errors"
	"auth-service/jwt"
	"auth-service/messages"
	"auth-service/users/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Change Password
// Header Authorization - token
// Parameter(JSON)
// - oldPassword (string) ; old password ; 8 <= length <= 50
// - newPassword (string) ; new password ; 8 <= length <= 50
//
// Response
// - 200 OK ; Update password success
// - 400 bad request (invalid format password)
// - 401 Unautorize ; invalid old password
// - 422 ; New password == Old password
// - 500 internal server error
func (h *usersHttpHandler) ChangePassword(c echo.Context) error {
	reqBody := new(models.ChangePassword)
	if err := c.Bind(reqBody); err != nil {
		return baseResponse(c, http.StatusBadRequest, messages.BAD_REQUEST)
	}

	claims, err := jwt.ValidateAndExtractClaims(c)
	if err != nil {
		fmt.Println("INVALID TOKEN")
		return baseResponse(c, http.StatusUnauthorized, messages.UNAUTHORIZED)
	}
	username := claims.Username
	fmt.Println("VALID TOKEN, username", username)

	if err := h.usersUsecase.ChangePassword(reqBody, username); err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			fmt.Println("REQUEST ERROR", er.Error())
			return baseResponse(c, er.StatusCode, er.Error())
		} else {
			return baseResponse(c, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}
	}
	return baseResponse(c, http.StatusOK, messages.SUCCESSFUL_CHANGE_PASSWORD)
}
