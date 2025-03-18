package main

import (
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
		hostname = flag.String("hostname", "0.0.0.0", "")
		dbName   = flag.String("db-name", "mp1", "")
		dbPort   = flag.Uint("db-port", 5432, "")
		apiPort  = flag.Uint("api-port", 3000, "")
		migrate  = flag.Bool("migrate", false, "")
		logLevel = flag.String("log-level", "info", "")
	)
	flag.Parse()
	util.Check(util.LogLevel(*logLevel))

	var servicesWg sync.WaitGroup
	servicesWg.Add(SERVICES_COUNT)
	// connecting to the database
	go func() {
		util.Info("db", "connecting to %s @ %s:%d", *dbName, "172.19.0.1", *dbPort)
		util.Check(db.Open("172.19.0.1", *dbPort, *dbName))
		util.Info("db", "connected")
		if *migrate {
			servicesWg.Add(1)
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
			servicesWg.Done()
		}
		servicesWg.Done()
	}()

	// starting the api
	var server *http.Server
	go func() {
		server = api.Init(*hostname, *apiPort)
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
		sqlDB, err := db.Database.DB()
		util.Check(err, sqlDB.Close())
		util.Info("db", "stopped")
		servicesWg.Done()
	}()
	servicesWg.Wait()
}
