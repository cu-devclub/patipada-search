package handlers

import (
	"auth-service/errors"
	"auth-service/messages"
	"auth-service/users/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Login handles the login request.
//
// It takes in a `c` parameter of type `echo.Context`
// Parameters (JSON) :
// - username : string ; 3 <= length <= 50
// - password : string ; 8 <= length <= 50
//
// Response
// - 200 , role and token
// - 400 bad request ; some field missing or input invalid
// - 401 unauthorized ;  username or password incorrect
// - 500 internal server error
func (h *usersHttpHandler) Login(c echo.Context) error {
	reqBody := new(models.LoginDto)
	handlerOpts := NewHandlerOpts(c)
	handlerOpts.Params = reqBody

	if err := c.Bind(reqBody); err != nil {
		return h.errorResponse(c, handlerOpts, http.StatusBadRequest, messages.BAD_REQUEST)
	}

	handlerOpts.Params = &models.LoginLogDto{
		Username: reqBody.Username,
	}

	token, role, err := h.usersUsecase.Authentication(reqBody)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return h.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
		} else {
			return h.errorResponse(c, handlerOpts, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}

	}

	res := ResponseOptions{
		Response: &loginResponseStruct{
			Role:    role,
			Token:   token,
			Message: messages.SUCCESSFUL_LOGIN,
		},
		LogResponseOptional: &loginResponseLogStruct{
			Username: reqBody.Username,
			Role:     role,
		},
	}

	return h.successResponse(c, handlerOpts, http.StatusOK, res)
}

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
	handlerOpts := NewHandlerOpts(c)
	handlerOpts.Params = map[string]string{"requiredRole": ""}

	requiredRole := c.QueryParam("requiredRole")
	if requiredRole == "" {
		return h.errorResponse(c, handlerOpts, http.StatusBadRequest, messages.MISSING_REQUIRED_ROLE)
	}

	handlerOpts.Params = map[string]string{"requiredRole": requiredRole}

	result, err := h.usersUsecase.Authorize(c, requiredRole)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return h.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
		} else {
			return h.errorResponse(c, handlerOpts, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}
	}

	resp := ResponseOptions{
		Response: &verifyStruct{
			Message: messages.SUCCESSFUL_AUTHORIZE,
			Result:  result,
		},
	}

	return h.successResponse(c, handlerOpts, http.StatusOK, resp)
}
