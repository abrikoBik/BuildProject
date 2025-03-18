package main

import (
	"errors"
	"flag"
	"mp1/api"
	"mp1/db"
	"mp1/models"
	"mp1/util"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	// TODO: provide flag usage
	var (
		envFile  = flag.String("env-file", ".env", "")
		hostname = flag.String("hostname", "localhost", "")
		dbName   = flag.String("db-name", "mp1", "")
		dbPort   = flag.Uint("db-port", 5432, "")
		apiPort  = flag.Uint("api-port", 3000, "")
		migrate  = flag.Bool("migrate", false, "")
		logLevel = flag.String("log-level", "info", "")
	)
	flag.Parse()
	util.Check(util.LogLevel(*logLevel))

	util.Info("*", "loading environment variables")
	util.Check(godotenv.Load(*envFile))
	util.Info("*", "loaded environment variables")

	// connecting to the database
	go func() {
		util.Info("db", "connecting to %s @ %s:%d", *dbName, *hostname, *dbPort)
		util.Check(db.Open(*hostname, *dbPort, *dbName))
		util.Info("db", "connected")
		if *migrate {
			util.Info("db", "running migrations")
			toMigrate := map[string]interface{}{
				"user":    models.User{},
				"session": models.Session{},
				"worker":  models.Worker{},
				"stage":   models.Stage{},
				"project": models.Project{},
			}
			for name, model := range toMigrate {
				util.Check(db.Database.AutoMigrate(&model))
				util.Debug("db", "model", name)
			}
			util.Info("db", "migrations complete")
		}
	}()

	// starting the api
	util.Info("api", "starting @ %s:%d", *hostname, *apiPort)
	server := api.Init(*hostname, *apiPort)
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			util.Error("api", err)
		}
	}()

	// graceful shutdown
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)
	<-exitChan
	var stoppedWg sync.WaitGroup
	stoppedWg.Add(2)
	go func() {
		util.Info("api", "stopping")
		util.Check(server.Close())
		util.Info("api", "stopped")
		stoppedWg.Done()
	}()
	go func() {
		util.Info("db", "stopping")
		sqlDB, err := db.Database.DB()
		util.Check(err, sqlDB.Close())
		util.Info("db", "stopped")
		stoppedWg.Done()
	}()
	stoppedWg.Wait()
}
