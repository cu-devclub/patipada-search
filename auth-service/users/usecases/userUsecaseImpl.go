package usecases

import (
	"auth-service/errors"
	"auth-service/jwt"
	"auth-service/users/entities"
	"auth-service/users/helper"
	"auth-service/users/messages"
	"auth-service/users/models"
	"auth-service/users/repositories"

	"github.com/go-playground/validator/v10"
)

type usersUsecaseImpl struct {
	usersRepository repositories.UsersRepository
}

func NewUsersUsecaseImpl(usersRepository repositories.UsersRepository) UsersUsecase {
	return &usersUsecaseImpl{
		usersRepository: usersRepository,
	}
}

// UsersRegisterDataProcessing processes the registration data for users.
//
// Parameter(s):
// - in: the data for adding users, including username, password, email, and role.
//
// Return type(s):
// - if username or email already exists return error with status code 409 Conflict
// - error: any error that occurred during the processing.
func (u *usersUsecaseImpl) UsersRegisterDataProcessing(in *models.AddUsersData) error {
	// Validate data
	validator := validator.New()
	if err := validator.Struct(in); err != nil {
		return errors.CreateError(400, err.Error())
	}

	// Check if username or email already exists
	users, err := u.usersRepository.GetAllUsersData()
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Username == in.Username || user.Email == in.Email {
			return errors.CreateError(409, messages.USERNAME_ALREADY_EXISTS)
		}
	}

	// INSERT USER
	uuid, err := helper.GenerateUUID()
	if err != nil {
		return err
	}

	password, salt, err := helper.GenerateHashedSaltedPassword(in.Password)
	if err != nil {
		return err
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
		return err
	}
	return nil
}

// Authentication performs user authentication.
//
// It takes a LoginDto as input and returns a token(string)and an error.
// If the user is not found or the username/password is incorrect, it returns an error of 401.
func (u *usersUsecaseImpl) Authentication(in *models.LoginDto) (string, error) {
	users, err := u.usersRepository.GetAllUsersData()
	if err != nil {
		return "", err
	}
	user := helper.GetUserFromUserLists(users, in.Username)
	if user == nil {
		return "", errors.CreateError(401, messages.WRONG_USERNAME_PASSWORD)
	}

	if err := helper.VerifyPassword(user.Password, in.Password+user.Salt); err != nil {
		return "", errors.CreateError(401, messages.WRONG_USERNAME_PASSWORD)
	}

	token, err := jwt.CreateToken(user.Id, user.Username, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
