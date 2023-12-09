package usecases

import "auth-service/users/models"

type UsersUsecase interface {
	// UsersRegisterDataProcessing processes the registration data for users.
	//
	// Parameter(s):
	// - in: the data for adding users, including username, password, email, and role.
	//
	// Return type(s):
	// - if username and email already exists return error with status code 409 Conflict
	// - error: any error that occurred during the processing.
	UsersRegisterDataProcessing(in *models.AddUsersData) error

	// Authentication performs user authentication.
	//
	// It takes a LoginDto as input and returns a token(string)and an error.
	// If the user is not found or the username/password is incorrect, it returns an error of 401.
	Authentication(in *models.LoginDto) (string, error)
}
