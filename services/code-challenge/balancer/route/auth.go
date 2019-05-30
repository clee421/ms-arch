package route

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aaclee/ms-arch/services/auth/authpb"
	"github.com/aaclee/ms-arch/services/code-challenge/service/ccpb"
	"github.com/gorilla/mux"

	// Required for postgres
	_ "github.com/lib/pq"
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

	authCC, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to auth service: %v", err)
	}

	defer authCC.Close()

	authClient := authpb.NewAuthServiceClient(authCC)

	codeChallengeCC, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to auth service: %v", err)
	}

	defer codeChallengeCC.Close()

	codeChallengeClient := ccpb.NewUserServiceClient(codeChallengeCC)

	userRequest := &ccpb.UserRequest{
		Username: body.Username,
	}
	ccRes, err := codeChallengeClient.User(context.Background(), userRequest)
	if err != nil {
		fmt.Printf("Could not find user: %v\n", err)
	} else {
		fmt.Printf("User found: %v\n", ccRes.User)
	}

	authRequest := &authpb.AuthenticateRequest{
		Auth: &authpb.Authentication{
			Username: body.Username,
			Password: body.Password,
		},
	}
	authRes, err := authClient.Authenticate(context.Background(), authRequest)
	if err != nil {
		encodeError(w, http.StatusNotFound, "Username or password incorrect!")
	} else {
		encode(w, authResponse{authRes.Token})
	}
}
