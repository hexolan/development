// Copyright (C) 2024 Declan Teevan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package controller

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hexolan/stocklet/internal/svc/auth"
	// pb "github.com/hexolan/stocklet/internal/pkg/protogen/auth/v1"
)

type postgresController struct {
	cl *pgxpool.Pool
}

func NewPostgresController(cl *pgxpool.Pool) auth.StorageController {
	return postgresController{cl: cl}
}

// todo: implement
func (c postgresController) SetPassword(ctx context.Context, userId string, password string) error {
	unimplemented
	return nil
}

// todo: implement
func (c postgresController) VerifyPassword(ctx context.Context, userId string, password string) (bool, error) {
	return false, nil
}

// todo: implement
func (c postgresController) DeleteAuthMethods(ctx context.Context, userId string) error {
	return nil
}
