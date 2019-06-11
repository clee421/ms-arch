package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/aaclee/ms-arch/services/auth/authpb"
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

// Server type for Auth Service
type Server struct {
	config Configuration
	log    log.Logger
}

// Authenticate will validate username with password
func (server *Server) Authenticate(ctx context.Context, req *authpb.AuthenticateRequest) (*authpb.AuthenticateResponse, error) {
	psqlAuthInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		server.config.Database.Host,
		server.config.Database.Port,
		server.config.Database.Username,
		server.config.Database.Password,
		server.config.Database.Name,
	)

	dbAuth, err := sql.Open("postgres", psqlAuthInfo)
	if err != nil {
		server.log.Fatalf("AUTH: Could not open connection to database: %v", err)
	}
	defer dbAuth.Close()

	err = dbAuth.Ping()
	if err != nil {
		server.log.Fatalf("AUTH: Could not ping database: %v", err)
	}

	uuid := req.GetAuth().GetUuid()
	password := req.GetAuth().GetPassword()

	authSelectQuery := `SELECT password FROM auth WHERE user_uuid=$1;`
	row := dbAuth.QueryRow(authSelectQuery, uuid)

	var passwordHash string
	err = row.Scan(&passwordHash)

	server.log.Infof("Authenticating: %s\n", uuid)

	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	}

	var response *authpb.AuthenticateResponse
	if err != nil {
		server.log.Infof("Error Authenticating %v: %v\n", passwordHash, err)
		err = errors.New("Invalid username or password")
		response = nil
	} else {
		// TODO: Add email / username to this claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss": server.config.Jwt.Issuer,
			"sub": string(uuid),
			"aud": "msp.code_challenge",
			"exp": time.Now().Add(time.Hour * 8),
			"nbf": time.Now(),
			"iat": time.Now(),
		})

		// TODO: Issue 500 error if token signing failed
		tokenString, _ := token.SignedString([]byte(server.config.Jwt.Secret))

		// fmt.Println(tokenString, err)
		response = &authpb.AuthenticateResponse{
			Token: tokenString,
		}
	}

	return response, err
}
