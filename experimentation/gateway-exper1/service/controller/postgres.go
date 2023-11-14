package controller

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hexolan/development/experimentation/gateway-exper1/service"
	"github.com/hexolan/development/experimentation/gateway-exper1/utils/errors"
	pb "github.com/hexolan/development/experimentation/gateway-exper1/protogen/testingv1"
)

// ---

const (
	itemBaseQuery string = "SELECT id, status, title, created_at FROM items"
)

// ---

type postgresController struct {
	pCl *pgxpool.Pool

	evtC service.EventController
}

func NewPostgresController(pCl *pgxpool.Pool, evtC service.EventController) service.StorageController {
	return postgresController{pCl: pCl, evtC: evtC}
}

// ---

func scanRowToItem(row pgx.Row) (*pb.Item, error) {
	var item pb.Item

	// todo: test enum conversion
	err := row.Scan(
		item.Id, 
		item.Status, 
		item.Title,
		item.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	
	return &item, nil
}

// ---

func (c postgresController) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	row := c.pCl.QueryRow(ctx, (itemBaseQuery + " WHERE id=$1"), req.GetItemId())
	item, err := scanRowToItem(row)
	if err != nil {
		return nil, errors.NewServiceError(errors.ErrCodeService, "failed to unmarshal database row")
	}

	return &pb.GetItemResponse{Data: item}, nil
}

func (c postgresController) GetItems(ctx context.Context, req *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	// todo:
	
	return &pb.GetItemsResponse{}, nil
}

func (c postgresController) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	// todo:
	
	c.evtC.DispatchUpdatedEvent(item)
	return &pb.UpdateItemResponse{Data: item}, nil
}

func (c postgresController) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	// todo:
	
	c.evtC.DispatchDeletedEvent(item)
	return &pb.DeleteItemResponse{}, nil
}

func (c postgresController) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error)  {
	// todo:
	
	c.evtC.DispatchCreatedEvent(item)
	return &pb.CreateItemResponse{Data: item}, nil
}