package main

import (
	"fmt"
	"log"
	"net"

	"github.com/aaclee/ms-arch/services/auth/authpb"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Authentication Server")

	config, err := getServerConfigs("./auth_server/config/config.development.json")
	if err != nil {
		log.Fatalf("Failed to get configs: %v", err)
	}

	address := fmt.Sprintf("%v:%d", config.Host, config.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Printf("Listening on address: %v\n", address)
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &Server{*config})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
