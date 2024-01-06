package handlers

import "github.com/labstack/echo/v4"

type baseResponseStruct struct {
	Message string `json:"message"`
}

type loginResponseStruct struct {
	Token   string `json:"token"`
	Role    string `json:"role"`
	Message string `json:"message"`
}
type registerResponseStruct struct {
	Message string `json:"message"`
	UserID  string `json:"user_id"`
}

type verifyTokenStruct struct {
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

type forgetPasswordStruct struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func baseResponse(c echo.Context, responseCode int, message string) error {
	return c.JSON(responseCode, &baseResponseStruct{
		Message: message,
	})
}

func loginResponse(c echo.Context, responseCode int, message string, token string, role string) error {
	return c.JSON(responseCode, &loginResponseStruct{
		Token:   token,
		Role:    role,
		Message: message,
	})
}

func registerResponse(c echo.Context, responseCode int, message string, userID string) error {
	return c.JSON(responseCode, &registerResponseStruct{
		Message: message,
		UserID:  userID,
	})
}

func verifyTokenResponse(c echo.Context, responseCode int, message string, valid bool) error {
	return c.JSON(responseCode, &verifyTokenStruct{
		Message: message,
		Result:  valid,
	})
}

func forgetPasswordResponse(c echo.Context, responseCode int, message string, token string) error {
	return c.JSON(responseCode, &forgetPasswordStruct{
		Message: message,
		Token:   token,
	})
}
