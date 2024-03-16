package handlers

import (
	"auth-service/errors"
	"auth-service/jwt"
	"auth-service/messages"
	"auth-service/users/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RegisterUser handles the HTTP request to register users.
// If new user role is "admin" or "super-admin"
// then requester role must be "admin" or "super-admin"
//
// It takes in a `c` parameter of type `echo.Context`
// Header - Authorization : <token>
// Parameters (JSON) :
// - username : string ; 3 <= length <= 50, unique
// - password : string ; 8 <= length <= 50, unique
// - email : string ; valid email, unique
// - role : string ; one of admin, super-admin, user
//
// Response
// - 201 and user id
// - 400 bad request ; or input invalid
//   - Email already exsits => message `Email already exists`
//   - Username already exsits => message `Username already exists`
//
// - 409 conflict ; no permission when requester is not super-admin/admin
// - 500 internal server error
func (h *usersHttpHandler) RegisterUser(c echo.Context) error {

	handlerOpts := &HandlerOpts{
		Name:   c.Request().URL.Path,
		Method: c.Request().Method,
		Params: &models.RegisterLogDto{},
	}

	reqBody := new(models.RegisterDto)
	if err := c.Bind(reqBody); err != nil {
		return h.errorResponse(c, handlerOpts, http.StatusBadRequest, messages.BAD_REQUEST)
	}

	handlerOpts.Params = &models.RegisterLogDto{
		Username: reqBody.Username,
		Email:    reqBody.Email,
		Role:     reqBody.Role,
	}

	requesterRole := "user"
	var err error
	if reqBody.Role == "admin" || reqBody.Role == "super-admin" {
		requesterRole, err = jwt.GetRole(c)
		if err != nil {
			if er, ok := err.(*errors.RequestError); ok {
				return h.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
			} else {
				return h.errorResponse(c, handlerOpts, http.StatusUnauthorized, messages.UNAUTHORIZED)
			}
		}
	}

	userID, err := h.usersUsecase.RegisterUser(requesterRole, reqBody)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return h.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
		} else {
			return h.errorResponse(c, handlerOpts, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}
	}

	res := ResponseOptions{
		Response: &registerResponseStruct{
			Message: messages.SUCCESSFUL_REGISTER,
			UserID:  userID,
		},
	}

	return h.successResponse(c, handlerOpts, http.StatusCreated, res)
}

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
	reqBody := new(models.RemoveUserDto)
	handlerOpts := &HandlerOpts{
		Name:   c.Request().URL.Path,
		Method: c.Request().Method,
		Params: reqBody,
	}

	username := c.Param("username")
	if username == "" {
		return h.errorResponse(c, handlerOpts, http.StatusBadRequest, messages.BAD_REQUEST)
	}

	reqBody.Username = username

	requesterRole, err := jwt.GetRole(c)
	if err != nil {
		return h.errorResponse(c, handlerOpts, http.StatusUnauthorized, messages.UNAUTHORIZED)
	}

	if err := h.usersUsecase.RemoveUser(requesterRole, reqBody); err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return h.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
		} else {
			return h.errorResponse(c, handlerOpts, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}
	}

	res := ResponseOptions{
		Response: &baseResponseStruct{
			Message: messages.SUCCESSFUL_REMOVE_USER,
		},
	}

	return h.successResponse(c, handlerOpts, http.StatusOK, res)
}
