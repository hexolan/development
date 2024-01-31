package controller

import (
	"time"
	"context"

	"github.com/rs/zerolog/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"

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
	// insert order query
	var id string
	err := c.pCl.QueryRow(
		ctx,
		"INSERT INTO orders (status, customer_id) VALUES ($1, $2) RETURNING id",
		order.Status,
		order.CustomerId,
	).Scan(&id)
	if err != nil {
		// todo: checking of error cause
		// connection error?
		// integrity constraint violation?
		return nil, errors.WrapServiceError(errors.ErrCodeService, "failed to create order", err)
	}

	// create records for any order items
	err = c.createOrderItems(ctx, id, order.Items)
	if err != nil {
		// failed to add order items??
		// undo above transaction and return error - or just error?
		// todo: think about how to handle this
		return nil, errors.WrapServiceError(errors.ErrCodeService, "failed to create order item records", err)
	}

	// todo: check if there are items included in the order
	// if there are - rows will have to be inserted for those as well
	// TODO: ^^^ this is important and needs doing

	// Return the created order
	order.Id = id
	order.CreatedAt = time.Now().Unix()
	return order, nil
}

// todo: IMPLEMENT
// todo: proper documentation for all funcs
func (c postgresController) createOrderItems(ctx context.Context, orderId string, itemQuantities map[string]int32) error {
	// check if there are any items to add
	if len(itemQuantities) > 1 {
		statementVals := [][]interface{}{}  // what the fuck?
		// won't let me do []goqu.Vals{} but this is okay????
		for product_id, quantity := range itemQuantities {
			statementVals = append(
				statementVals, 
				goqu.Vals{orderId, product_id, quantity},
			)
		}

		statement, args, _ := goqu.Dialect("postgres").From(
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
		_, err := c.pCl.Exec(ctx, statement, args...)
		if err != nil {
			return err
		}

		// ~~todo: assemble into single query~~ DONE ABOVE
		// HOWEVER NEED TO LOOK AT POSSIBLE IMPROVEMENTS
		// then execute. instead of making multiple insert queries
		// this is just a prototype, but need to improve in future
		// maybe implement goka or whatever query builder I used in panels?
		//
		// while on this train of thought.
		// need to experiment with the update mask in UpdateOrder
		// think of a way to implement that. may require goka anyway
		// just need to ensure validation works properly

		// create records for the order items
		/*
		for key, val := range itemQuantities {
			_, err := c.pCl.Exec(
				ctx,
				"INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)",
				orderId,
				key,
				val,
			)
			if err != nil {
				return err
			}
		}
		*/
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
func (c postgresController) UpdateOrder(ctx context.Context, order *pb.Order) error {
	// todo: actual SQL statement

	// assumption here is that Order is properly validated for patching
	// if it isn't broken - don't fix - get working for now then improve later

	// TODO: all order attrs (excluding Items, CreatedAt and UpdateAt)
	// 	-> patchData below

	patchData := goqu.Record{
		"updated_at": goqu.L("timezone('utc', now())"),
	}
	statement, args, _ := goqu.Dialect("postgres").Update("orders").Prepared(true).Set(patchData).Where(goqu.C("id").Eq(order.Id)).ToSQL()

	_, err := c.pCl.Exec(ctx, statement, args...)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeExtService, "failed to update order", err)
	}
	
	return nil
}

// todo: IMPLEMENT
func (c postgresController) DeleteOrderById(ctx context.Context, id string) error {
	return nil
}