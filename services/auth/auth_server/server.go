package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/aaclee/ms-arch/services/auth/authpb"

	"google.golang.org/grpc"
)

// Server type for Auth Service
type Server struct{}

// Authenticate will validate username with password
func (*Server) Authenticate(ctx context.Context, req *authpb.AuthenticateRequest) (*authpb.AuthenticateResponse, error) {
	username := req.GetAuth().GetUsername()
	password := req.GetAuth().GetPassword()

	fmt.Printf("Authenticating: %v\n", username)

	var response *authpb.AuthenticateResponse
	var err error
	if username == "steve.jobs@apple.com" && password == "password" {
		err = nil
		response = &authpb.AuthenticateResponse{
			Token: "928refhunasf89ys9d8f",
		}
	} else {
		err = errors.New("Invalid username or password")
		response = nil
	}

	return response, err
}

func main() {
	fmt.Println("Authentication Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
