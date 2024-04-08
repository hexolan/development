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

	"github.com/hexolan/stocklet/internal/svc/warehouse"
	"github.com/hexolan/stocklet/internal/pkg/errors"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/warehouse/v1"
)

const (
	pgProductStockBaseQuery string = "SELECT product_id, quantity FROM product_stock"
	pgReservationBaseQuery string = "SELECT id, order_id, created_at FROM reservations"
	pgReservationItemsBaseQuery string = "SELECT product_id, quantity FROM reservation_items"
)

type postgresController struct {
	cl *pgxpool.Pool
}

func NewPostgresController(cl *pgxpool.Pool) warehouse.StorageController {
	return postgresController{cl: cl}
}

//
// todo: implement
//

func (c postgresController) getReservationItems(ctx context.Context, tx *pgx.Tx, reservationId string) ([]*pb.ReservationStock, error) {
	// Determine if transaction is being used.
	var rows pgx.Rows
	var err error
	if tx == nil {
		rows, err = c.cl.Query(ctx, pgReservationItemsBaseQuery + " WHERE reservation_id=$1", reservationId)
	} else {
		rows, err = (*tx).Query(ctx, pgReservationItemsBaseQuery + " WHERE reservation_id=$1", reservationId)
	}
	
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "query error whilst fetching reserved items", err)
	}

	items := []*pb.ReservationStock{}
	for rows.Next() {
		var reservStock pb.ReservationStock

		err := rows.Scan(
			&reservStock.ProductId,
			&reservStock.Quantity,
		)
		if err != nil {
			return nil, errors.WrapServiceError(errors.ErrCodeService, "failed to scan a reservation item", err)
		}
		
		items = append(items, &reservStock)
	}

	if rows.Err() != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "error scanning item rows", rows.Err())
	}

	return items, nil
}


// Append items to the reservation
func (c postgresController) appendItemsToReservation(ctx context.Context, tx *pgx.Tx, reservation *pb.Reservation) (*pb.Reservation, error) {
	reservedItems, err := c.getReservationItems(ctx, tx, reservation.Id)
	if err != nil {
		return nil, err
	}

	// Add the items to the reservation protobuf
	reservation.ReservedStock = reservedItems

	return reservation, nil
}

// Scan a postgres row to a protobuf object
func scanRowToProductStock(row pgx.Row) (*pb.ProductStock, error) {
	var stock pb.ProductStock

	err := row.Scan(
		&stock.ProductId,
		&stock.Quantity,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.WrapServiceError(errors.ErrCodeNotFound, "stock not found", err)
		} else {
			return nil, errors.WrapServiceError(errors.ErrCodeExtService, "something went wrong scanning object", err)
		}
	}

	return &stock, nil
}

// Scan a postgres row to a protobuf object
func scanRowToReservation(row pgx.Row) (*pb.Reservation, error) {
	var reservation pb.Reservation

	// Temporary variables that require conversion
	var tmpCreatedAt pgtype.Timestamp

	err := row.Scan(
		&reservation.Id,
		&reservation.OrderId,
		&tmpCreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.WrapServiceError(errors.ErrCodeNotFound, "reservation not found", err)
		} else {
			return nil, errors.WrapServiceError(errors.ErrCodeExtService, "something went wrong scanning object", err)
		}
	}

	// convert postgres timestamps to unix format
	if tmpCreatedAt.Valid {
		reservation.CreatedAt = tmpCreatedAt.Time.Unix()
	}
	
	return &reservation, nil
}