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

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hexolan/stocklet/internal/svc/payment"
	"github.com/hexolan/stocklet/internal/pkg/errors"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/payment/v1"
)

const (
	pgTransactionBaseQuery string = "SELECT id, order_id, customer_id, amount, reversed_at, processed_at FROM transactions"
	pgCustomerBalanceBaseQuery string = "SELECT customer_id, balance FROM customer_balances"
)

type postgresController struct {
	cl *pgxpool.Pool
}

func NewPostgresController(cl *pgxpool.Pool) payment.StorageController {
	return postgresController{cl: cl}
}

func (c postgresController) GetBalance(ctx context.Context, customerId string) (*pb.CustomerBalance, error) {
	
}

func (c postgresController) GetTransaction(ctx context.Context, orderId string) (*pb.Transaction, error) {
	
}
	
func (c postgresController) CreateBalance(ctx context.Context, customerId string) error {
	
}

func (c postgresController) CreditBalance(ctx context.Context, customerId string, amount float32) error {
	
}

func (c postgresController) DebitBalance(ctx context.Context, customerId string, amount float32) error {
	
}

func (c postgresController) CloseBalance(ctx context.Context, customerId string) error {
	
}

func (c postgresController) PaymentForOrder(ctx context.Context, orderId string, customerId string, amount float32) error {
	
}

// Scan a postgres row to a protobuf object
func scanRowToTransaction(row pgx.Row) (*pb.Transaction, error) {
	var transaction pb.Transaction

	// Temporary variables that require conversion
	var tmpProcessedAt pgtype.Timestamp
	var tmpReversedAt pgtype.Timestamp

	err := row.Scan(
		&transaction.Id,
		&transaction.OrderId,
		&transaction.CustomerId,
		&transaction.Amount,
		&tmpReversedAt,
		&tmpProcessedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.WrapServiceError(errors.ErrCodeNotFound, "transaction not found", err)
		} else {
			return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to scan object from database", err)
		}
	}

	// Convert the temporary variables
	// - converting postgres timestamps to unix format
	if tmpProcessedAt.Valid {
		transaction.ProcessedAt = tmpProcessedAt.Time.Unix()
	} else {
		return nil, errors.NewServiceError(errors.ErrCodeUnknown, "failed to scan object from database (timestamp conversion)")
	}

	if tmpReversedAt.Valid {
		unixReversed := tmpReversedAt.Time.Unix()
		transaction.ReversedAt = &unixReversed
	}
	
	return &transaction, nil
}

// Scan a postgres row to a protobuf object
func scanRowToCustomerBalance(row pgx.Row) (*pb.CustomerBalance, error) {
	var balance pb.CustomerBalance

	err := row.Scan(
		&balance.CustomerId,
		&balance.Balance,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.WrapServiceError(errors.ErrCodeNotFound, "balance not found", err)
		} else {
			return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to scan object from database", err)
		}
	}
	
	return &balance, nil
}