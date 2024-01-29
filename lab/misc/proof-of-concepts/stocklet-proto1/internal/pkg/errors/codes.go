package errors

import (
	"google.golang.org/grpc/codes"
)

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

// Maps the custom service error codes 
// to their gRPC status code equivalents.
func (c ErrorCode) GRPCCode() codes.Code {
	codeMap := map[ErrorCode]codes.Code{
		ErrCodeUnknown: codes.Unknown,

		ErrCodeService: codes.Internal,
		ErrCodeExtService: codes.Unavailable,

		ErrCodeNotFound: codes.NotFound,
		ErrCodeAlreadyExists: codes.AlreadyExists,

		ErrCodeForbidden: codes.PermissionDenied,
		ErrCodeUnauthorised: codes.PermissionDenied,

		ErrCodeInvalidArgument: codes.InvalidArgument,
	}

	grpcCode, mapped := codeMap[c]
	if mapped {
		return grpcCode
	}
	return codes.Unknown
}