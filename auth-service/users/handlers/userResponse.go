package handlers

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

func (handler *usersHttpHandler) successResponse(c echo.Context, handlerOpts *HandlerOpts, responseCode int, resp ResponseOptions) error {
	var body interface{}
	if resp.LogResponseOptional == nil {
		body = resp.Response
	} else {
		body = resp.LogResponseOptional
	}

	logRes := &Response{
		Code: responseCode,
		Body: body,
	}
	slog.Info(
		"Success Request",
		slog.Any("Handler", handlerOpts),
		slog.Any("Response", logRes),
	)
	return c.JSON(responseCode, resp.Response)
}

func (handler *usersHttpHandler) errorResponse(c echo.Context, handlerOpts *HandlerOpts, responseCode int, errMessage string) error {
	res := &Response{
		Code: responseCode,
		Body: errMessage,
	}
	slog.Error(
		"Error Request",
		slog.Any("Handler", handlerOpts),
		slog.Any("Response", res),
	)

	return c.JSON(responseCode, &errorResponseStruct{
		ErrMessage: errMessage,
	})
}