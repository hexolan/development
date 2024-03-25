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
	"time"
	"context"
	"slices"
	"encoding/json"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/hexolan/stocklet/internal/svc/order"
	"github.com/hexolan/stocklet/internal/pkg/errors"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

const (
	pgOrderBaseQuery string = "SELECT id, status, customer_id, transaction_id, created_at, updated_at FROM orders"
	pgOrderItemsBaseQuery string = "SELECT product_id, quantity FROM order_items"
)

type postgresController struct {
	cl *pgxpool.Pool
}

func NewPostgresController(cl *pgxpool.Pool) order.StorageController {
	return postgresController{cl: cl}
}

// Get an order by its specified id from the database.
//
// Internal method. Assumes inputs are valid.
func (c postgresController) GetOrder(ctx context.Context, id string) (*pb.Order, error) {
	return c.getOrder(ctx, nil, id)
}

func (c postgresController) getOrder(ctx context.Context, tx *pgx.Tx, id string) (*pb.Order, error) {
	// Query order
	var row pgx.Row
	if tx == nil {
		row = c.cl.QueryRow(ctx, pgOrderBaseQuery + " WHERE id=$1", id)
	} else {
		row = (*tx).QueryRow(ctx, pgOrderBaseQuery + " WHERE id=$1", id)	
	}

	// Scan row to protobuf order
	order, err := scanRowToOrder(row)
	if err != nil {
		return nil, err
	}

	// Append the order items
	order, err = c.appendItemsToOrderObj(ctx, tx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// Create a new order in the database.
//
// Internal method. Validation is already assumed.
func (c postgresController) CreateOrder(ctx context.Context, orderObj *pb.Order) (*pb.Order, error) {
	// Begin a DB transaction
	tx, err := c.cl.Begin(ctx)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	// Prepare and perform insert query
	newOrder := pb.Order{
		Items: orderObj.Items,
		Status: orderObj.Status,
		CustomerId: orderObj.CustomerId,
		CreatedAt: time.Now().Unix(),
	}
	err = tx.QueryRow(
		ctx,
		"INSERT INTO orders (status, customer_id) VALUES ($1, $2) RETURNING id",
		newOrder.Status,
		newOrder.CustomerId,
	).Scan(&newOrder.Id)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to create order", err)
	}

	// Create records for any order items
	err = c.createOrderItems(ctx, tx, newOrder.Id, newOrder.Items)
	if err != nil {
		// The deffered rollback will be called (so the transaction will not be commited)
		return nil, err
	}

	// Prepare a created event.
	//
	// Then add the event to the outbox table with the transaction
	// to ensure that the event will be dispatched if
	// the transaction succeeds.
	evt, evtTopic, err := order.PrepareCreatedEvent(&newOrder)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "failed to create order event", err)
	}

	_, err = tx.Exec(ctx, "INSERT INTO event_outbox (aggregateid, aggregatetype, payload) VALUES ($1, $2, $3)", newOrder.Id, evtTopic, evt)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to insert order event", err)
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to commit transaction", err)
	}

	return &newOrder, nil
}

// Update an order in the database.
//
// Internal method. Inputs are assumed valid.
func (c postgresController) UpdateOrder(ctx context.Context, id string, order *pb.Order, mask *fieldmaskpb.FieldMask) (*pb.Order, error) {
	return c.updateOrder(ctx, id, order, mask)
}

func (c postgresController) updateOrder(ctx context.Context, orderId string, orderObj *pb.Order, mask *fieldmaskpb.FieldMask) (*pb.Order, error) {
	// Begin a DB transaction
	tx, err := c.cl.Begin(ctx)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	// Forbidden attributes
	if slices.Contains(mask.Paths, "id") {
		return nil, errors.NewServiceError(errors.ErrCodeInvalidArgument, "cannot update primary key")
	}

	// Set the updatable attrs (TEMP ALL BELOW NEEDS CLEANING UP)
	updateVals := goqu.Record{"updated_at": goqu.L("timezone('utc', now())")}

	// Check if order items are to be updated
	if slices.Contains(mask.Paths, "items") {
		err = c.setOrderItems(ctx, &tx, orderId, orderObj.Items)
		if err != nil {
			return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to update order items", err)
		}

		// remove Items from orderObj
		// to prevent (invalidly) being added to update statement below
		orderObj.Items = nil
	}

	// put filtered mask order into patchData
	marshalledMask, err := json.Marshal(orderObj)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "error whilst assembling patch data", err)
	}
	err = json.Unmarshal(marshalledMask, &updateVals)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "error whilst assembling patch data", err)
	}

	// Build an update statement
	statement, args, err := goqu.Dialect("postgres").Update("orders").Prepared(true).Set(updateVals).Where(goqu.C("id").Eq(orderId)).ToSQL()
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "failed to build SQL statement", err)
	}

	// Execute the statement as part of the transaction
	_, err = tx.Exec(ctx, statement, args...)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to update order", err)
	}

	// Get the full updated order object
	orderObj, err = c.getOrder(ctx, &tx, orderId)
	if err != nil {
		return nil, err
	}
	
	// Create an order updated event.
	evt, evtTopic, err := order.PrepareUpdatedEvent(orderObj)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "failed to create order event", err)
	}

	_, err = tx.Exec(ctx, "INSERT INTO event_outbox (aggregateid, aggregatetype, payload) VALUES ($1, $2, $3)", orderObj.Id, evtTopic, evt)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to add order event", err)
	}
	
	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to commit transaction", err)
	}

	return orderObj, nil
}

