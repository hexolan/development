package service

import (
	"context"

	pb "github.com/hexolan/development/experimentation/gateway-exper1/protogen/testingv1"
)

type StorageController interface {

}

type EventController interface {

}

type Service struct {
	EvtCtrl EventController
	StrCtrl StorageController

	pb.UnimplementedTestServiceServer
}

func NewTestService(evtCtrl EventController, strCtrl StorageController) Service {
	return Service{
		EvtCtrl: evtCtrl,
		StrCtrl: strCtrl,
	}
}

// -- -- --

func (svc Service) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {

}

func (svc Service) GetItems(ctx context.Context, req *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {

}

func (svc Service) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {

}

func (svc Service) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {

}

func (svc Service) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error)  {

}