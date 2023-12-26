package handlers

import (
	"auth-service/errors"
	"auth-service/jwt"
	"auth-service/messages"
	"auth-service/users/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Remove user by username & requestor role must be higher
// Header - Authorization : <token>
// Parameters (Route Param) :
// - username (string)
//
// Response
// - 200 OK
// - 400 bad request (invalid/missing username)
// - 401 Unauthorize ; missing token
// - 403 Forbidden ; no permission
// - 404 User not found (invalid username)
// - 500 internal server error
func (h *usersHttpHandler) RemoveUser(c echo.Context) error {
	username := c.Param("username")
	if username == "" {
		return baseResponse(c, http.StatusBadRequest, messages.BAD_REQUEST)
	}
	reqBody := &models.RemoveUserDto{
		Username: username,
	}

	requesterRole, err := jwt.GetRole(c)
	if err != nil {
		return baseResponse(c, http.StatusUnauthorized, messages.UNAUTHORIZED)
	}

	if err := h.usersUsecase.RemoveUser(requesterRole, reqBody); err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return baseResponse(c, er.StatusCode, er.Error())
		} else {
			return baseResponse(c, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}
	}
	return baseResponse(c, http.StatusOK, messages.SUCCESSFUL_REMOVE_USER)

}
