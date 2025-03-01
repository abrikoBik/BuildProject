package db

import (
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global database connection.
var Database *gorm.DB

// Opens a PostgreSQL database with the name dbname at hostname:port.
func Open(hostname string, port int, dbname string) (*gorm.DB, error) {
	var (
		username = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
	)
	if username == "" || password == "" {
		return nil, errors.New("postgres credentials not found in environment")
	}
	// TODO: sslmode=verify-full
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, port, username, password, dbname)
	database, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return database, nil
}
