package endpoints

import (
	"mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
)

func AddUsers(r chi.Router, jwtKey []byte) {
	r.Route("/users", func(r chi.Router) {
		r.With(middleware.JSONResponse).Post("/", handlers.CreateUser)
		r.Route("/{userID}", func(r chi.Router) {
			r.Use(middleware.ParseID(middleware.UserID, "userID"))
			r.With(middleware.JSONResponse).Get("/", handlers.GetUser)
			r.With(middleware.JSONResponse).Get("/projects", handlers.ListProjects)
			r.With(middleware.JSONResponse).Patch("/", handlers.PatchUser)
			AddUserWorker(r, jwtKey)
		})
	})
}

func AddProjectWorkers(r chi.Router, jwtKey []byte) {
	r.Route("/workers", func(r chi.Router) {
		r.With(middleware.JSONResponse).Get("/", handlers.ListWorkers)
		r.Route("/{workerID}", func(r chi.Router) {
			r.Use(middleware.ParseID(middleware.WorkerID, "workerID"))
			r.Put("/", handlers.AddWorker)
			r.Delete("/", handlers.RemoveWorker)
		})
	})
}
