package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/aaclee/ms-arch/services/auth/authpb"

	_ "github.com/lib/pq"
)

// Server type for Auth Service
type Server struct {
	Config Configuration
}

// Authenticate will validate username with password
func (server *Server) Authenticate(ctx context.Context, req *authpb.AuthenticateRequest) (*authpb.AuthenticateResponse, error) {
	psqlAuthInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		server.Config.Database.Host,
		server.Config.Database.Port,
		server.Config.Database.Username,
		server.Config.Database.Password,
		server.Config.Database.Name,
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
			"iss": server.Config.Jwt.Issuer,
			"sub": string(uuid),
			"aud": "msp.code_challenge",
			"exp": time.Now().Add(time.Hour * 8),
			"nbf": time.Now(),
			"iat": time.Now(),
		})

		// TODO: Issue 500 error if token signing failed
		tokenString, err := token.SignedString([]byte(server.Config.Jwt.Secret))

		fmt.Println(tokenString, err)
		response = &authpb.AuthenticateResponse{
			Token: tokenString,
		}
	}

	return response, err
}
