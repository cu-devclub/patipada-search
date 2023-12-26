package models

import (
	"auth-service/users/helper"
)

type RegisterDto struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=8,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" validate:"required,oneof=super-admin admin user"`
}

func (r *RegisterDto) MockData() {
	r.Username = helper.GenerateRandomUsername()
	r.Password = "test-password"
	r.Email = helper.GenerateRandomEmail()
	r.Role = "user"
}

type LoginDto struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type ResetPassword struct {
	Token    string `json:"token"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type ForgetPassword struct {
	Email string `json:"email" validate:"required,email"`
}

type RemoveUserDto struct {
	Username string `json:"username"`
}

type ChangePassword struct {
	OldPassword string `json:"oldPassword" validate:"required,min=8,max=50"`
	NewPassword string `json:"newPassword" validate:"required,min=8,max=50"`
}