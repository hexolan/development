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
	orderBaseQuery string = "SELECT id, status, customer_id, transaction_id, created_at FROM orders"
	orderItemBaseQuery string = "SELECT product_id, quantity FROM order_items"
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
	var order pb.Order

	// todo: test enum conversion
	err := row.Scan(
		order.Id, 
		order.Status, 
		order.CustomerId,
		order.TransactionId, 
		order.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	
	return &order, nil
}

func scanRowToOrderItem(row pgx.Row) (*pb.OrderItem, error) {
	var orderItem pb.OrderItem
	err := row.Scan(orderItem.ProductId, orderItem.Quantity)
	if err != nil {
		return nil, err
	}

	return &orderItem, nil
}

// Get all items included in an order (by order id).
func (c postgresController) getOrderItemsById(ctx context.Context, id string) ([]*pb.OrderItem, error) {
	rows, err := c.pCl.Query(ctx, (orderItemBaseQuery + " WHERE order_id=$1"), id)
	if err != nil {
		// todo: wrap in service error
		return nil, err
	}

	items := []*pb.OrderItem{}
	for rows.Next() {
		item, err := scanRowToOrderItem(rows)
		if err != nil {
			// todo: handling
			// return service error - something went wrong?
			continue
		}
		
		items = append(items, item)
	}

	if rows.Err() != nil {
		// todo: wrap with service error
		return nil, rows.Err()
	}

	return items, nil
}

func (c postgresController) getOrderWithItems(ctx context.Context, order *pb.Order) (*pb.Order, error) {
	// Load the order items
	items, err := c.getOrderItemsById(ctx, order.Id)
	if err != nil {
		return nil, err
	}

	// Add the order items to the order protobuf
	order.Items = items

	// Return the order
	return order, nil
}

// Get an order by its specified id
func (c postgresController) GetOrderById(ctx context.Context, id string) (*pb.Order, error) {
	// Load the order data
	row := c.pCl.QueryRow(ctx, (orderBaseQuery + " WHERE id=$1"), id)
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