package jwt

import (
	"auth-service/config"
	"auth-service/errors"
	"auth-service/messages"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CustomClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// CreateToken generates a token for the given user ID, username, and role.
//
// Parameters:
// - userID: the ID of the user (string)
// - username: the username of the user (string)
// - role: the role of the user (string)
//
// Returns:
// - token: the generated token (string)
// - error: an error if the token generation fails (error)
func CreateToken(username, role string) (string, error) {
	claims := CustomClaims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	secretKey := config.GetConfig().App.JWTKey

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ValidateAndExtractClaims validates the JWT token in the Authorization header
// and extracts the claims.
func ValidateAndExtractClaims(c echo.Context) (*CustomClaims, *errors.RequestError) {
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return nil, errors.CreateError(http.StatusBadRequest, messages.MISSING_AUTHORIZATION)
	}

	claims, err := ValidateAndExtractToken(tokenString)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func ValidateAndExtractToken(tokenString string) (*CustomClaims, *errors.RequestError) {
	secretKey := config.GetConfig().App.JWTKey
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, errors.CreateError(http.StatusUnauthorized, messages.INVALID_TOKEN)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.CreateError(http.StatusUnauthorized, messages.INVALID_TOKEN)
	}

	return claims, nil
}
