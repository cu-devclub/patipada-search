package handlers

import "github.com/labstack/echo/v4"

type UsersHandler interface {
	// InsertUsers handles the HTTP request to insert users.
	//
	// It takes in a `c` parameter of type `echo.Context` which represents the current
	// HTTP context. It returns an error.
	// If the request body is not valid, it returns an error of 409 conflict .
	// If there is an error while inserting users, it returns an error of 500 .
	// Otherwise, it returns a 201 Created status code.
	InsertUsers(c echo.Context) error

	// Login handles the login request.
	//
	// It takes a `c` parameter of type `echo.Context` and returns an `error`.
	// If the request body is not valid, it returns an error of 400 bad request.
	// If username or password is incorrect, it returns an error of 401.
	// Otherwise, it returns a 200 OK status code and the token.
	Login(c echo.Context) error

}
