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

// Postgres Storage Controller
type postgresController struct {
	pCl *pgxpool.Pool
}

func NewPostgresController(pCl *pgxpool.Pool) order.StorageController {
	return postgresController{pCl: pCl}
}

// Get an order by its specified id
func (c postgresController) GetOrder(ctx context.Context, orderId string) (*pb.Order, error) {
	return c.getOrder(ctx, nil, orderId)
}

// Internal controller method
//
// Get an order by its id.
// Optionally using a specified transaction.
func (c postgresController) getOrder(ctx context.Context, tx *pgx.Tx, orderId string) (*pb.Order, error) {
	// Load the order data
	var row pgx.Row
	if tx == nil {
		row = c.pCl.QueryRow(ctx, pgOrderBaseQuery + " WHERE id=$1", orderId)
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
// This method presumes the necessary validation has already
// taken place on the order object.
func (c postgresController) CreateOrder(ctx context.Context, orderObj *pb.Order) (*pb.Order, error) {
	// Begin a DB transaction
	tx, err := c.pCl.Begin(ctx)
	if err != nil {
		// todo: implement check to determine if database problem (likely is a connection problem)
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	// Insert order query
	var id string
	err = tx.QueryRow(
		ctx,
		"INSERT INTO orders (status, customer_id) VALUES ($1, $2) RETURNING id",
		orderObj.Status,
		orderObj.CustomerId,
	).Scan(&id)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to create order", err)
	}

	// Create records for any order items
	err = c.createOrderItems(ctx, tx, id, orderObj.Items)
	if err != nil {
		// The deffered rollback will be called
		// so the transaction will not be commited as a result of this error
		return nil, err
	}

	// Attach attrs to the order object
	orderObj.Id = id
	orderObj.CreatedAt = time.Now().Unix()

	// Prepare a created event.
	//
	// Then add the event to the outbox table with the transaction
	// to ensure that the event will be dispatched if
	// the transaction succeeds.
	evt, evtTopic, err := order.PrepareCreatedEvent(orderObj)
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

	// Return the created order
	return orderObj, nil
}



// todo: IMPLEMENT
func (c postgresController) UpdateOrder(ctx context.Context, orderId string, orderObj *pb.Order, mask *fieldmaskpb.FieldMask) (*pb.Order, error) {
	// Begin a DB transaction
	tx, err := c.pCl.Begin(ctx)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	// Set the updatable attrs (TEMP ALL BELOW NEEDS CLEANING UP)
	statementSetData := goqu.Record{"updated_at": goqu.L("timezone('utc', now())")}

	// https://github.com/mennanov/fieldmask-utils
	// https://github.com/mennanov/fmutils

	/*
	// Filter to only field mask elements
	maskPaths := mask.Paths
	// todo: ensuring that order.Items is not in maskPath
	// all order attrs (excluding Id, Items, CreatedAt and UpdatedAt)
	//
	// handle settings items in seperate DB route
	// (maybe even in a seperate /v1/order/<order_id>/items route)
	//
	// disallow setting id / created_at / updated_at entirely

	fmutils.Filter(order, maskPaths)
	log.Info().Any("filteredOrd", order).Any("paths", maskPaths).Msg("")
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

	// Get the full updated order object.
	//
	// Append the order items
	row := tx.QueryRow(ctx, pgOrderBaseQuery + " WHERE id=$1", orderId)
	orderObj, err = scanRowToOrder(row)
	if err != nil {
		return nil, err
	}

	orderObj, err = c.appendItemsToOrderObj(ctx, &tx, orderObj)
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
func (c postgresController) DeleteOrder(ctx context.Context, orderId string) error {
	return c.deleteOrder(ctx, nil, orderId)
}

// todo: IMPLEMENT
func (c postgresController) deleteOrder(ctx context.Context, tx *pgx.Tx, orderId string) error {
	return nil
}

// Get items included in an order (by order id)
func (c postgresController) GetOrderItems(ctx context.Context, orderId string) (*map[string]int32, error) {
	return c.getOrderItems(ctx, nil, orderId)
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
		rows, err = c.pCl.Query(ctx, pgOrderItemsBaseQuery + " WHERE order_id=$1", orderId)
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
func (c postgresController) SetOrderItems(ctx context.Context, orderId string, itemQuantities map[string]int32) (*map[string]int32, error) {
	// Begin a DB transaction
	tx, err := c.pCl.Begin(ctx)
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

func (c postgresController) DeleteOrderItem(ctx context.Context, orderId string, itemId string) (*map[string]int32, error) {
	// Begin a DB transaction
	tx, err := c.pCl.Begin(ctx)
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

	return &orderObj.Items, nil
}

// Get all orders related to a specified customer.
//
// TODO: implement pagination - currently limited to showing first 10
func (c postgresController) GetOrdersByCustomerId(ctx context.Context, custId string) ([]*pb.Order, error) {
	rows, err := c.pCl.Query(ctx, pgOrderBaseQuery + " WHERE customer_id=$1 LIMIT 10", custId)
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