package handlers

import "github.com/labstack/echo/v4"

type baseResponse struct {
	Message string `json:"message"`
}

type loginResponseStruct struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

func insertResponse(c echo.Context, responseCode int, message string) error {
	return c.JSON(responseCode, &baseResponse{
		Message: message,
	})
}

func loginResponse(c echo.Context, responseCode int, message string, token string) error {
	return c.JSON(responseCode, &loginResponseStruct{
		Token:   token,
		Message: message,
	})
}
