package jwt

import (
	"auth-service/config"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// ValidateToken is a function that validates the token present in the Authorization header of a request.
//
// It takes a `next` echo.HandlerFunc as input and returns an echo.HandlerFunc as output.
// The echo.HandlerFunc passed as `next` will be executed if the token is valid.
// If the token is invalid or missing, an appropriate error response will be returned.
func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Missing Authorization header",
			})
		}
		secretKey := config.GetConfig().App.JWTKey

		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil // Replace with your actual secret key
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid || err == jwt.ErrTokenMalformed {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid token",
				})
			}
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		claims, ok := token.Claims.(*CustomClaims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid token",
			})
		}

		c.Set("user", claims) // Store claims in context
		return next(c)
	}
}

// ClaimToken retrieves the user claims from the echo.Context object.
//
// It takes a single parameter, c, of type echo.Context.
// It returns a pointer to CustomClaims and an error.
func ClaimToken(c echo.Context) (*CustomClaims, error) {
	claims, ok := c.Get("user").(*CustomClaims)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Unable to retrieve user claims")
	}
	return claims, nil
}
