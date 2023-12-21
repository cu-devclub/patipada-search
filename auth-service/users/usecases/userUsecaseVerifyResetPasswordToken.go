package usecases

import (
	"time"
)

// Verify Reset Token to verify the time valid of token (15 minute)
// Parameter
// - token ; string
//
// Response
// - 200 OK & result (true/false)
// - 404 Not found ; token == "" or not attach token
// - 500 internal server error
func (u *UsersUsecaseImpl) VerifyResetToken(token string) (bool, error) {

	users, err := u.usersRepository.GetAllUsersData()
	if err != nil {
		return false, err
	}

	for _, user := range users {
		if user.ResetToken == token {
			if time.Now().Before(user.ResetTokenExpiresAt) {
				return true, nil
			}
		}
	}
	return false, nil
}
