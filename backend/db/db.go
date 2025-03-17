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
func Open(hostname string, port int, dbname string) error {
	var (
		username, password string
		exists             bool
	)
	if username, exists = os.LookupEnv("POSTGRES_USER"); !exists {
		return errors.New("POSTGRES_USER not found in environment")
	}
	if password, exists = os.LookupEnv("POSTGRES_PASSWORD"); !exists {
		return errors.New("POSTGRES_PASSWORD not found in environment")
	}
	// TODO: sslmode=verify-full
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, port, username, password, dbname)
	database, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}
	Database = database
	return nil
}
