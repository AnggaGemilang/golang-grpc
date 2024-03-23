package server

import (
	"context"
	"len/go-grpc-api/internal/services"
	grpc "len/go-grpc-api/proto"
)

type vehiclesCRUDServer struct {
	grpc.UnimplementedVehiclesCRUDServer
}

func NewVehiclesCRUDServer() grpc.VehiclesCRUDServer {
	return &vehiclesCRUDServer{}
}

func (u vehiclesCRUDServer) CreateVehicle(ctx context.Context, input *grpc.VehicleInput) (*grpc.VehicleResponse, error) {
	s := services.GetVehicleService()
	return s.CreateVehicle(input)
}

func (u vehiclesCRUDServer) ListVehicles(ctx context.Context, req *grpc.ListVehicleReq) (*grpc.ListVehicleRes, error) {
	s := services.GetVehicleService()
	return s.ListVehicles()
}

func (u vehiclesCRUDServer) MostVehicles(ctx context.Context, filter *grpc.MostVehicleReq) (*grpc.ListMostVehicleRes, error) {
	s := services.GetVehicleService()
	return s.MostVehicles(filter)
}

func (u vehiclesCRUDServer) DeleteVehicle(ctx context.Context, id *grpc.ID) (*grpc.VehicleResponse, error) {
	s := services.GetVehicleService()
	return s.DeleteVehicle(id)
}
