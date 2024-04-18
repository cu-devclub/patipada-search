package usecases

import (
	"auth-service/config"
	"auth-service/errors"
	"auth-service/jwt"
	"auth-service/messages"
	"auth-service/users/entities"
	"auth-service/users/helper"
	"auth-service/users/models"
	"sort"

	"github.com/go-playground/validator"
)

// RegisterUser
// If new user role is "admin" or "super-admin"
// then requester role must be "admin" or "super-admin"
// Parameters (JSON) :
// - requesterRole : string ; one of admin, super-admin, user
// - models.RegisterDto
//   - username : string ; 3 <= length <= 50, unique
//   - password : string ; 8 <= length <= 50, unique
//   - email : string ; valid email, unique
//   - role : string ; one of admin, super-admin, user
//
// Response
// - 201 and user id
// - 400 bad request ; or input invalid
//   - Email already exsits => message `Email already exists`
//   - Username already exsits => message `Username already exists`
//
// - 409 conflict ; no permission when requester is not super-admin/admin
// - 500 internal server error
func (u *UsersUsecaseImpl) RegisterUser(requesterRole string, in *models.RegisterDto) (string, error) {
	// Validate data
	validator := validator.New()
	if err := validator.Struct(in); err != nil {
		return "", errors.CreateError(400, err.Error())
	}

	// Check roles
	if in.Role == "admin" || in.Role == "super-admin" {
		ch := jwt.HasAuthorizeRole(requesterRole, in.Role, true)
		if !ch {
			return "", errors.CreateError(409, messages.NO_PERMISSION)
		}
	}

	// Check if username or email already exists
	users, err := u.usersRepository.GetAllUsersData()
	if err != nil {
		return "", err
	}
	for _, user := range users {
		if user.Username == in.Username {
			return "", errors.CreateError(400, messages.USERNAME_ALREADY_EXISTS)
		}
		if user.Email == in.Email {
			return "", errors.CreateError(400, messages.EMAIL_ALREADY_EXISTS)
		}
	}

	// INSERT USER
	uuid, err := helper.GenerateUUID()
	if err != nil {
		return "", err
	}

	password, salt, err := helper.GenerateHashedSaltedPassword(in.Password)
	if err != nil {
		return "", err
	}

	insertUsersData := &entities.InsertUsersDto{
		Id:       uuid,
		Username: in.Username,
		Password: password,
		Salt:     salt,
		Email:    in.Email,
		Role:     in.Role,
	}

	if err := u.usersRepository.InsertUsersData(insertUsersData); err != nil {
		return "", err
	}
	return uuid, nil
}

func (u *UsersUsecaseImpl) GetAllUsersData() ([]*models.Users, error) {
	en, err := u.usersRepository.GetAllUsersData()
	if err != nil {
		return nil, err
	}

	var users []*models.Users
	for _, user := range en {
		users = append(users, &models.Users{
			ID:       user.Id,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		})
	}

	cfg := config.GetConfig()
	sort.Slice(users, func(i, j int) bool {
		roleI := cfg.App.RolesMap[users[i].Role]
		roleJ := cfg.App.RolesMap[users[j].Role]
		if roleI != roleJ {
			return roleI > roleJ
		}

		return users[i].Username < users[j].Username
	})
	
	return users, nil
}

// Remove user by id & requestor role must be higher
// Parameters  :
// - requester Role (string) ; one of admin, super-admin, user
// - models.RemoveUserDto
//   - id (string)
//
// Response
// - 200 OK
// - 400 bad request (invalid/missing id)
// - 401 Unauthorize ; missing token
// - 403 Forbidden ; no permission
// - 404 User not found (invalid id)
// - 500 internal server error
func (u *UsersUsecaseImpl) RemoveUser(requesterRole string, in *models.RemoveUserDto) error {
	// Get User data
	users, err := u.usersRepository.GetAllUsersData()
	if err != nil {
		return err
	}

	var user *entities.Users
	for _, u := range users {
		if u.Id == in.ID {
			user = u
			break
		}
	}
	if user == nil {
		return errors.CreateError(404, messages.USER_NOT_FOUND)
	}

	// Check roles
	ch := jwt.HasAuthorizeRole(requesterRole, user.Role, false)
	if !ch {
		return errors.CreateError(403, messages.NO_PERMISSION)
	}

	// Remove User
	return u.usersRepository.RemoveUser(in.ID)
}
