package main

import (
	"context"
	"mp1/api"
	"mp1/db"
	"mp1/util"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	util.Check(godotenv.Load())
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)

	// connecting to the database
	database, err := db.Open("localhost", 5432, "mp1")
	db.Database = database
	util.Check(err, db.Migrate(db.User{}))
	util.Log("db", "connected to database")

	// starting the api
	api.Database = database
	apiServer := api.Init(3000)
	util.Log("api", "server started")

	// listening for SIGINT and SIGTERM (graceful shutdown)
	<-exitChan
	apiServer.Shutdown(context.Background())
	util.Log("api", "server shutdown")
	sqlDB, err := database.DB()
	util.Check(err, sqlDB.Close())
}
