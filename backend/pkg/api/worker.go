package api

import (
	// "mp1/pkg/api/handlers"
	"mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"
	"mp1/pkg/util"

	"github.com/go-chi/chi/v5"
)

func addWorker(r chi.Router) {
	r.With(middleware.JSONResponse).Post("/", handlers.RegisterWorker)
	// r.With(middleware.JSONResponse).Get("/", handlers.GetWorkers)
	r.Route("/{workerID}", func(r chi.Router) {
		r.Use(middleware.ParseID(middleware.WorkerID))
		// r.With(middleware.JSONResponse).Get("/", handlers.GetWorker)
	})
	util.Debug("api:addWorker", "plugged in")
}
