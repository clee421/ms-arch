package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	authHost     = "localhost"
	authPort     = 5432
	authUser     = "ms_auth_psql"
	authPassword = "password"
	authDBname   = "auth_db"

	codeChallengeHost     = "localhost"
	codeChallengePort     = 5433
	codeChallengeUser     = "ms_cc_psql"
	codeChallengePassword = "password"
	codeChallengeDBname   = "code_challenge_db"
)

func main() {
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

	adminUUID := uuid.New()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("AUTH: Could not create password hash: %v", err)
	}

	authInsertQuery := `
		INSERT INTO auth (user_uuid, password)
		VALUES ($1, $2)`
	_, err = dbAuth.Exec(authInsertQuery, adminUUID, passwordHash)
	if err != nil {
		log.Fatalf("AUTH: Could not insert value: %v", err)
	}

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

	codeChallengeInsertQuery := `
		INSERT INTO users (uuid, email)
		VALUES ($1, $2)`
	_, err = dbCodeChallenge.Exec(codeChallengeInsertQuery, adminUUID, "admin@email.com")
	if err != nil {
		log.Fatalf("CODE_CHALLENGE: Could not insert value: %v", err)
	}
}
