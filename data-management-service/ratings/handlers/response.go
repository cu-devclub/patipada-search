package handlers

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type ResponseOptions struct {
	Response         any
	OptionalResponse any
}

type Response struct {
	Code int
	Body any
}

type RatingsLog struct {
	Amount int
}

func (r Response) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("Code", r.Code),
		slog.Any("Body", r.Body),
	)
}


func (r *ratingHandler) successResponse(c *gin.Context, handlerOpts HandlerOpts, responseCode int, resp ResponseOptions) {
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
		"Success Request",
		slog.Any("Handler", handlerOpts),
		slog.Any("Response", res),
	)

	c.JSON(responseCode, resp.Response)
}

func (r *ratingHandler) errorResponse(c *gin.Context, handlerOpts *HandlerOpts, responseCode int, response any) {
	res := &Response{
		Code: responseCode,
		Body: response,
	}

	slog.Error(
		"Error Request",
		slog.Any("Handler", handlerOpts),
		slog.Any("Response", res),
	)

	c.JSON(responseCode, gin.H{"error": response})
}
