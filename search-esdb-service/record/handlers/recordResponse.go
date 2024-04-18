package handlers

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

// Implement as success response for future improvement
// if return has more condition, can implement successResponse further
func (r *recordHttpHandler) successResponse(c *gin.Context, handlerOpts *HandlerOpts, responseCode int, resp ResponseOptions) {
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

	c.JSON(responseCode, resp.Response)
}

func (r *recordHttpHandler) errorResponse(c *gin.Context, handlerOpts *HandlerOpts, responseCode int, response any, logMessage string) {
	res := &Response{
		Code: responseCode,
		Body: logMessage,
	}

	slog.Error(
		"Error Response",
		slog.Any("Handler", handlerOpts),
		slog.Any("Response", res),
	)

	c.JSON(responseCode, response)
}
