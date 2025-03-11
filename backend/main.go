package main

import (
	"mp1/api"
	"mp1/db"
	"mp1/models"
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
	util.Check(err, db.Migrate(models.UserDB{}, models.Session{}, models.WorkerDB{}, models.Stage{}, models.Project{}))
	util.Log("db", "connected to database")

	// starting the api
	api.Database = database
	util.Check(api.Init(3000))

	// listening for SIGINT and SIGTERM (graceful shutdown)
	<-exitChan
	sqlDB, err := database.DB()
	util.Check(err, sqlDB.Close())
}
