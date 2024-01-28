package usecases

import (
	"auth-service/errors"
	"auth-service/jwt"
	"auth-service/messages"
	"auth-service/users/helper"
	"auth-service/users/models"
	"log"

	"github.com/go-playground/validator"
)

// Authentication
// Parameters (models.LoginDto) :
// - username : string ; 3 <= length <= 50
// - password : string ; 8 <= length <= 50
//
// Response
// - 200 , role and token
// - 400 bad request ; some field missing or input invalid
// - 401 unauthorized ;  username or password incorrect
// - 500 internal server error
func (u *UsersUsecaseImpl) Authentication(in *models.LoginDto) (string, string, error) {
	// Validate data
	validator := validator.New()
	if err := validator.Struct(in); err != nil {
		log.Println("Authentication : Error while validating request body: ", err)
		return "", "", errors.CreateError(400, err.Error())
	}

	user, err := u.usersRepository.GetUserByUsername(in.Username)
	if err != nil {
		log.Println("Authentication : Error while getting user by username: ", err)
		return "", "", errors.CreateError(401, messages.WRONG_USERNAME_PASSWORD)
	}

	if err := helper.VerifyPassword(user.Password, in.Password+user.Salt); err != nil {
		log.Println("Authentication : Error while verifying password: ", err)
		return "", "", errors.CreateError(401, messages.WRONG_USERNAME_PASSWORD)
	}

	token, err := jwt.CreateToken(user.Username, user.Role)
	if err != nil {
		log.Println("Authentication : Error while creating token: ", err)
		return "", "", err
	}

	log.Println("Authentication : Login Success",user.Username, " & Token created: ", token)

	return token, user.Role, nil
}
