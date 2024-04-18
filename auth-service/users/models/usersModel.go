package models

import (
	"auth-service/users/helper"
	"fmt"
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

func (r *RegisterDto) ToString() string {
	return fmt.Sprintf("Username: %s, Email: %s, Role: %s", r.Username, r.Email, r.Role)
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
	ID string `json:"id"`
}

type ChangePassword struct {
	OldPassword string `json:"oldPassword" validate:"required,min=8,max=50"`
	NewPassword string `json:"newPassword" validate:"required,min=8,max=50"`
}

type Users struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