// Delete an order from the database
func (c postgresController) DeleteOrder(ctx context.Context, id string) error {
	// Begin a DB transaction
	tx, err := c.cl.Begin(ctx)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeExtService, "failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	// Delete the order
	cmdTag, err := tx.Exec(ctx, "DELETE FROM orders WHERE id=$1", id)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeExtService, "failed to execute delete query", err)
	}

	// Ensure an order was deleted
	if cmdTag.RowsAffected() < 1 {
		return errors.NewServiceError(errors.ErrCodeNotFound, "order does not exist")
	}

	// Prepare an order deleted event.
	evt, evtTopic, err := order.PrepareDeletedEvent(id)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeService, "failed to create order event", err)
	}

	_, err = tx.Exec(ctx, "INSERT INTO event_outbox (aggregateid, aggregatetype, payload) VALUES ($1, $2, $3)", id, evtTopic, evt)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeExtService, "failed to add order event", err)
	}
	
	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeExtService, "failed to commit transaction", err)
	}

	return nil
}

// Get items included in an order (by order id)
func (c postgresController) GetOrderItems(ctx context.Context, id string) (*map[string]int32, error) {
	return c.getOrderItems(ctx, nil, id)
}

// Internal controller method
//
// Get an order's items.
// Optionally using a specified transaction.
func (c postgresController) getOrderItems(ctx context.Context, tx *pgx.Tx, orderId string) (*map[string]int32, error) {
	// Determine if transaction is being used.
	var rows pgx.Rows
	var err error
	if tx == nil {
		rows, err = c.cl.Query(ctx, pgOrderItemsBaseQuery + " WHERE order_id=$1", orderId)
	} else {
		rows, err = (*tx).Query(ctx, pgOrderItemsBaseQuery + " WHERE order_id=$1", orderId)
	}
	
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "query error whilst fetching order items", err)
	}

	items := make(map[string]int32)
	for rows.Next() {
		var (
			itemId string
			quantity int32
		)
		err := rows.Scan(
			&itemId,
			&quantity,
		)
		if err != nil {
			return nil, errors.WrapServiceError(errors.ErrCodeService, "failed to scan an order item", err)
		}
		
		items[itemId] = quantity
	}

	if rows.Err() != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "error whilst scanning order item rows", rows.Err())
	}

	return &items, nil
}

// Sets the items attached to an order.
//
// Currently overrides any existing order items.
//
// TODO: operate as a patch function with any existing items
// TODO: ENSURING THAT THE ORDER ACTUALLY EXISTS
func (c postgresController) SetOrderItems(ctx context.Context, orderId string, items map[string]int32) (*map[string]int32, error) {
	// Begin a DB transaction
	tx, err := c.cl.Begin(ctx)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	// Set the order items
	c.setOrderItems(ctx, &tx, orderId, items)

	// Prepare an updated event and add to outbox table.
	evt, evtTopic, err := order.PrepareUpdatedEvent(&pb.Order{Items: items})
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "failed to create order event", err)
	}

	_, err = tx.Exec(ctx, "INSERT INTO event_outbox (aggregateid, aggregatetype, payload) VALUES ($1, $2, $3)", orderId, evtTopic, evt)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to commit event", err)
	}
	
	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to commit transaction", err)
	}

	return &items, nil
}

func (c postgresController) setOrderItems(ctx context.Context, tx *pgx.Tx, orderId string, items map[string]int32) error {
	// Delete any existing order items
	var err error
	if tx == nil {
		_, err = c.cl.Exec(ctx, "DELETE FROM order_items WHERE order_id=$1", orderId)
	} else {
		_, err = (*tx).Exec(ctx, "DELETE FROM order_items WHERE order_id=$1", orderId)
	}
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeExtService, "failed to override existing items", err)
	}

	// Create the order items
	err = c.createOrderItems(ctx, *tx, orderId, items)
	if err != nil {
		return err
	}

	return nil
}

