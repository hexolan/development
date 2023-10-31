package order

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hexolan/stocklet/internal/pkg/errors"
	pb "github.com/hexolan/stocklet/internal/pkg/protogen/order/v1"
)

type DatabaseRepo struct {
	db *pgxpool.Pool
}

func NewDatabaseRepo(db *pgxpool.Pool) DatabaseRepo {
	return DatabaseRepo{db: db}
}

func scanRowToOrder(row pgx.Row) *pb.Order {
	var order pb.Order
	err := row.Scan(order.Id, order.Status, order.CustomerId, order.CreatedAt)
	if err != nil {
		return nil
	}

	return &order
}

func scanRowToOrderItem(row pgx.Row) *pb.OrderItem {
	var orderItem pb.OrderItem
	err := row.Scan(orderItem.ProductId, orderItem.Quantity)
	if err != nil {
		return nil
	}

	return &orderItem
}

func scanRowsToOrderItems(rows []pgx.Row) []*pb.OrderItem {
	items := []*pb.OrderItem{}
	for _, row := range rows {
		orderItem := scanRowToOrderItem(row)
		if orderItem == nil {
			continue
		}
		
		items = append(items, orderItem)
	}

	return items
}

func (r DatabaseRepo) GetOrder(ctx context.Context, id string) (*pb.Order, error) {
	row := r.db.QueryRow(ctx, "SELECT id, status, customer_id, created_at FROM panels WHERE id=$1", id)
	order := scanRowToOrder(row)
	if order == nil {
		return nil, errors.NewServiceError(errors.ErrCodeService, "failed to unmarshal database row")
	}

	// todo: fetching order items
	// append to loaded order record
}