package handlers

import (
	"auth-service/errors"
	"auth-service/jwt"
	"auth-service/messages"
	"auth-service/users/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RegisterUser handles the HTTP request to register users.
// If new user role is "admin" or "super-admin"
// then requester role must be "admin" or "super-admin"
//
// It takes in a `c` parameter of type `echo.Context`
// Header - Authorization : <token>
// Parameters (JSON) :
// - username : string ; 3 <= length <= 50, unique
// - password : string ; 8 <= length <= 50, unique
// - email : string ; valid email, unique
// - role : string ; one of admin, super-admin, user
//
// Response
// - 201 and user id
// - 400 bad request ; or input invalid
//   - Email already exsits => message `Email already exists`
//   - Username already exsits => message `Username already exists`
//
// - 409 conflict ; no permission when requester is not super-admin/admin
// - 500 internal server error
func (h *usersHttpHandler) RegisterUser(c echo.Context) error {
	log.Println("RegisterUser : Starting handler")
	reqBody := new(models.RegisterDto)

	if err := c.Bind(reqBody); err != nil {
		log.Println("RegisterUser : Error while binding request body: ", err)
		return baseResponse(c, http.StatusBadRequest, messages.BAD_REQUEST)

	}

	requesterRole := "user"
	var err error
	if reqBody.Role == "admin" || reqBody.Role == "super-admin" {
		requesterRole, err = jwt.GetRole(c)
		if err != nil {
			if er, ok := err.(*errors.RequestError); ok {
				log.Println("RegisterUser : Error while validating token: ", er.Error())
				return baseResponse(c, er.StatusCode, er.Error())
			} else {
				log.Println("RegisterUser : Error while validating token: ", err)
				return baseResponse(c, http.StatusUnauthorized, messages.UNAUTHORIZED)
			}
		}
	}

	log.Println("RegisterUser : request: ", reqBody.ToString())

	userID, err := h.usersUsecase.RegisterUser(requesterRole, reqBody)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			log.Println("RegisterUser : Error while registering user: ", er.Error())
			return baseResponse(c, er.StatusCode, er.Error())
		} else {
			log.Println("RegisterUser : Error while registering user: ", err)
			return baseResponse(c, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}
	}

	return registerResponse(c, http.StatusCreated, messages.SUCCESSFUL_REGISTER, userID)
}
