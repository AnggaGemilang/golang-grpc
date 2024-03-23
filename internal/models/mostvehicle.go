package models

import grpc "len/go-grpc-api/proto"

type MostVehicle struct {
	VehicleType string `gorm:"column:vehicle_type;not null;" json:"vehicle_type" form:"vehicle_type"`
	Count       uint64 `gorm:"column:count;not null;" json:"count" form:"count"`
}

func (u MostVehicle) ToGRPC() *grpc.MostVehicle {
	return &grpc.MostVehicle{
		VehicleType: u.VehicleType,
		Count:       u.Count,
	}
}
