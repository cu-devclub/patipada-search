package handlers

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

func (handler *usersHttpHandler) successResponse(c echo.Context, handlerOpts *HandlerOpts, responseCode int, resp ResponseOptions) error {
	var body interface{}
	if resp.OptionalResponse == nil {
		body = resp.Response
	} else {
		body = resp.OptionalResponse
	}

	res := &Response{
		Code: responseCode,
		Body: body,
	}
	slog.Info(
		"Success Response",
		slog.Any("Handler", handlerOpts),
		slog.Any("Response", res),
	)
	return c.JSON(responseCode, resp.Response)
}

func (handler *usersHttpHandler) errorResponse(c echo.Context, handlerOpts *HandlerOpts, responseCode int, errMessage string) error {
	res := &Response{
		Code: responseCode,
		Body: errMessage,
	}
	slog.Error(
		"Error Response",
		slog.Any("Handler", handlerOpts),
		slog.Any("Response", res),
	)

	return c.JSON(responseCode, &errorResponseStruct{
		ErrMessage: errMessage,
	})
}