package main

import (
	"fmt"
	"net"
	"os"

	"github.com/aaclee/ms-arch/services/code-challenge/service/ccpb"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	configFile = "./config/config.development.json"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Info("Code Challenge Service: Status - Booting")

	log.Infof("Loading configs from: %s\n", configFile)
	config, err := getServerConfigs(configFile)
	if err != nil {
		log.Fatalf("Failed to get configs: %v", err)
	}

	address := fmt.Sprintf("%v:%d", config.Host, config.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Infof("Listening on address: %v\n", address)
	s := grpc.NewServer()
	ccpb.RegisterUserServiceServer(s, &Server{*config, *log.New()})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
