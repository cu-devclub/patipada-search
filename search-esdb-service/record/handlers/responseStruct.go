package handlers

import "log/slog"

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

type SearchRecordLogResponse struct {
	Length int
	Status string
	Amount int
}

type RecordIndexLogResponse struct {
	Index string
}
