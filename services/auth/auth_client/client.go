package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aaclee/ms-arch/services/auth/authpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Authentication Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := authpb.NewAuthServiceClient(cc)
	// fmt.Printf("Created client %f", c)

	req := &authpb.AuthenticateRequest{
		Auth: &authpb.Authentication{
			Username: "bill.jobs@apple.com",
			Password: "password",
		},
	}
	res, err := c.Authenticate(context.Background(), req)
	if err != nil {
		log.Fatalf("Error from authenticating: %v\n", err)
	}

	log.Printf("Token: %v", res.Token)
}
