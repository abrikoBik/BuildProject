package api

import (
	"mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
)

func addUsers(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.With(middleware.JSONResponse).Post("/", handlers.Register)
		r.Route("/{userID}", func(r chi.Router) {
			r.Use(middleware.ParseID("userID", "userID"))
			r.With(middleware.JSONResponse).Get("/", handlers.GetUser)
			// r.Patch("/", handlers.PatchUser)
		})
	})
}
