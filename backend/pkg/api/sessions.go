package api

import (
	"mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"
	"mp1/pkg/util"

	"github.com/go-chi/chi/v5"
)

func addSessions(r chi.Router) {
	r.Route("/sessions", func(r chi.Router) {
		r.With(middleware.JSONResponse).Post("/", handlers.Login)
		r.Delete("/", handlers.Logout)
	})
	util.Debug("api:addSessions", "plugged in")
}
