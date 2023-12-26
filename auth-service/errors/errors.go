package errors

import (
	"errors"
	"fmt"
)

type RequestError struct {
	StatusCode int
	Err        error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("%v",  r.Err)
}

func CreateError(statusCode int, errMessage string) *RequestError {
	return &RequestError{
		StatusCode: statusCode,
		Err:        errors.New(errMessage),
	}
}
