package endpoints

import (
	"mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
)

func AddSessions(r chi.Router, jwtKey []byte) {
	r.Route("/sessions", func(r chi.Router) {
		r.With(middleware.JSONResponse).Post("/", handlers.CreateSession)
		r.Delete("/", handlers.DeleteSession)
	})
}
