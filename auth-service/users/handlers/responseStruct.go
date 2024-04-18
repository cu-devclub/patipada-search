package handlers

import (
	"auth-service/users/models"
	"log/slog"
)

type ResponseOptions struct {
	Response            any
	LogResponseOptional any
}

type baseResponseStruct struct {
	Message string `json:"message"`
}

type errorResponseStruct struct {
	ErrMessage string `json:"errMessage"`
}

type loginResponseStruct struct {
	Token   string `json:"token"`
	Role    string `json:"role"`
	Message string `json:"message"`
}

type loginResponseLogStruct struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (r loginResponseStruct) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("Message", r.Message),
	)
}

type registerResponseStruct struct {
	Message string `json:"message"`
	UserID  string `json:"user_id"`
}

type verifyStruct struct {
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

type forgetPasswordStruct struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type forgetPasswordResLogStruct struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

type getAllUserResponse struct {
	Amount  int             `json:"amount"`
	Users   []*models.Users `json:"users"`
	Message string          `json:"message"`
}

type getAllUserLogResponse struct {
	Amount  int    `json:"amount"`
	Message string `json:"message"`
}

type Response struct {
	Code int
	Body any
}

func (r Response) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("Code", r.Code),
		slog.Any("Body", r.Body),
	)
}
