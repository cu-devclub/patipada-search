package usecases

import (
	"auth-service/errors"
	"auth-service/messages"
	"auth-service/users/entities"
	"auth-service/users/helper"
	"auth-service/users/models"
	"time"

	"github.com/go-playground/validator"
)

// Reset Password
// Parameters(JSON)
// - token (string) ; reset password token
// - password (string) ; new password ; 8 <= length <= 50
//
// Response
// - 201 Created ; Update password success
// - 400 bad request (invalid format password)
// - 401 Unautorize ; invalid reset password
// - 500 internal server error
func (u *UsersUsecaseImpl) ResetPassword(in *models.ResetPassword) error {

	// Validate data (password)
	validator := validator.New()
	if err := validator.Struct(in); err != nil {
		return errors.CreateError(400, err.Error())
	}

	if in.Token == "" {
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	// Check token exists and valid
	users, err := u.usersRepository.GetAllUsersData()
	if err != nil {
		return err
	}
	user := &entities.Users{}
	ch := false
	for _, u := range users {
		if u.ResetToken == in.Token {
			if time.Now().Before(u.ResetTokenExpiresAt) {
				user = u
				ch = true
			}
		}
	}
	if !ch {
		return errors.CreateError(401, messages.UNAUTHORIZED)
	}

	password, salt, err := helper.GenerateHashedSaltedPassword(in.Password)
	if err != nil {
		return err
	}

	user.Password = password
	user.Salt = salt
	user.ResetToken = ""

	return u.usersRepository.UpdateUser(user)
}
