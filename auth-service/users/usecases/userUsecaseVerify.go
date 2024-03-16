package usecases

import (
	"auth-service/jwt"
	"time"

	"github.com/labstack/echo/v4"
)

// Verify Token to verify the time valid of auth token
// Header - Authorization : <token>
//
// Response
// - 200 OK & result (true/false)
// - 400 Bad request ; missing token
// - 401 Unauthorize ; invalid token
// - 500 internal server error
func (u *UsersUsecaseImpl) VerifyToken(c echo.Context) (bool, error) {
	_, err := jwt.ValidateAndExtractClaims(c)
	if err != nil {
		return false, err
	}
	return true, nil
}

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

// * Called by GRPC
func (s *UsersUsecaseImpl) VerifyUsername(username string) (bool, error) {
	_, err := s.usersRepository.GetUserByUsername(username)
	if err != nil {
		return false, err
	}
	return true, nil
}
