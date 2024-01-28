package usecases

import (
	"auth-service/errors"
	"auth-service/messages"
	"auth-service/users/helper"
	"auth-service/users/models"

	"github.com/go-playground/validator"
)

// Change Password
// Parameter(JSON)
// - oldPassword (string) ; old password ; 8 <= length <= 50
// - newPassword (string) ; new password ; 8 <= length <= 50
//
// Response
// - 200 OK ; Update password success
// - 400 bad request (invalid format password)
// - 401 Unautorize ; invalid old password
// - 500 internal server error
func (u *UsersUsecaseImpl) ChangePassword(in *models.ChangePassword, username string) error {
	validator := validator.New()
	if err := validator.Struct(in); err != nil {
		return errors.CreateError(400, messages.BAD_REQUEST)
	}

	// Check if old password is correct
	user, err := u.usersRepository.GetUserByUsername(username)
	if err != nil {
		return errors.CreateError(401, messages.BAD_REQUEST)
	}

	if err := helper.VerifyPassword(user.Password, in.OldPassword+user.Salt); err != nil {
		return errors.CreateError(401, messages.BAD_REQUEST)
	}

	// Generate new credentials
	password, salt, err := helper.GenerateHashedSaltedPassword(in.NewPassword)
	if err != nil {
		return err
	}

	user.Password = password
	user.Salt = salt

	if err := u.usersRepository.UpdateUser(user); err != nil {
		return err
	}

	return nil
}
