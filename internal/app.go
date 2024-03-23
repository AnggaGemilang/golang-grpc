package internal

import (
	"fmt"
	"log"
	"net"

	"len/go-grpc-api/config"
	"len/go-grpc-api/internal/server"
	pb "len/go-grpc-api/proto"

	"google.golang.org/grpc"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	config.SetupDB()
}

func Run() {
	setConfiguration("config.yml")
	conf := config.GetConfig()

	fmt.Println("Go gRPC Running on port " + conf.Server.Port)
	fmt.Println("==================>")

	listen, err := net.Listen("tcp", ":"+conf.Server.Port)
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterVehiclesCRUDServer(s, server.NewVehiclesCRUDServer())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
