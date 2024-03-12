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
	pb.UnimplementedOrderStorageServiceServer

	cl *pgxpool.Pool
}

func NewPostgresController(cl *pgxpool.Pool) order.StorageController {
	return postgresController{cl: cl}
}

// Get an order by its specified id from the database.
//
// Internal method. Assumes inputs are valid.
func (c postgresController) GetOrder(ctx context.Context, req *pb.DbGetOrderRequest) (*pb.DbGetOrderResponse, error) {
	order, err := c.getOrder(ctx, nil, req.OrderId)
	if err != nil {
		return nil, err
	}

	return &pb.DbGetOrderResponse{Order: order}, nil 
}

func (c postgresController) getOrder(ctx context.Context, tx *pgx.Tx, orderId string) (*pb.Order, error) {
	// Load the order data
	var row pgx.Row
	if tx == nil {
		row = c.cl.QueryRow(ctx, pgOrderBaseQuery + " WHERE id=$1", orderId)
	} else {
		row = (*tx).QueryRow(ctx, pgOrderBaseQuery + " WHERE id=$1", orderId)	
	}
	orderObj, err := scanRowToOrder(row)
	if err != nil {
		return nil, err
	}

	// Add the order items and return
	orderObj, err = c.appendItemsToOrderObj(ctx, tx, orderObj)
	if err != nil {
		return nil, err
	}

	return orderObj, nil
}

// Create a new order in the Postgres database.
//
// Internal method. Validation is already assumed.
func (c postgresController) CreateOrder(ctx context.Context, req *pb.DbCreateOrderRequest) (*pb.DbCreateOrderResponse, error) {
	// Begin a DB transaction
	tx, err := c.cl.Begin(ctx)
	if err != nil {
		// todo: implement check to determine if database problem (likely is a connection problem)
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	// Insert order query
	newOrder := pb.Order{
		Items: req.Order.Items,
		Status: req.Order.Status,
		CustomerId: req.Order.CustomerId,
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

	_, err = tx.Exec(
		ctx,
		"INSERT INTO event_outbox (aggregateid, aggregatetype, payload) VALUES ($1, $2, $3)",
		newOrder.Id,
		evtTopic,
		evt,
	)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to insert order event", err)
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to commit transaction", err)
	}

	return &pb.DbCreateOrderResponse{Order: &newOrder}, nil
}

// todo: FINISH IMPLEMENTING
//
// Internal method. Inputs are assumed valid.
func (c postgresController) UpdateOrder(ctx context.Context, req *pb.DbUpdateOrderRequest) (*pb.DbUpdateOrderResponse, error) {
	order, err := c.updateOrder(ctx, req.OrderId, req.Order, req.Mask)
	if err != nil {
		return nil, err
	}

	return &pb.DbUpdateOrderResponse{Order: order}, nil
}

func (c postgresController) updateOrder(ctx context.Context, orderId string, orderObj *pb.Order, mask *fieldmaskpb.FieldMask) (*pb.Order, error) {
	// Begin a DB transaction
	tx, err := c.cl.Begin(ctx)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	// Set the updatable attrs (TEMP ALL BELOW NEEDS CLEANING UP)
	statementSetData := goqu.Record{"updated_at": goqu.L("timezone('utc', now())")}

	// https://github.com/mennanov/fieldmask-utils
	// https://github.com/mennanov/fmutils

	/*

	maskPaths := mask.Paths
	// todo: ensuring that order.Items is not in maskPath
	// all order attrs (excluding Id, Items, CreatedAt and UpdatedAt)
	//
	// handle settings items in seperate DB route
	// (maybe even in a seperate /v1/order/<order_id>/items route)
	//
	// disallow setting id / created_at / updated_at entirely

	*/

	// put filtered mask order into patchData
	marshalledMask, err := json.Marshal(orderObj)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "error whilst assembling patch data", err)
	}
	err = json.Unmarshal(marshalledMask, &statementSetData)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "error whilst assembling patch data", err)
	}

	// Build an update statement
	statement, args, err := goqu.Dialect("postgres").Update("orders").Prepared(true).Set(statementSetData).Where(goqu.C("id").Eq(orderId)).ToSQL()
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

	_, err = tx.Exec(
		ctx,
		"INSERT INTO event_outbox (aggregateid, aggregatetype, payload) VALUES ($1, $2, $3)",
		orderObj.Id,
		evtTopic,
		evt,
	)
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

// Delete an order by its specified id.
// TODO: implement
func (c postgresController) DeleteOrder(ctx context.Context, req *pb.DbDeleteOrderRequest) (*pb.DbDeleteOrderResponse, error) {
	err := c.deleteOrder(ctx, nil, req.OrderId)
	if err != nil {
		return nil, err
	}

	return &pb.DbDeleteOrderResponse{}, nil
}

// todo: IMPLEMENT
func (c postgresController) deleteOrder(ctx context.Context, tx *pgx.Tx, orderId string) error {
	return nil
}

