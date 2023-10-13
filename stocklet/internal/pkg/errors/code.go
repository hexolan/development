package errors

type ErrorCode int32

const (
	ErrCodeUnknown ErrorCode = iota

	ErrCodeService
	ErrCodeExtService

	ErrCodeNotFound
	ErrCodeAlreadyExists
	
	ErrCodeForbidden
	ErrCodeUnauthorised
	
	ErrCodeInvalidArgument
)