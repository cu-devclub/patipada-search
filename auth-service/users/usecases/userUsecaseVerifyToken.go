package usecases

import (
	"auth-service/jwt"

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
