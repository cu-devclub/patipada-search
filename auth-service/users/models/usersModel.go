package models

type AddUsersData struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=8,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" validate:"required,oneof=super-admin admin user"`
}

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

