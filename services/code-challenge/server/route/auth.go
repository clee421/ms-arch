package route

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aaclee/ms-arch/services/auth/authpb"
	"github.com/aaclee/ms-arch/services/code-challenge/service/ccpb"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	// Required for postgres
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

type authRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authResponse struct {
	Token string `json:"token"`
}

// AuthRoutes for the authentication path
func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth", authenticate).Methods(http.MethodPost, http.MethodOptions)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	if r.Method == http.MethodOptions {
		return
	}

	decoder := json.NewDecoder(r.Body)

	var body authRequestBody
	err := decoder.Decode(&body)

	if err != nil {
		log.Info("Error with post body")
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
		log.Infof("Could not find user: %v\n", err)
		encodeError(w, http.StatusUnauthorized, "Username or password incorrect!")
		return
	}

	authRequest := &authpb.AuthenticateRequest{
		Auth: &authpb.Authentication{
			Uuid:     ccRes.User.GetUuid(),
			Password: body.Password,
		},
	}
	authRes, err := authClient.Authenticate(context.Background(), authRequest)
	if err != nil {
		encodeError(w, http.StatusUnauthorized, "Username or password incorrect!")
	} else {
		encode(w, authResponse{authRes.Token})
	}
}
