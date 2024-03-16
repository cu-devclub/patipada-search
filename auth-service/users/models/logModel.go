package models

type LoginLogDto struct {
	Username string `json:"username"`
}

type RegisterLogDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type ResetPasswordLog struct {
	Token string `json:"token"`
}
