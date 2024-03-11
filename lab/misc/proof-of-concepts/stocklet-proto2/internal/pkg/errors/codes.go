// Copyright 2024 Declan Teevan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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