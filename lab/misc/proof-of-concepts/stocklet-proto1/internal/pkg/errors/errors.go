package errors

import (
	"fmt"
)

type ServiceError struct {
	code    ErrorCode
	msg     string
	wrapped error
}

func NewServiceError(code ErrorCode, msg string) error {
	return &ServiceError{
		code: code,
		msg:  msg,
	}
}

func NewServiceErrorf(code ErrorCode, msg string, args ...interface{}) error {
	return NewServiceError(code, fmt.Sprintf(msg, args...))
}

func WrapServiceError(code ErrorCode, msg string, wrapped error) error {
	return &ServiceError{
		code:    code,
		msg:     msg,
		wrapped: wrapped,
	}
}

func (e ServiceError) Error() string {
	if e.wrapped != nil {
		return fmt.Sprintf("%s: %s", e.msg, e.wrapped.Error())
	}

	return e.msg
}

func (e ServiceError) Code() ErrorCode {
	return e.code
}

func (e ServiceError) Unwrap() error {
	return e.wrapped
}
