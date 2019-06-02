package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/aaclee/ms-arch/services/auth/authpb"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	authHost     = "localhost"
	authPort     = 5432
	authUser     = "ms_auth_psql"
	authPassword = "password"
	authDBname   = "auth_db"
)

// Server type for Auth Service
type Server struct{}

// Authenticate will validate username with password
func (*Server) Authenticate(ctx context.Context, req *authpb.AuthenticateRequest) (*authpb.AuthenticateResponse, error) {
	psqlAuthInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		authHost,
		authPort,
		authUser,
		authPassword,
		authDBname,
	)

	dbAuth, err := sql.Open("postgres", psqlAuthInfo)
	if err != nil {
		log.Fatalf("AUTH: Could not open connection to database: %v", err)
	}
	defer dbAuth.Close()

	err = dbAuth.Ping()
	if err != nil {
		log.Fatalf("AUTH: Could not ping database: %v", err)
	}

	uuid := req.GetAuth().GetUuid()
	password := req.GetAuth().GetPassword()

	authSelectQuery := `SELECT password FROM auth WHERE user_uuid=$1;`
	row := dbAuth.QueryRow(authSelectQuery, uuid)

	var passwordHash string
	err = row.Scan(&passwordHash)

	fmt.Printf("Authenticating: %s\n", uuid)

	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	}

	var response *authpb.AuthenticateResponse
	if err != nil {
		fmt.Printf("Error Authenticating %v: %v\n", passwordHash, err)
		err = errors.New("Invalid username or password")
		response = nil
	} else {
		// TODO: Add email / username to this claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss": "msp.auth_service",
			"sub": string(uuid),
			"aud": "msp.code_challenge",
			"exp": time.Now().Add(time.Hour * 8),
			"nbf": time.Now(),
			"iat": time.Now(),
		})

		// TODO: Issue 500 error if token signing failed
		tokenString, err := token.SignedString([]byte("msp_secret"))

		fmt.Println(tokenString, err)
		response = &authpb.AuthenticateResponse{
			Token: tokenString,
		}
	}

	return response, err
}

func main() {
	fmt.Println("Authentication Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
