package services

import (
	"fmt"
	"len/go-grpc-api/internal/models"
	"len/go-grpc-api/internal/repository"
	grpc "len/go-grpc-api/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VehicleService struct{}

var vehicleService *VehicleService

func GetVehicleService() *VehicleService {
	if vehicleService == nil {
		vehicleService = &VehicleService{}
	}
	return vehicleService
}

func (u VehicleService) CreateVehicle(input *grpc.VehicleInput) (*grpc.VehicleResponse, error) {
	r := repository.GetVehicleRepository()
	vehicle := models.Vehicle{
		VehicleType:  input.VehicleType,
		VehicleModel: input.VehicleModel,
		Longitude:    input.Longitude,
		Latitude:     input.Latitude,
		CreatedAt:    input.CreatedAt,
	}
	if vehicle.VehicleType == "" || vehicle.ToGRPC().VehicleModel == "" || vehicle.Longitude == 0.0 || vehicle.Latitude == 0.0 || vehicle.CreatedAt == "" {
		fmt.Println("Masukkan data dengan lengkap")
		return nil, status.Errorf(codes.InvalidArgument, "Masukkan data dengan lengkap")
	} else if vehicle.VehicleType != "truck" && vehicle.VehicleType != "bus" && vehicle.VehicleType != "car" {
		fmt.Println("Jenis kendaraan tidak diketahui")
		return nil, status.Errorf(codes.InvalidArgument, "Jenis kendaraan tidak diketahui")
	} else if err := r.Add(&vehicle); err != nil {
		fmt.Println(err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	} else {
		fmt.Println("Data berhasil ditambahkan")
		return &grpc.VehicleResponse{
			Status: "OK",
			Id:     vehicle.ID,
			Desc:   "Data berhasil ditambahkan",
		}, nil
	}
}

func (u VehicleService) ListVehicles() (*grpc.ListVehicleRes, error) {
	r := repository.GetVehicleRepository()
	if vehicles, err := r.All(); err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.NotFound, "Data tidak ditemukan")
	} else {
		fmt.Println("Data berhasil ditampilkan")
		var res []*grpc.Vehicle
		for _, u := range *vehicles {
			res = append(res, u.ToGRPC())
		}
		return &grpc.ListVehicleRes{
			Status:   "OK",
			Amount:   uint64(len(res)),
			Vehicles: res,
		}, nil
	}
}

func (u VehicleService) MostVehicles(filter *grpc.MostVehicleReq) (*grpc.ListMostVehicleRes, error) {
	r := repository.GetVehicleRepository()
	if filter.Filter != "hourly" && filter.Filter != "daily" && filter.Filter != "lifetime" {
		fmt.Println("Jenis filter tidak diketahui")
		return nil, status.Errorf(codes.InvalidArgument, "Jenis filter tidak diketahui")
	} else if mostVehicles, err := r.Query(filter.Filter); err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.NotFound, "Data tidak ditemukan")
	} else {
		fmt.Println("Data berhasil ditampilkan")
		var res []*grpc.MostVehicle
		for _, u := range *mostVehicles {
			res = append(res, u.ToGRPC())
		}
		return &grpc.ListMostVehicleRes{
			Status:       "OK",
			Amount:       uint64(len(res)),
			MostVehicles: res,
		}, nil
	}
}

func (u VehicleService) DeleteVehicle(id *grpc.ID) (*grpc.VehicleResponse, error) {
	r := repository.GetVehicleRepository()
	if vehicle, err := r.Get(id.Id); err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.NotFound, "Data tidak ditemukan")
	} else {
		if err := r.Delete(vehicle); err != nil {
			fmt.Println(err.Error())
			return nil, status.Errorf(codes.Internal, err.Error())
		} else {
			fmt.Println("Data dengan id ", id.Id, " berhasil dihapus")
			return &grpc.VehicleResponse{
				Status: "OK",
				Id:     id.Id,
				Desc:   "Data berhasil dihapus",
			}, nil
		}
	}
}
