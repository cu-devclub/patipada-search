package entities

import "time"

type (
	InsertUsersDto struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Salt     string `json:"salt"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}

	Users struct {
		Id                  string    `json:"id"`
		Username            string    `json:"username"`
		Password            string    `json:"password"`
		Salt                string    `json:"salt"`
		Email               string    `json:"email"`
		Role                string    `json:"role"`
		Is_Active           bool      `json:"is_active"`
		ResetToken          string    `json:"reset_token"`
		ResetTokenExpiresAt time.Time `json:"reset_token_expires_at"`
	}

	Email struct {
		Subject     string
		Content     string
		To          []string
		CC          []string
		BCC         []string
		AttachFiles []string
	}


)

func (u *Users) ToString() string{
	return u.Id + u.Username  + u.Email + u.Role
}
