package controller

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

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

func scanRowToOrder(row pgx.Row) (*pb.Order, error) {
	var order *pb.Order

	err := row.Scan(
		&order.Id, 
		&order.Status, 
		&order.CustomerId,
		&order.TransactionId, 
		&order.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	
	return order, nil
}

// Get items included in an order (by order id)
func (c postgresController) getOrderItemsByOrderId(ctx context.Context, id string) (*map[string]int32, error) {
	rows, err := c.pCl.Query(ctx, (pgOrderItemsBaseQuery + " WHERE order_id=$1"), id)
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeService, "query error wilst fetching order items", err)
	}

	orderItems := make(map[string]int32)
	for rows.Next() {
		var (
			productId string
			productQuantity int32
		)
		err := rows.Scan(
			productId,
			productQuantity,
		)
		if err != nil {
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

func (c postgresController) getOrderWithItems(ctx context.Context, order *pb.Order) (*pb.Order, error) {
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
		return nil, errors.NewServiceError(errors.ErrCodeService, "failed to unmarshal database row")
	}

	// Add the order items and return
	order, err = c.getOrderWithItems(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (c postgresController) GetOrdersByCustomerId(ctx context.Context, custId string) ([]*pb.Order, error) {
	return nil, nil
}

func (c postgresController) UpdateOrder(ctx context.Context, order *pb.Order) error {
	// todo: actual SQL statement
	query := "UPDATE orders SET xyz WHERE abc"
	_, err := c.pCl.Exec(ctx, query)
	if err != nil {
		return errors.WrapServiceError(errors.ErrCodeExtService, "failed to update order", err)
	}

	return nil
}

func (c postgresController) DeleteOrderById(ctx context.Context, id string) error {
	return nil
}