// Set an order item's quantity.
func (c postgresController) SetOrderItem(ctx context.Context, orderId string, itemId string, quantity int32) (*map[string]int32, error) {
	items, err := c.getOrderItems(ctx, nil, orderId)
	if err != nil {
		return nil, err
	}

	(*items)[itemId] = quantity
	_, err = c.SetOrderItems(ctx, orderId, *items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// Delete an order item.
func (c postgresController) DeleteOrderItem(ctx context.Context, orderId string, itemId string) (*map[string]int32, error) {
	// Begin a DB transaction
	tx, err := c.cl.Begin(ctx)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	// Delete the order item
	_, err = tx.Exec(ctx, "DELETE FROM order_items WHERE order_id=$1 AND product_id=$2", orderId, itemId)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to override existing items", err)
	}

	// Get the updated order object.
	orderObj, err := c.getOrder(ctx, &tx, orderId)
	if err != nil {
		return nil, err
	}
	
	// Prepare an order updated event.
	evt, evtTopic, err := order.PrepareUpdatedEvent(orderObj)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "failed to create order event", err)
	}

	_, err = tx.Exec(ctx, "INSERT INTO event_outbox (aggregateid, aggregatetype, payload) VALUES ($1, $2, $3)", orderObj.Id, evtTopic, evt)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to add order event", err)
	}
	
	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to commit transaction", err)
	}

	return &orderObj.Items, nil
}

// Get all orders related to a specified customer.
//
// TODO: implement pagination - currently limited to showing first 10
func (c postgresController) GetCustomerOrders(ctx context.Context, customerId string) ([]*pb.Order, error) {
	rows, err := c.cl.Query(ctx, pgOrderBaseQuery + " WHERE customer_id=$1 LIMIT 10", customerId)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "query error whilst fetching customer orders", err)
	}

	orders := []*pb.Order{}
	for rows.Next() {
		orderObj, err := scanRowToOrder(rows)
		if err != nil {
			return nil, err
		}

		// Append the order's items
		orderObj, err = c.appendItemsToOrderObj(ctx, nil, orderObj)
		if err != nil {
			return nil, err
		}

		orders = append(orders, orderObj)
	}

	if rows.Err() != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "error whilst scanning order rows", rows.Err())
	}

	return orders, nil
}

// Internal controller method
//
// Appends order items to an order object.
func (c postgresController) appendItemsToOrderObj(ctx context.Context, tx *pgx.Tx, orderObj *pb.Order) (*pb.Order, error) {
	// Load the order items
	orderItems, err := c.getOrderItems(ctx, tx, orderObj.Id)
	if err != nil {
		return nil, err
	}

	// Add the order items to the order protobuf
	orderObj.Items = *orderItems

	// Return the order
	return orderObj, nil
}

// Build and exec an insert statement for a map of order items
func (c postgresController) createOrderItems(ctx context.Context, tx pgx.Tx, orderId string, items map[string]int32) error {
	// check there are items to add
	if len(items) > 1 {
		statementVals := [][]interface{}{}
		for product_id, quantity := range items {
			statementVals = append(
				statementVals, 
				goqu.Vals{orderId, product_id, quantity},
			)
		}

		statement, args, err := goqu.Dialect("postgres").From(
			"order_items",
		).Insert().Cols(
			"order_id",
			"product_id",
			"quantity",
		).Vals(
			statementVals...
		).Prepared(
			true,
		).ToSQL()
		if err != nil {
			return errors.WrapServiceError(errors.ErrCodeService, "failed to build SQL statement", err)
		}
		
		// Execute the statement on the transaction
		_, err = tx.Exec(ctx, statement, args...)
		if err != nil {
			return errors.WrapServiceError(errors.ErrCodeExtService, "failed to add items to order", err)
		}
	}

	return nil
}

// Scan a postgres row to a protobuf order object.
func scanRowToOrder(row pgx.Row) (*pb.Order, error) {
	var order pb.Order

	// Temporary variables that require conversion
	var tmpCreatedAt pgtype.Timestamp
	var tmpUpdatedAt pgtype.Timestamp

	err := row.Scan(
		&order.Id, 
		&order.Status, 
		&order.CustomerId,
		&order.TransactionId, 
		&tmpCreatedAt,
		&tmpUpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.WrapServiceError(errors.ErrCodeNotFound, "order not found", err)
		} else {
			return nil, errors.WrapServiceError(errors.ErrCodeExtService, "something went wrong scanning order", err)
		}
	}

	// Convert the temporary variables
	//
	// This includes converting postgres timestamps to unix format
	if tmpCreatedAt.Valid {
		order.CreatedAt = tmpCreatedAt.Time.Unix()
	} else {
		return nil, errors.NewServiceError(errors.ErrCodeUnknown, "failed to convert order (created_at) timestamp")
	}

	if tmpUpdatedAt.Valid {
		unixUpdated := tmpUpdatedAt.Time.Unix()
		order.UpdatedAt = &unixUpdated
	}
	
	return &order, nil
}