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
	"fmt"

	"google.golang.org/grpc/status"
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

// Set the gRPC status to only expose the top error message.
//
// This is to prevent any full error contexts (from wrapped errors)
// being exposed to users by the gateway.
// e.g. "{"code":2,"message":"something went wrong scanning order: failed to connect to `host=postgres user=postgres database=postgres`: hostname resolving error (lookup postgres on 127.0.0.11:53: server misbehaving)","details":[]}"
//
func (e ServiceError) GRPCStatus() *status.Status {
	return status.New(e.Code().GRPCCode(), e.msg)
}

func (e ServiceError) Unwrap() error {
	return e.wrapped
}
