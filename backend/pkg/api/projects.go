package api

import (
	"mp1/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
)

func addProjects(r chi.Router) {
	r.Route("/projects", func(r chi.Router) {
		// r.With(middleware.JSONResponse).Post("/", handlers.CreateProject)
		r.Route("/{projectID}", func(idRouter chi.Router) {
			idRouter.Use(middleware.ParseID(middleware.ProjectID, "projectID"))
			// r.With(middleware.JSONResponse).Get("/", handlers.GetProject)
			addStages(idRouter)
			// addWorkers(idRouter)
		})
	})
}
