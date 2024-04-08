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

	"github.com/hexolan/stocklet/internal/svc/shipping"
	"github.com/hexolan/stocklet/internal/pkg/errors"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/shipping/v1"
)

const (
	pgShipmentBaseQuery string = "SELECT id, order_id, dispatched, created_at FROM shipments"
	pgShipmentItemsBaseQuery string = "SELECT shipment_id, product_id, quantity FROM shipment_items"
)

type postgresController struct {
	cl *pgxpool.Pool
}

func NewPostgresController(cl *pgxpool.Pool) shipping.StorageController {
	return postgresController{cl: cl}
}

//
// todo: implement
//

func (c postgresController) getShipmentItems(ctx context.Context, tx *pgx.Tx, shipmentId string) (*[]*pb.ShipmentItem, error) {
	// Determine if transaction is being used.
	var rows pgx.Rows
	var err error
	if tx == nil {
		rows, err = c.cl.Query(ctx, pgShipmentItemsBaseQuery + " WHERE shipment_id=$1", shipmentId)
	} else {
		rows, err = (*tx).Query(ctx, pgShipmentItemsBaseQuery + " WHERE shipment_id=$1", shipmentId)
	}
	
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "query error whilst fetching items", err)
	}

	shipmentItems := []*pb.ShipmentItem{}
	for rows.Next() {
		var shipmentItem pb.ShipmentItem
		shipmentItem.ShipmentId = shipmentId
		err := rows.Scan(
			&shipmentItem.ProductId,
			&shipmentItem.Quantity,
		)
		if err != nil {
			return nil, errors.WrapServiceError(errors.ErrCodeService, "failed to scan an order item", err)
		}
		
		shipmentItems = append(shipmentItems, &shipmentItem)
	}

	if rows.Err() != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "error whilst scanning order item rows", rows.Err())
	}

	return &shipmentItems, nil
}

// Scan a postgres row to a protobuf object
func scanRowToShipment(row pgx.Row) (*pb.Shipment, error) {
	var shipment pb.Shipment

	// Temporary variables that require conversion
	var tmpCreatedAt pgtype.Timestamp

	err := row.Scan(
		&shipment.Id, 
		&shipment.OrderId, 
		&shipment.Dispatched,
		&tmpCreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.WrapServiceError(errors.ErrCodeNotFound, "shipment not found", err)
		} else {
			return nil, errors.WrapServiceError(errors.ErrCodeExtService, "something went wrong scanning object", err)
		}
	}

	// convert postgres timestamps to unix format
	if tmpCreatedAt.Valid {
		shipment.CreatedAt = tmpCreatedAt.Time.Unix()
	} else {
		return nil, errors.NewServiceError(errors.ErrCodeUnknown, "something went wrong scanning objec (timestamp conversion error)")
	}
	
	return &shipment, nil
}