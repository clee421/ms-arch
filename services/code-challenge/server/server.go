package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aaclee/ms-arch/services/code-challenge/server/route"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Code Challenge Server")

	config, err := getServerConfigs("./config/config.development.json")
	if err != nil {
		log.Fatalf("Failed to get configs: %v", err)
	}

	r := mux.NewRouter()

	route.Handler(r)

	port := getPort(fmt.Sprintf(":%d", config.Port))
	fmt.Printf("Running server on port %v...\n", port)

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
