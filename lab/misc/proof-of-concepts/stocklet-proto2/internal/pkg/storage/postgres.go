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

package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hexolan/stocklet/internal/pkg/config"
	"github.com/hexolan/stocklet/internal/pkg/errors"
)

func NewPostgresConn(conf *config.PostgresConfig) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	pCl, err := pgxpool.New(ctx, conf.GetDSN())
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to connect to postgres", err)
	}
	return pCl, nil
}
