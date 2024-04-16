package helper

import (
	"auth-service/users/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserFromUserLists(t *testing.T) {
	users := []*entities.Users{
		{Username: "user1"},
		{Username: "user2"},
		{Username: "user3"},
	}

	t.Run("Existing User", func(t *testing.T) {
		username := "user2"
		expectedUser := &entities.Users{Username: "user2"}

		result := GetUserFromUserLists(users, username)

		assert.Equal(t, expectedUser, result)
	})

	t.Run("Non-Existing User", func(t *testing.T) {
		username := "user4"

		result := GetUserFromUserLists(users, username)

		assert.Nil(t, result)
	})
}
