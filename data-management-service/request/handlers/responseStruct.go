package handlers

import (
	"data-management/request/models"
	"log/slog"
)

type ResponseOptions struct {
	Response         any
	OptionalResponse any
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

type RequestLog struct {
	RequestID string
	Status    string
}

type ArrayRequestsLog struct {
	Length int
}

type RequestResponse struct {
	RequestS []*models.Request `json:"request"`
	Amount   int               `json:"amount"`
}
