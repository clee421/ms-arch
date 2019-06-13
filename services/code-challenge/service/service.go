package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aaclee/ms-arch/services/code-challenge/service/ccpb"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// Server type for Auth Service
type Server struct {
	config Configuration
	log    log.Logger
}

// User will get user by username
func (server *Server) User(ctx context.Context, req *ccpb.UserRequest) (*ccpb.UserResponse, error) {
	username := req.GetUsername()

	server.log.Infof("Fetching user from database: %v\n", username)

	// Database fetch
	psqlCodeChallengeInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		server.config.Database.Host,
		server.config.Database.Port,
		server.config.Database.Username,
		server.config.Database.Password,
		server.config.Database.Name,
	)

	dbCodeChallenge, err := sql.Open("postgres", psqlCodeChallengeInfo)
	if err != nil {
		server.log.Fatalf("CODE_CHALLENGE: Could not open connection to database: %v", err)
	}
	defer dbCodeChallenge.Close()

	err = dbCodeChallenge.Ping()
	if err != nil {
		server.log.Fatalf("CODE_CHALLENGE: Could not ping database: %v", err)
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
