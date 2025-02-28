package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// Opens a PostgreSQL database with the name dbname at hostname:port.
func Open(hostname string, port int, dbname string) (*sql.DB, error) {
	var (
		username = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
	)
	if username == "" || password == "" {
		err := errors.New("postgres credentials not found in environment")
		return nil, err
	}
	// TODO: sslmode=verify-full
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, port, username, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
