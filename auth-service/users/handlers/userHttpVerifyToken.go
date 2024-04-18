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

	handlerOpts := NewHandlerOpts(c)
	handlerOpts.Params = map[string]string{"token": ""}

	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return h.errorResponse(c, handlerOpts, http.StatusBadRequest, messages.INVALID_TOKEN)
	}

	handlerOpts.Params = map[string]string{"token": token}

	result, err := h.usersUsecase.VerifyToken(c)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return h.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
		} else {
			return h.errorResponse(c, handlerOpts, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}
	}

	res := ResponseOptions{
		Response: &verifyStruct{
			Message: messages.SUCCESSFUL_VERIFIY_AUTH_TOKEN,
			Result:  result,
		},
	}

	return h.successResponse(c, handlerOpts, http.StatusOK, res)
}

// Verify Reset Token to verify the time valid of token (15 minute)
// Route Params - `token`
//
// Response
// - 200 OK & result (true/false)
// - 404 Not found ; token == "" or not attach token
// - 500 internal server error
func (h *usersHttpHandler) VerifyResetToken(c echo.Context) error {
	handlerOpts := NewHandlerOpts(c)
	handlerOpts.Params = map[string]string{"token": ""}

	token := c.Param("token")
	if token == "" {
		return h.errorResponse(c, handlerOpts, http.StatusNotFound, messages.INVALID_TOKEN)
	}

	handlerOpts.Params = map[string]string{"token": token}

	ch, err := h.usersUsecase.VerifyResetToken(token)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return h.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
		} else {
			return h.errorResponse(c, handlerOpts, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}
	}

	res := ResponseOptions{
		Response: &verifyStruct{
			Message: messages.SUCCESSFUL_VERIFY_RESET_TOKEN,
			Result:  ch,
		},
	}

	return h.successResponse(c, handlerOpts, http.StatusOK, res)
}
