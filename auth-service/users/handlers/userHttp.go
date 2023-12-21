package handlers

import (
	"net/http"

	"auth-service/errors"
	"auth-service/jwt"
	"auth-service/messages"
	"auth-service/users/models"
	"auth-service/users/usecases"

	"github.com/labstack/echo/v4"
)

type usersHttpHandler struct {
	usersUsecase usecases.UsersUsecase
}

func NewUsersHttpHandler(usersUsecase usecases.UsersUsecase) UsersHandler {
	return &usersHttpHandler{
		usersUsecase: usersUsecase,
	}
}

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
// - 401 Unauthorized ;  no token attach
// - 409 conflict ; no permission when requester is not super-admin/admin
// - 500 internal server error
func (h *usersHttpHandler) RegisterUser(c echo.Context) error {
	reqBody := new(models.RegisterDto)

	if err := c.Bind(reqBody); err != nil {
		return baseResponse(c, http.StatusBadRequest, messages.BAD_REQUEST)
	}

	requesterRole := "user"
	var err error
	if reqBody.Role == "admin" || reqBody.Role == "super-admin" {
		requesterRole, err = jwt.GetRole(c)
		if err != nil {
			return baseResponse(c, http.StatusUnauthorized, messages.UNAUTHORIZED)
		}
	}

	userID, err := h.usersUsecase.RegisterUser(requesterRole, reqBody)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return baseResponse(c, er.StatusCode, er.Error())
		} else {
			return baseResponse(c, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}
	}

	return registerResponse(c, http.StatusCreated, messages.SUCCESSFUL_REGISTER, userID)
}

// Login handles the login request.
//
// It takes in a `c` parameter of type `echo.Context`
// Parameters (JSON) :
// - username : string ; 3 <= length <= 50
// - password : string ; 8 <= length <= 50
//
// Response
// - 200 , role and token
// - 400 bad request ; some field missing or input invalid
// - 401 unauthorized ;  username or password incorrect
// - 500 internal server error
func (h *usersHttpHandler) Login(c echo.Context) error {
	reqBody := new(models.LoginDto)
	if err := c.Bind(reqBody); err != nil {
		return loginResponse(c, http.StatusBadRequest, messages.BAD_REQUEST, "", "")
	}
	token, role, err := h.usersUsecase.Authentication(reqBody)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return loginResponse(c, er.StatusCode, er.Error(), "", "")
		} else {
			return loginResponse(c, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR, "", "")
		}
	}
	return loginResponse(c, http.StatusOK, messages.SUCCESSFUL_LOGIN, token, role)

}

// Request the link to reset password
// Link when sent to input email if valid
// Route Parameter
// - email (string,email)

// Response
// - 200 OK & reset password token (also send to email)
// - 400 bad request (invalid email)
// - 404 User not found (email not exists)
// - 500 internal server error
func (h *usersHttpHandler) ForgetPassword(c echo.Context) error {
	email := c.Param("email")
	if email == "" {
		return forgetPasswordResponse(c, http.StatusBadRequest, messages.BAD_REQUEST, "")
	}

	in := &models.ForgetPassword{
		Email: email,
	}

	token, err := h.usersUsecase.ForgetPassword(in)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return forgetPasswordResponse(c, er.StatusCode, er.Error(), "")
		} else {
			return forgetPasswordResponse(c, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR, "")
		}
	}
	return forgetPasswordResponse(c, http.StatusOK, messages.SUCCESSFUL_SEND_EMAIL_FORGET_PASSWORD, token)
}

// Reset Password
// Parameters(JSON)
// - token (string) ; reset password token
// - password (string) ; new password ; 8 <= length <= 50
//
// Response
// - 201 Created ; Update password success
// - 400 bad request (invalid format password)
// - 401 Unautorize ; invalid reset password
// - 422 ; New password == Old password
// - 500 internal server error
func (h *usersHttpHandler) ResetPassword(c echo.Context) error {
	reqBody := new(models.ResetPassword)
	if err := c.Bind(reqBody); err != nil {
		return baseResponse(c, http.StatusBadRequest, messages.BAD_REQUEST)
	}

	if err := h.usersUsecase.ResetPassword(reqBody); err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return baseResponse(c, er.StatusCode, er.Error())
		} else {
			return baseResponse(c, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}
	}
	return baseResponse(c, http.StatusCreated, messages.SUCCESSFUL_RESET_PASSWORD)

}

// Remove user by username & requestor role must be higher
// Header - Authorization : <token>
// Parameters (Route Param) :
// - username (string)
//
// Response
// - 200 OK
// - 400 bad request (invalid/missing username)
// - 401 Unauthorize ; missing token
// - 403 Forbidden ; no permission
// - 404 User not found (invalid username)
// - 500 internal server error
func (h *usersHttpHandler) RemoveUser(c echo.Context) error {
	username := c.Param("username")
	if username == "" {
		return baseResponse(c, http.StatusBadRequest, messages.BAD_REQUEST)
	}
	reqBody := &models.RemoveUserDto{
		Username: username,
	}

	requesterRole, err := jwt.GetRole(c)
	if err != nil {
		return baseResponse(c, http.StatusUnauthorized, messages.UNAUTHORIZED)
	}

	if err := h.usersUsecase.RemoveUser(requesterRole, reqBody); err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return baseResponse(c, er.StatusCode, er.Error())
		} else {
			return baseResponse(c, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR)
		}
	}
	return baseResponse(c, http.StatusOK, messages.SUCCESSFUL_REMOVE_USER)

}

// Verify Reset Token to verify the time valid of token (15 minute)
// Route Params - `token`
//
// Response
// - 200 OK & result (true/false)
// - 404 Not found ; token == "" or not attach token
// - 500 internal server error
func (h *usersHttpHandler) VerifyResetToken(c echo.Context) error {
	token := c.Param("token")
	ch, err := h.usersUsecase.VerifyResetToken(token)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			return verifyTokenResponse(c, er.StatusCode, er.Error(), false)
		} else {
			return verifyTokenResponse(c, http.StatusInternalServerError, messages.INTERNAL_SERVER_ERROR, false)
		}
	}
	return verifyTokenResponse(c, http.StatusOK, messages.SUCCESSFUL_VERIFY_RESET_TOKEN, ch)
}
