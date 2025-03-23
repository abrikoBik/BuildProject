package api

import (
	// "mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
)

func addWorkers(r chi.Router) {
	r.Route("/workers", func(r chi.Router) {
		// r.With(middleware.JSONResponse).Get("/", handlers.GetWorkers)
		r.Route("/{workerID}", func(r chi.Router) {
			r.Use(middleware.ParseID("workerID", "workerID"))
			// r.With(middleware.JSONResponse).Get("/", handlers.GetWorker)
		})
	})
}
