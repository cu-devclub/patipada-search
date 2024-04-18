package usecases

import (
	"auth-service/config"
	"auth-service/errors"
	"auth-service/messages"
	"auth-service/users/entities"
	"auth-service/users/helper"
	"auth-service/users/models"
	"fmt"
	"log/slog"
	"time"

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

// Request the link to reset password
// Link when sent to input email if valid
// Parameter(JSON)
// - email (string,email)

// Response
// - 200 OK & reset password token (also send to email)
// - 400 bad request (invalid email format)
// - 404 User not found (email not exists)
// - 500 internal server error
func (u *UsersUsecaseImpl) ForgetPassword(in *models.ForgetPassword) (string, error) {
	// Validate data
	validator := validator.New()
	if err := validator.Struct(in); err != nil {
		return "", errors.CreateError(400, err.Error())
	}

	// Check if email exists
	users, err := u.usersRepository.GetAllUsersData()
	if err != nil {
		return "", err
	}
	user := &entities.Users{}
	for _, u := range users {
		if u.Email == in.Email {
			user = u
		}
	}

	if helper.GetUserFromUserLists(users, user.Username) == nil {
		return "", errors.CreateError(404, messages.USER_NOT_FOUND)
	}

	// if current token is not expire yet
	// extend the time from now on and not change the token also not send the email
	if user.ResetTokenExpiresAt.After(time.Now()) {
		user.ResetTokenExpiresAt = time.Now().Add(15 * time.Minute)
		err = u.usersRepository.UpdateUser(user)
		if err != nil {
			return "", err
		}
		return user.ResetToken, nil
	}

	//Generate Token
	resetPasswordToken := helper.GenerateResetToken()
	resetPasswordExpireTime := helper.GenerateResetTokenExpiration()

	user.ResetToken = resetPasswordToken
	user.ResetTokenExpiresAt = resetPasswordExpireTime
	err = u.usersRepository.UpdateUser(user)
	if err != nil {
		return "", err
	}

	slog.Info("Update reset token for", slog.String("username", user.Username))
	// Sending email
	slog.Info("Sending email to", slog.String("email", in.Email))

	cfg := config.GetConfig()
	subject := "Reset Password"
	resetLink := fmt.Sprintf("%s/user/reset-password/%s", cfg.Link.URL, resetPasswordToken)
	content := fmt.Sprintf(`
	<h1>Hello %s </h1>
	<p>We received a request to reset your password. Click the link below to reset your password:</p>
	<p><a href="%s">Reset Your Password</a></p>
	<p>If you didn't request a password reset, you can ignore this email.</p>
`, user.Username, resetLink)
	to := []string{in.Email}
	attachFiles := []string{}

	e := &entities.Email{
		Subject:     subject,
		Content:     content,
		To:          to,
		AttachFiles: attachFiles,
		CC:          nil,
		BCC:         nil,
	}

	err = u.UserEmailing.SendEmail(e)
	if err != nil {
		return "", errors.CreateError(500, fmt.Sprintf("Error sending email: %v", err))
	}

	return resetPasswordToken, err
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
