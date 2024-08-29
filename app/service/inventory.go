package service

import (
	"context"
	"fmt"
	"grpc_inventory/service/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type inventorycService struct {
	pb.UnimplementedInventoryServiceServer
}

func NewInventoryServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	invService := inventorycService{}

	pb.RegisterInventoryServiceServer(grpcServer, &invService)
	return grpcServer
}

func (svc *inventorycService) FindInventoryByName(ctx context.Context, req *pb.FindInventoryByNameRequest) (*pb.FindInventoryResponse, error) {
	if req.Name == "bogus" {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("No items found with name %s", req.Name))
	}

	inv1 := pb.ItemInventory{
		Item: &pb.Item{
			Id:   "001",
			Sku:  "sku-1",
			Name: req.Name,
		},
		Quantity: 10,
		Location: "warehouse-1",
	}

	inv2 := pb.ItemInventory{
		Item: &pb.Item{
			Id:   "002",
			Sku:  "sku-2",
			Name: req.Name,
		},
		Quantity: 15,
		Location: "warehouse-2",
	}

	resp := pb.FindInventoryResponse{
		InventoryItems: []*pb.ItemInventory{&inv1, &inv2},
	}

	return &resp, nil
}
func (svc *inventorycService) TransferInventory(ctx context.Context, req *pb.TransferInventoryRequest) (*pb.TransferInventoryResponse, error) {
	if req.ItemId == "invalid" {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("ItemId %s does not exist", req.ItemId))
	}
	return &pb.TransferInventoryResponse{FromLocationQuantity: 20, ToLocationQuantity: 15}, nil
}
