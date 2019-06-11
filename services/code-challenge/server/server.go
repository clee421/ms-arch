package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aaclee/ms-arch/services/code-challenge/server/route"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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
	log.Info("Code Challenge Server: Status - Booting")

	log.Infof("Loading configs from: %s\n", configFile)
	config, err := getServerConfigs(configFile)
	if err != nil {
		log.Fatalf("Failed to get configs: %v", err)
	}

	r := mux.NewRouter()

	route.Handler(r)

	port := getPort(fmt.Sprintf(":%d", config.Port))
	log.Infof("Running server on port %v...\n", port)

	// Pass router into the server
	http.ListenAndServe(port, r)
}

func getPort(fallback string) string {
	e := os.Getenv("PORT")
	if e == "" {
		return fallback
	}

	return e
}
