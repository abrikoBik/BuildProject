package main

import (
	"flag"
	"mp1/cmd/api"
	"mp1/cmd/db"
	"mp1/pkg/util"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const SERVICES_COUNT = 2

func main() {
	var (
		servicesWg  sync.WaitGroup
		dbName      = flag.String("db-name", "mp1", "PostgreSQL database name")
		dbHostname  = flag.String("db-hostname", "0.0.0.0", "PostgreSQL database IP")
		dbPort      = flag.Uint("db-port", 5432, "PostgreSQL database port")
		apiHostname = flag.String("api-hostname", "0.0.0.0", "Hostname for the API to listen to")
		apiPort     = flag.Uint("api-port", 3000, "Port for the API to listen to")
		migrate     = flag.Bool("migrate", false, "Run migrations (exits after successful migrations)")
		logLevel    = flag.String("log-level", "info", "General log level (debug, warn, info)")
	)
	flag.Parse()
	util.Check(util.LogLevel(*logLevel))

	servicesWg.Add(SERVICES_COUNT)
	go db.Start(*dbHostname, *dbPort, *dbName, servicesWg.Done)
	if *migrate {
		servicesWg.Done()
		servicesWg.Wait()
		db.Migrate(func() {
			db.Stop(func() { os.Exit(0) })
		})
	}
	go api.Start(*apiHostname, *apiPort, servicesWg.Done)
	servicesWg.Wait()

	// graceful shutdown
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, syscall.SIGTERM)
	<-exitChan
	servicesWg.Add(SERVICES_COUNT)
	go api.Stop(servicesWg.Done)
	go db.Stop(servicesWg.Done)
	servicesWg.Wait()
}
