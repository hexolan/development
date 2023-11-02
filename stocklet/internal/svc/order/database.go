package order

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hexolan/stocklet/internal/pkg/errors"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

const (
	orderBaseQuery string = "SELECT id, status, customer_id, transaction_id, created_at FROM orders"
	orderItemBaseQuery string = "SELECT product_id, quantity FROM order_items"
)

type DatabaseRepo struct {
	db *pgxpool.Pool
}

func NewDatabaseRepo(db *pgxpool.Pool) DatabaseRepo {
	return DatabaseRepo{db: db}
}

func scanRowToOrder(row pgx.Row) (*pb.Order, error) {
	var order pb.Order

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
func (r DatabaseRepo) getOrderItemsById(ctx context.Context, id string) ([]*pb.OrderItem, error) {
	rows, err := r.db.Query(ctx, (orderItemBaseQuery + " WHERE order_id=$1"), id)
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

// Get an order by its specified id
func (r DatabaseRepo) GetOrderById(ctx context.Context, id string) (*pb.Order, error) {
	// Load the order data
	row := r.db.QueryRow(ctx, (orderBaseQuery + " WHERE id=$1"), id)
	order, err := scanRowToOrder(row)
	if err != nil {
		return nil, errors.NewServiceError(errors.ErrCodeService, "failed to unmarshal database row")
	}

	// Load the order items
	items, err := r.getOrderItemsById(ctx, order.Id)
	if err != nil {
		return nil, err
	}

	// Add the order items to the order protobuf
	order.Items = items

	// Return the order
	return order, nil
}