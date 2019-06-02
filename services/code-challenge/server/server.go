package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aaclee/ms-arch/services/code-challenge/server/route"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Code Challenge Server")
	// New router from gorilla-mux
	r := mux.NewRouter()

	route.Handler(r)

	port := getPort(":8000")
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
