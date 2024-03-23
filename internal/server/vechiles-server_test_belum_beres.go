package server

import (
	"context"
	pb "len/go-grpc-api/proto"
	"len/go-grpc-api/util"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(2048 * 2048)

	server := grpc.NewServer()

	pb.RegisterVehiclesCRUDServer(server, &vehiclesCRUDServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestCreateVehicleData(t *testing.T) {
	tests := []struct {
		name    string
		input   *pb.VehicleInput
		res     *pb.VehicleResponse
		errCode codes.Code
		errMsg  string
	}{
		{
			"Terdapat input data yang kosong",
			&pb.VehicleInput{
				VehicleType:  util.RandomString(6),
				VehicleModel: util.RandomString(8),
				Latitude:     float64(util.RandomInt(0, 10)),
				CreatedAt:    "2024-03-10 10:00:00",
			},
			nil,
			codes.InvalidArgument,
			"Proses tambah data gagal karena terdapat input yang kosong",
		},
		{
			"Data berhasil ditambahkan",
			&pb.VehicleInput{
				VehicleType:  util.RandomString(6),
				VehicleModel: util.RandomString(8),
				Longitude:    float64(util.RandomInt(0, 10)),
				Latitude:     float64(util.RandomInt(0, 10)),
				CreatedAt:    "2024-03-10 10:00:00",
			},
			&pb.VehicleResponse{Status: "OK"},
			codes.OK,
			"",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewVehiclesCRUDClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &pb.VehicleInput{
				VehicleType:  util.RandomString(6),
				VehicleModel: util.RandomString(8),
				Longitude:    float64(util.RandomInt(0, 10)),
				Latitude:     float64(util.RandomInt(0, 10)),
				CreatedAt:    "2024-03-10 10:00:00",
			}

			response, err := client.CreateVehicle(ctx, request)

			if response != nil {
				if response.Status != tt.res.Status {
					t.Error("response: expected", tt.res.Status, "received", response.Status)
				}
			}

			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != tt.errCode {
						t.Error("error code: expected", codes.InvalidArgument, "received", er.Code())
					}
					if er.Message() != tt.errMsg {
						t.Error("error message: expected", tt.errMsg, "received", er.Message())
					}
				}
			}
		})
	}
}

func TestDeleteVehicleData(t *testing.T) {
	tests := []struct {
		name    string
		input   *pb.ID
		res     *pb.VehicleResponse
		errCode codes.Code
		errMsg  string
	}{
		{
			"Data tidak ditemukan",
			&pb.ID{Id: 1000},
			nil,
			codes.NotFound,
			"Proses hapus data gagal karena data tidak ditemukan",
		},
		{
			"Data berhasil dihapus",
			&pb.ID{Id: 1200},
			&pb.VehicleResponse{Status: "OK"},
			codes.OK,
			"",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewVehiclesCRUDClient(conn)

	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if index == 1 {
				request := &pb.VehicleInput{
					VehicleType:  util.RandomString(6),
					VehicleModel: util.RandomString(8),
					Longitude:    float64(util.RandomInt(0, 10)),
					Latitude:     float64(util.RandomInt(0, 10)),
					CreatedAt:    "2024-03-10 10:00:00",
				}

				client.CreateVehicle(ctx, request)
			}

			request := &pb.ID{Id: 1200}
			response, err := client.DeleteVehicle(ctx, request)

			if response != nil {
				if response.Status != tt.res.Status {
					t.Error("response: expected", tt.res.Status, "received", response.Status)
				}
			}

			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != tt.errCode {
						t.Error("error code: expected", codes.NotFound, "received", er.Code())
					}
					if er.Message() != tt.errMsg {
						t.Error("error message: expected", tt.errMsg, "received", er.Message())
					}
				}
			}
		})
	}
}
