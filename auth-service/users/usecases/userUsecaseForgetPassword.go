package usecases

import (
	"auth-service/config"
	"auth-service/errors"
	"auth-service/messages"
	"auth-service/users/entities"
	"auth-service/users/helper"
	"auth-service/users/models"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator"
)

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
		log.Println("Forget Password : Error while validating request body: ", err)
		return "", errors.CreateError(400, err.Error())
	}

	// Check if email exists
	users, err := u.usersRepository.GetAllUsersData()
	if err != nil {
		log.Println("Forget Password : Error while getting all users data: ", err)
		return "", err
	}
	user := &entities.Users{}
	for _, u := range users {
		if u.Email == in.Email {
			user = u
		}
	}

	if helper.GetUserFromUserLists(users, user.Username) == nil {
		log.Println("Forget Password : not found user from request email")
		return "", errors.CreateError(404, messages.USER_NOT_FOUND)
	}

	// if current token is not expire yet
	// extend the time from now on and not change the token also not send the email
	if user.ResetTokenExpiresAt.After(time.Now()) {
		user.ResetTokenExpiresAt = time.Now().Add(15 * time.Minute)
		log.Println("Forget Password : Token is not expire yet,  Updating user (token time): ", user)
		err = u.usersRepository.UpdateUser(user)
		if err != nil {
			log.Println("Forget Password : Error while updating user(token time): ", err)
			return "", err
		}
		return user.ResetToken, nil
	}

	log.Println("Token already expired, Generate new token .....")
	//Generate Token
	resetPasswordToken := helper.GenerateResetToken()
	resetPasswordExpireTime := helper.GenerateResetTokenExpiration()

	user.ResetToken = resetPasswordToken
	user.ResetTokenExpiresAt = resetPasswordExpireTime
	log.Println("Forget Password : Updating user(token time): ", user)
	err = u.usersRepository.UpdateUser(user)
	if err != nil {
		log.Println("Forget Password : Error while updating user(token time): ", err)
		return "", err
	}

	// Sending email
	log.Println("Sending email to ", in.Email," ......")
	
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
		log.Println("Forget Password : Error while sending email: ", err)
		return "", err
	}

	return resetPasswordToken, err
}
