package api

import (
	"mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"
	"mp1/pkg/util"

	"github.com/go-chi/chi/v5"
)

func addUsers(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.With(middleware.JSONResponse).Post("/", handlers.RegisterUser)
		r.Route("/{userID}", func(r chi.Router) {
			r.Use(middleware.ParseID(middleware.UserID))
			r.With(middleware.JSONResponse).Get("/", handlers.GetUser)
			r.Patch("/", handlers.PatchUser)
			r.Route("/worker", func(r chi.Router) {
				addWorker(r)
			})
		})
	})
	util.Debug("api:addUsers", "plugged in")
}
