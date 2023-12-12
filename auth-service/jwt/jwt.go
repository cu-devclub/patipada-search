package jwt

import (
	"auth-service/config"
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
func CreateToken(userID, username, role string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
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
func ValidateAndExtractClaims(c echo.Context) (*CustomClaims, error) {
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Missing Authorization header")
	}

	secretKey := config.GetConfig().App.JWTKey

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid || err == jwt.ErrTokenMalformed {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	return claims, nil
}
