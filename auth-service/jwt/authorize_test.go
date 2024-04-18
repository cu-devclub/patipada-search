package jwt

import (
	"auth-service/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasAuthorizeRole(t *testing.T) {
	config.InitializeViper("../")
	config.ReadConfig()

	t.Run("AllowEqualRole = true", func(t *testing.T) {
		requesterRole := "user"
		requiredRole := "admin"
		allowEqualRole := true

		result := HasAuthorizeRole(requesterRole, requiredRole, allowEqualRole)

		assert.False(t, result)
	})

	t.Run("AllowEqualRole = false, requesterRole < requiredRole", func(t *testing.T) {
		requesterRole := "user"
		requiredRole := "admin"
		allowEqualRole := false

		result := HasAuthorizeRole(requesterRole, requiredRole, allowEqualRole)

		assert.False(t, result)
	})

	t.Run("AllowEqualRole = false, requesterRole = requiredRole", func(t *testing.T) {
		requesterRole := "admin"
		requiredRole := "admin"
		allowEqualRole := false

		result := HasAuthorizeRole(requesterRole, requiredRole, allowEqualRole)

		assert.False(t, result)
	})

	t.Run("AllowEqualRole = false, requesterRole > requiredRole", func(t *testing.T) {
		requesterRole := "admin"
		requiredRole := "user"
		allowEqualRole := false

		result := HasAuthorizeRole(requesterRole, requiredRole, allowEqualRole)

		assert.True(t, result)
	})
}
