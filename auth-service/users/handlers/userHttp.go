package handlers

import (
	"net/http"

	"auth-service/errors"
	"auth-service/users/models"
	"auth-service/users/usecases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type usersHttpHandler struct {
	usersUsecase usecases.UsersUsecase
}

func NewUsersHttpHandler(usersUsecase usecases.UsersUsecase) UsersHandler {
	return &usersHttpHandler{
		usersUsecase: usersUsecase,
	}
}

// InsertUsers handles the HTTP request to insert users.
//
// It takes in a `c` parameter of type `echo.Context` which represents the current
// HTTP context. It returns an error.
// If the request body is not valid, it returns an error of 409 conflict .
// If there is an error while inserting users, it returns an error of 500 .
// Otherwise, it returns a 201 Created status code.
func (h *usersHttpHandler) InsertUsers(c echo.Context) error {
	reqBody := new(models.AddUsersData)

	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return insertResponse(c, http.StatusBadRequest, "Bad request")
	}

	if err := h.usersUsecase.UsersRegisterDataProcessing(reqBody); err != nil {
		if err, ok := err.(*errors.RequestError); ok {
			return insertResponse(c, err.StatusCode, err.Error())
		} else {
			return insertResponse(c, http.StatusInternalServerError, "Add user failed")
		}
	}

	return insertResponse(c, http.StatusCreated, "Success insert user")
}

// Login handles the login request.
//
// It takes a `c` parameter of type `echo.Context` and returns an `error`.
// If the request body is not valid, it returns an error of 400 bad request.
// If username or password is incorrect, it returns an error of 401.
// Otherwise, it returns a 200 OK status code and the token.
func (h *usersHttpHandler) Login(c echo.Context) error {
	reqBody := new(models.LoginDto)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return loginResponse(c, http.StatusBadRequest, "Bad request", "")
	}
	token, err := h.usersUsecase.Authentication(reqBody)
	if err != nil {
		if err, ok := err.(*errors.RequestError); ok {
			return loginResponse(c, err.StatusCode, err.Error(), "")
		} else {
			return loginResponse(c, http.StatusInternalServerError, "Authentication failed", "")
		}
	}
	return loginResponse(c, http.StatusOK, "Success login", token)

}
