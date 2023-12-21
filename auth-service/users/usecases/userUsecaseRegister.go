package usecases

import (
	"auth-service/errors"
	"auth-service/jwt"
	"auth-service/messages"
	"auth-service/users/entities"
	"auth-service/users/helper"
	"auth-service/users/models"

	"github.com/go-playground/validator"
)

// RegisterUser
// If new user role is "admin" or "super-admin"
// then requester role must be "admin" or "super-admin"
// Parameters (JSON) :
// - requesterRole : string ; one of admin, super-admin, user
// - models.RegisterDto
//   - username : string ; 3 <= length <= 50, unique
//   - password : string ; 8 <= length <= 50, unique
//   - email : string ; valid email, unique
//   - role : string ; one of admin, super-admin, user
//
// Response
// - 201 and user id
// - 400 bad request ; or input invalid
// - 409 conflict ; no permission when requester is not super-admin/admin
// - 500 internal server error
func (u *UsersUsecaseImpl) RegisterUser(requesterRole string, in *models.RegisterDto) (string, error) {
	// Validate data
	validator := validator.New()
	if err := validator.Struct(in); err != nil {
		return "", errors.CreateError(400, err.Error())
	}

	// Check roles
	if in.Role == "admin" || in.Role == "super-admin" {
		ch, _ := jwt.HasAuthorizeRole(requesterRole, in.Role,true)
		if !ch {
			return "", errors.CreateError(409, messages.NO_PERMISSION)
		}
	}

	// Check if username or email already exists
	users, err := u.usersRepository.GetAllUsersData()
	if err != nil {
		return "", err
	}
	for _, user := range users {
		if user.Username == in.Username || user.Email == in.Email {
			return "", errors.CreateError(400, messages.USERNAME_ALREADY_EXISTS)
		}
	}

	// INSERT USER
	uuid, err := helper.GenerateUUID()
	if err != nil {
		return "", err
	}

	password, salt, err := helper.GenerateHashedSaltedPassword(in.Password)
	if err != nil {
		return "", err
	}

	insertUsersData := &entities.InsertUsersDto{
		Id:       uuid,
		Username: in.Username,
		Password: password,
		Salt:     salt,
		Email:    in.Email,
		Role:     in.Role,
	}

	if err := u.usersRepository.InsertUsersData(insertUsersData); err != nil {
		return "", err
	}
	return uuid, nil
}