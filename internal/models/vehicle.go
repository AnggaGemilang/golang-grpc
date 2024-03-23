package models

import grpc "len/go-grpc-api/proto"

type Vehicle struct {
	ID           uint64 `gorm:"column:id;primary_key;auto_increment;" json:"id"`
	VehicleType  string `gorm:"column:vehicle_type;not null;" json:"vehicle_type" form:"vehicle_type"`
	VehicleModel string `gorm:"column:vehicle_model;not null;" json:"vehicle_model" form:"vehicle_model"`
	Longitude    float64 `gorm:"column:longitude;not null;" json:"longitude" form:"longitude"`
	Latitude     float64 `gorm:"column:latitude;not null;" json:"latitude" form:"latitude"`
	CreatedAt    string `gorm:"column:created_at;not null;" json:"created_at" form:"created_at"`
}

func (u Vehicle) ToGRPC() *grpc.Vehicle {
	return &grpc.Vehicle{
		Id:           u.ID,
		VehicleType:  u.VehicleType,
		VehicleModel: u.VehicleModel,
		Longitude:    u.Longitude,
		Latitude:     u.Latitude,
		CreatedAt:    u.CreatedAt,
	}
}
