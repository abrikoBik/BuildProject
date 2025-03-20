package main

import (
	"database/sql"
	"errors"
	"flag"
	"mp1/api"
	"mp1/db"
	"mp1/models"
	"mp1/util"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const SERVICES_COUNT = 2

func main() {
	// TODO: provide flag usage
	var (
		dbName      = flag.String("db-name", "mp1", "")
		dbHostname  = flag.String("db-hostname", "0.0.0.0", "")
		dbPort      = flag.Uint("db-port", 5432, "")
		apiHostname = flag.String("api-hostname", "0.0.0.0", "")
		apiPort     = flag.Uint("api-port", 3000, "")
		migrate     = flag.Bool("migrate", false, "")
		logLevel    = flag.String("log-level", "info", "")

		servicesWg sync.WaitGroup
		sqlDB      *sql.DB
		err        error
	)
	flag.Parse()
	util.Check(util.LogLevel(*logLevel))

	servicesWg.Add(SERVICES_COUNT)
	// connecting to the database
	go func() {
		util.Info("db", "connecting to %s @ %s:%d", *dbName, *dbHostname, *dbPort)
		util.Check(db.Open(*dbHostname, *dbPort, *dbName))
		sqlDB, err = db.Database.DB()
		util.Check(err)
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
			util.Check(err, sqlDB.Close())
			os.Exit(0)
		}

		servicesWg.Done()
	}()

	// starting the api
	var server *http.Server
	go func() {
		if *migrate {
			return
		}
		server = api.Init(*apiHostname, *apiPort)
		util.Info("api", "starting @ %s", server.Addr)
		listener, err := net.Listen("tcp", server.Addr)
		util.Check(err)
		util.Info("api", "started")
		servicesWg.Done()
		if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
			util.Error("api", err)
		}
	}()

	servicesWg.Wait()
	// graceful shutdown
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)
	<-exitChan
	servicesWg.Add(SERVICES_COUNT)
	go func() {
		util.Info("api", "stopping")
		util.Check(server.Close())
		util.Info("api", "stopped")
		servicesWg.Done()
	}()
	go func() {
		util.Info("db", "stopping")
		util.Check(err, sqlDB.Close())
		util.Info("db", "stopped")
		servicesWg.Done()
	}()
	servicesWg.Wait()
}
