package usecases

import (
	"auth-service/errors"
	"auth-service/jwt"
	"auth-service/messages"
	"auth-service/users/models"
)

// Remove user by username & requestor role must be higher
// Parameters  :
// - requester Role (string) ; one of admin, super-admin, user
// - models.RemoveUserDto
//   - username (string)
//
// Response
// - 200 OK
// - 400 bad request (invalid/missing username)
// - 401 Unauthorize ; missing token
// - 403 Forbidden ; no permission
// - 404 User not found (invalid username)
// - 500 internal server error
func (u *UsersUsecaseImpl) RemoveUser(requesterRole string, in *models.RemoveUserDto) error {
	// Get User data
	user, err := u.usersRepository.GetUserByUsername(in.Username)
	if err != nil {
		return errors.CreateError(404, messages.USER_NOT_FOUND)
	}

	// Check roles
	ch, _ := jwt.HasAuthorizeRole(requesterRole, user.Role,false)
	if !ch {
		return errors.CreateError(403, messages.NO_PERMISSION)
	}

	// Remove User
	return u.usersRepository.RemoveUser(in.Username)
}
