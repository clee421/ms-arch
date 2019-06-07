package main

import (
	"fmt"
	"log"
	"net"

	"github.com/aaclee/ms-arch/services/code-challenge/service/ccpb"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Code Challenge Service")

	config, err := getServerConfigs("./config/config.development.json")
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
	ccpb.RegisterUserServiceServer(s, &Server{*config})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
