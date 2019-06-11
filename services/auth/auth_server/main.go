package main

import (
	"fmt"
	"net"
	"os"

	"github.com/aaclee/ms-arch/services/auth/authpb"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("Authentication Server: Status - Booting")

	config, err := getServerConfigs("./auth_server/config/config.development.json")
	if err != nil {
		log.Fatalf("Failed to get configs: %v", err)
	}

	address := fmt.Sprintf("%v:%d", config.Host, config.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Infof("Listening on address: %s\n", address)
	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &Server{*config, *log.New()})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
