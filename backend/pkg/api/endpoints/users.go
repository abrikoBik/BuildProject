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
			r.With(middleware.ValidateToken(jwtKey)).Patch("/", handlers.PatchUser)
			r.Route("/worker", func(r chi.Router) {
				AddWorker(r, jwtKey)
			})
		})
	})
}
