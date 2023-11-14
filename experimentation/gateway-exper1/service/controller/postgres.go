package controller

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hexolan/development/experimentation/gateway-exper1/service"
	pb "github.com/hexolan/development/experimentation/gateway-exper1/protogen/testingv1"
)

type postgresController struct {
	pCl *pgxpool.Pool

	evtC service.EventController
}

func NewPostgresController(pCl *pgxpool.Pool, evtC service.EventController) service.StorageController {
	return postgresController{pCl: pCl, evtC: evtC}
}

// ---

func (c postgresController) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {

}

func (c postgresController) GetItems(ctx context.Context, req *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {

}

func (c postgresController) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {

}

func (c postgresController) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {

}

func (c postgresController) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error)  {

}