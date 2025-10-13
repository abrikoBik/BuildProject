package db

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// Global database connection.
	Database *gorm.DB

	username, password string
)

func init() {
	var exists bool
	if username, exists = os.LookupEnv("POSTGRES_USER"); !exists {
		panic("POSTGRES_USER not found in environment")
	}
	if password, exists = os.LookupEnv("POSTGRES_PASSWORD"); !exists {
		panic("POSTGRES_PASSWORD not found in environment")
	}
}

// Opens a PostgreSQL database with the name dbname at hostname:port.
func Open(hostname string, port uint, dbname string) error {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=prefer",
		hostname,
		port,
		username,
		password,
		dbname,
	)
	opt := &gorm.Config{
		PrepareStmt:          true,
		DisableAutomaticPing: true,
		TranslateError:       true,
	}
	database, err := gorm.Open(postgres.Open(connStr), opt)
	if err != nil {
		return err
	}
	Database = database
	return nil
}
