package controller

import (
	"time"
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
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
	pgOrderBaseQuery string = "SELECT id, status, customer_id, transaction_id, created_at FROM orders"
	pgOrderItemsBaseQuery string = "SELECT product_id, quantity FROM order_items"
)

// Postgres Storage Controller
type postgresController struct {
	pCl *pgxpool.Pool
	evtC order.EventController
}

func NewPostgresController(pCl *pgxpool.Pool, evtC order.EventController) order.StorageController {
	return postgresController{pCl: pCl, evtC: evtC}
}

// todo: clean up func
func scanRowToOrder(row pgx.Row) (*pb.Order, error) {
	var order pb.Order

	// convert from postgres timestamp format to int64 unix timestamp
	var tmpCreatedAt pgtype.Timestamp

	// todo: implementing updated_at into the query as well

	err := row.Scan(
		&order.Id, 
		&order.Status, 
		&order.CustomerId,
		&order.TransactionId, 
		&tmpCreatedAt,
	)
	if err != nil {
		// todo: delete after
		log.Error().Err(err).Msg("debug scan")
		if err == pgx.ErrNoRows {
			return nil, errors.WrapServiceError(errors.ErrCodeNotFound, "order not found", err)
		} else {
			return nil, errors.WrapServiceError(errors.ErrCodeExtService, "something went wrong scanning order", err)
		}
	}

	// convert timestamp to unix
	if tmpCreatedAt.Valid {
		order.CreatedAt = tmpCreatedAt.Time.Unix()
	} else {
		return nil, errors.NewServiceError(errors.ErrCodeUnknown, "failed to convert order (created_at) timestamp")
	}
	
	return &order, nil
}

// Get items included in an order (by order id)
func (c postgresController) getOrderItemsByOrderId(ctx context.Context, id string) (*map[string]int32, error) {
	rows, err := c.pCl.Query(ctx, (pgOrderItemsBaseQuery + " WHERE order_id=$1"), id)
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
			log.Error().Err(err).Msg("debug scan items")
			// something went wrong when scanning an order item
			return nil, errors.WrapServiceError(errors.ErrCodeService, "failed to scan an order item", err)
		}
		
		orderItems[productId] = productQuantity
	}

	if rows.Err() != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "error whilst scanning order item rows", rows.Err())
	}

	return &orderItems, nil
}

func (c postgresController) appendOrderItems(ctx context.Context, order *pb.Order) (*pb.Order, error) {
	// Load the order items
	orderItems, err := c.getOrderItemsByOrderId(ctx, order.Id)
	if err != nil {
		return nil, err
	}

	// Add the order items to the order protobuf
	order.Items = *orderItems

	// Return the order
	return order, nil
}

// Get an order by its specified id
func (c postgresController) GetOrderById(ctx context.Context, id string) (*pb.Order, error) {
	// Load the order data
	row := c.pCl.QueryRow(
		ctx,
		pgOrderBaseQuery + " WHERE id=$1",
		id,
	)
	order, err := scanRowToOrder(row)
	if err != nil {
		return nil, err
	}

	// Add the order items and return
	order, err = c.appendOrderItems(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// Get all orders related to a specified customer.
//
// TODO: implement pagination - currently limited to showing first 10
func (c postgresController) GetOrdersByCustomerId(ctx context.Context, custId string) ([]*pb.Order, error) {
	rows, err := c.pCl.Query(ctx, (pgOrderBaseQuery + " WHERE customer_id=$1 LIMIT 10"), custId)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "query error whilst fetching customer orders", err)
	}

	orders := []*pb.Order{}
	for rows.Next() {
		order, err := scanRowToOrder(rows)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	if rows.Err() != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "error whilst scanning order rows", rows.Err())
	}

	return orders, nil
}

// todo: IMPLEMENT
// Create a new order in the Postgres database.
//
// This interface assumes that the necessary validation has already
// taken place on the order object.
func (c postgresController) CreateOrder(ctx context.Context, order *pb.Order) (*pb.Order, error) {
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
		order.Status,
		order.CustomerId,
	).Scan(&id)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to create order", err)
	}

	// Create records for any order items
	err = c.createOrderItems(ctx, tx, id, order.Items)
	if err != nil {
		// the deffered rollback will be called
		// - and the order will not be commited as a result of this err
		return nil, err
	}

	// TODO: implement event outbox table
	// and add order created event to the outbox
	// (which is commited with the transaction)

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to commit transaction", err)
	}

	// Return the created order
	order.Id = id
	order.CreatedAt = time.Now().Unix()
	return order, nil
}

// todo: IMPLEMENT
// todo: proper documentation for all funcs
func (c postgresController) createOrderItems(ctx context.Context, tx pgx.Tx, orderId string, itemQuantities map[string]int32) error {
	// check if there are any items to add
	if len(itemQuantities) > 1 {
		statementVals := [][]interface{}{}  // what the fuck?
		// won't let me do []goqu.Vals{} but this is okay???
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

// todo: IMPLEMENT
func (c postgresController) UpdateOrder(ctx context.Context, order *pb.Order, mask *fieldmaskpb.FieldMask) error {
	// Begin a DB transaction
	tx, err := c.pCl.Begin(ctx)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeExtService, "failed to begin transaction", err)
	}
	defer tx.Rollback(ctx)

	// Set the updatable attrs (TEMP ALL BELOW NEEDS CLEANING UP)
	orderId := order.Id
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
	marshalledMask, err := json.Marshal(order)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeService, "error whilst assembling patch data", err)
	}
	err = json.Unmarshal(marshalledMask, &statementSetData)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeService, "error whilst assembling patch data", err)
	}

	// build an update statement
	statement, args, err := goqu.Dialect("postgres").Update("orders").Prepared(true).Set(statementSetData).Where(goqu.C("id").Eq(orderId)).ToSQL()
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeService, "failed to build SQL statement", err)
	}

	// Execute the statement as part of the transaction
	_, err = tx.Exec(ctx, statement, args...)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeExtService, "failed to update order", err)
	}
	
	// TODO: implement event outbox table
	// and add order updated event to the outbox
	// (which is commited with the transaction)

	// https://stackoverflow.com/questions/52289338/how-can-i-catch-an-event-of-a-new-postgresql-record-with-golang

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeExtService, "failed to commit transaction", err)
	}

	return nil
}

// todo: IMPLEMENT
func (c postgresController) DeleteOrderById(ctx context.Context, id string) error {
	return nil
}