// Get items included in an order (by order id)
func (c postgresController) GetOrderItems(ctx context.Context, req *pb.DbGetOrderItemsRequest) (*pb.DbGetOrderItemsResponse, error) {
	orderItems, err := c.getOrderItems(ctx, nil, req.OrderId)
	if err != nil {
		return nil, err
	}

	return &pb.DbGetOrderItemsResponse{Items: *orderItems}, nil
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

	orderItems := make(map[string]int32)
	for rows.Next() {
		var (
			productId string
			productQuantity int32
		)
		err := rows.Scan(
			&productId,
			&productQuantity,
		)
		if err != nil {
			return nil, errors.WrapServiceError(errors.ErrCodeService, "failed to scan an order item", err)
		}
		
		orderItems[productId] = productQuantity
	}

	if rows.Err() != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "error whilst scanning order item rows", rows.Err())
	}

	return &orderItems, nil
}

// Sets the items attached to an order.
//
// Currently overrides any existing order items.
// TODO: operate as a patch function with any existing items
//
// TODO: ENSURING THAT THE ORDER ACTUALLY EXISTS
func (c postgresController) SetOrderItems(ctx context.Context, req *pb.DbSetOrderItemsRequest) (*pb.DbSetOrderItemsResponse, error) {
	orderItems, err := c.setOrderItems(ctx, req.OrderId, req.Items)
	if err != nil {
		return nil, err
	}

	return &pb.DbSetOrderItemsResponse{Items: *orderItems}, nil
}

func (c postgresController) setOrderItems(ctx context.Context, orderId string, itemQuantities map[string]int32) (*map[string]int32, error) {
	// Begin a DB transaction
	tx, err := c.cl.Begin(ctx)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	// Delete any existing order items
	_, err = tx.Exec(ctx, "DELETE FROM order_items WHERE order_id=$1", orderId)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to override existing items", err)
	}

	// Create the order items
	err = c.createOrderItems(ctx, tx, orderId, itemQuantities)
	if err != nil {
		return nil, err
	}

	// Prepare an updated event and add to outbox table.
	evt, evtTopic, err := order.PrepareUpdatedEvent(&pb.Order{Items: itemQuantities})
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

	return &itemQuantities, nil
}

// Set an order item's quantity.
// or add an order item if it does not already exist.
// todo: implement method
func (c postgresController) SetOrderItem(ctx context.Context, orderId string, itemId string, quantity int32) (*map[string]int32, error) {
	return nil, nil
}

///
func (c postgresController) DeleteOrderItem(ctx context.Context, req *pb.DbDeleteOrderItemRequest) (*pb.DbDeleteOrderItemResponse, error) {
	orderItems, err := c.deleteOrderItem(ctx, req.OrderId, req.ItemId)
	if err != nil {
		return nil, err
	}

	return &pb.DbDeleteOrderItemResponse{Items: *orderItems}, nil
}

func (c postgresController) deleteOrderItem(ctx context.Context, orderId string, itemId string) (*map[string]int32, error) {
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

	return &orderObj.Items, nil
}

// Get all orders related to a specified customer.
// TODO: implement pagination - currently limited to showing first 10
func (c postgresController) GetCustomerOrders(ctx context.Context, req *pb.DbGetCustomerOrdersRequest) (*pb.DbGetCustomerOrdersResponse, error) {
	orders, err := c.getCustomerOrders(ctx, req.CustomerId)
	if err != nil {
		return nil, err
	}

	return &pb.DbGetCustomerOrdersResponse{Orders: orders}, nil
}

//
func (c postgresController) getCustomerOrders(ctx context.Context, custId string) ([]*pb.Order, error) {
	rows, err := c.cl.Query(ctx, pgOrderBaseQuery + " WHERE customer_id=$1 LIMIT 10", custId)
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

// todo: proper documentation for all funcs
func (c postgresController) createOrderItems(ctx context.Context, tx pgx.Tx, orderId string, itemQuantities map[string]int32) error {
	// check if there are any items to add
	if len(itemQuantities) > 1 {
		statementVals := [][]interface{}{}
		for product_id, quantity := range itemQuantities {
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

// todo: IMPLEMENT FOLLOWING METHODS
// addOrderItem( orderId, productId, quantity )
// >	adds a new order item
//
// setOrderItem( orderId, productId, quantity )
// >	sets an order item
//
// setOrderItems( orderId, orderItems, deletePrev )   (CHANGE FROM addOrderItems - when adding deletePrev false)
// >	replace/add/update the items in an order
//
// removeOrderItem ( orderId, productId )
// >	remove an order item

// Scan a postgres row to a protobuf order object.
func scanRowToOrder(row pgx.Row) (*pb.Order, error) {
	var orderObj pb.Order

	// Temporary variables that require conversion
	var tmpCreatedAt pgtype.Timestamp
	var tmpUpdatedAt pgtype.Timestamp

	err := row.Scan(
		&orderObj.Id, 
		&orderObj.Status, 
		&orderObj.CustomerId,
		&orderObj.TransactionId, 
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
		orderObj.CreatedAt = tmpCreatedAt.Time.Unix()
	} else {
		return nil, errors.NewServiceError(errors.ErrCodeUnknown, "failed to convert order (created_at) timestamp")
	}

	if tmpUpdatedAt.Valid {
		unixUpdated := tmpUpdatedAt.Time.Unix()
		orderObj.UpdatedAt = &unixUpdated
	}
	
	return &orderObj, nil
}