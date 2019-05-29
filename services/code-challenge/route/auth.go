package route

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aaclee/ms-arch/services/auth/authpb"
	"github.com/gorilla/mux"

	// Required for postgres
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	codeChallengeHost     = "localhost"
	codeChallengePort     = 5433
	codeChallengeUser     = "ms_cc_psql"
	codeChallengePassword = "password"
	codeChallengeDBname   = "code_challenge_db"
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

	psqlCodeChallengeInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		codeChallengeHost,
		codeChallengePort,
		codeChallengeUser,
		codeChallengePassword,
		codeChallengeDBname,
	)

	dbCodeChallenge, err := sql.Open("postgres", psqlCodeChallengeInfo)
	if err != nil {
		log.Fatalf("CODE_CHALLENGE: Could not open connection to database: %v", err)
	}
	defer dbCodeChallenge.Close()

	err = dbCodeChallenge.Ping()
	if err != nil {
		log.Fatalf("CODE_CHALLENGE: Could not ping database: %v", err)
	}

	codeChallengeSelectQuery := `SELECT uuid FROM users WHERE email=$1;`
	row := dbCodeChallenge.QueryRow(codeChallengeSelectQuery, body.Username)

	var uuid string
	var email string
	err = row.Scan(&uuid, &email)
	if err != nil {
		encodeError(w, http.StatusNotFound, "Username or password incorrect!")
		fmt.Println("Could not select query")
		return
	}

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
