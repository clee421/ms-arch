package route

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aaclee/ms-arch/services/auth/authpb"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type authRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authResponse struct {
	Token string `json:"token"`
}

// AuthRoutes for the authentication path
func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth", authenticate).Methods("POST")
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var body authRequestBody
	err := decoder.Decode(&body)

	if err != nil {
		fmt.Println("Error with post body")
	}

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := authpb.NewAuthServiceClient(cc)

	req := &authpb.AuthenticateRequest{
		Auth: &authpb.Authentication{
			Username: body.Username,
			Password: body.Password,
		},
	}
	res, err := c.Authenticate(context.Background(), req)
	if err != nil {
		encodeError(w, http.StatusNotFound, "Username or password incorrect!")
	} else {
		encode(w, authResponse{res.Token})
	}
}
