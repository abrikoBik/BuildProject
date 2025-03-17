package main

import (
	"errors"
	"mp1/api"
	"mp1/db"
	"mp1/models"
	"mp1/util"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

const (
	API_PORT = 3000
	DB_HOST  = "localhost"
	DB_PORT  = 5432
	DB_NAME  = "mp1"
)

func main() {
	// connecting to the database
	util.Info("db", "connecting to %s @ %s:%d", DB_NAME, DB_HOST, DB_PORT)
	go func() {
		util.Check(
			godotenv.Load(),
			db.Open(DB_HOST, DB_PORT, DB_NAME),
			db.Migrate(models.User{}, models.Session{}, models.Worker{}, models.Stage{}, models.Project{}),
		)
		util.Info("db", "connected")
	}()

	// starting the api
	server := api.Init(API_PORT)
	util.Info("api", "starting @ :%d", API_PORT)
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			util.Error("api", err)
		}
		util.Info("api", "stopped")
	}()

	// graceful shutdown
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)
	<-exitChan

	util.Info("api", "stopping")
	util.Check(server.Close())
	util.Info("db", "stopping")
	sqlDB, err := db.Database.DB()
	util.Check(err, sqlDB.Close())
}
