package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/aaclee/ms-arch/services/code-challenge/service/ccpb"
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

// Server type for Auth Service
type Server struct{}

// User will get user by username
func (*Server) User(ctx context.Context, req *ccpb.UserRequest) (*ccpb.UserResponse, error) {
	username := req.GetUsername()

	fmt.Printf("Fetching user from database: %v\n", username)

	// Database fetch
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

	codeChallengeSelectQuery := `SELECT id, uuid, email FROM users WHERE email=$1;`
	row := dbCodeChallenge.QueryRow(codeChallengeSelectQuery, username)

	var id uint32
	var uuid []byte
	var email string
	err = row.Scan(&id, &uuid, &email)

	var response *ccpb.UserResponse

	if err != nil {
		// err = errors.New("Username does not exist")
		response = nil
	} else {
		response = &ccpb.UserResponse{
			User: &ccpb.User{
				Id:    id,
				Uuid:  uuid,
				Email: email,
			},
		}
	}

	return response, err
}

func main() {
	fmt.Println("Code Challenge Service")

	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	ccpb.RegisterUserServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
