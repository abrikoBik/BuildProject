package api

import (
	"mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
)

func addSessions(r chi.Router) {
	r.Route("/sessions", func(r chi.Router) {
		r.With(middleware.JSONResponse).Post("/", handlers.Login)
		r.Delete("/", handlers.Logout)
	})
}
