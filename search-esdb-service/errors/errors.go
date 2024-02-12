package errors

type RequestError struct {
	StatusCode int
	Message    string
}

func (e *RequestError) Error() string {
	return e.Message
}

func CreateError(statusCode int, message string) *RequestError {
	if message == "" {
		return nil
	}
	return &RequestError{
		StatusCode: statusCode,
		Message:    message,
	}
}