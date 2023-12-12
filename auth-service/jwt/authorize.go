package jwt 
import (
	"github.com/labstack/echo/v4"
	"net/http"
)
// Authorize authorizes the request by validating and extracting claims from the provided echo.Context.
//
// It returns an error if the validation fails, otherwise it returns the role of the claims in a JSON response.
func Authorize (c echo.Context) error {
	claims, err := ValidateAndExtractClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
		return err
	}
	c.JSON(http.StatusOK, map[string]string{
		"role": claims.Role,
	})
	return nil
}