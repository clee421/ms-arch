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
	host     = "localhost"
	port     = 5432
	user     = "ms_auth_psql"
	password = "password"
	dbname   = "auth_db"
)

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Could not open connection to database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not ping database: %v", err)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Could not create password hash: %v", err)
	}

	sqlStatement := `
		INSERT INTO auth (user_id, password)
		VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, uuid.New(), passwordHash)
	if err != nil {
		panic(err)
	}
}
