package helper

import "auth-service/users/entities"

// GetUserFromUserLists returns a user from a list of users based on the username.
//
// It takes in a slice of user entities and a string representing the username as parameters.
// It returns a pointer to a user entity or nil if no user is found.
func GetUserFromUserLists(users []*entities.Users, username string) *entities.Users {
	for _, user := range users {
		if user.Username == username {
			return user
		}
	}
	return nil
}
