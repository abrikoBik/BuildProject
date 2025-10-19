package endpoints

import (
	"mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
)

func AddUserWorker(r chi.Router, jwtKey []byte) {
	r.Route("/worker", func(r chi.Router) {
		r.With(middleware.JSONResponse).Get("/", handlers.GetWorker)
		r.With(middleware.JSONResponse).Post("/", handlers.CreateWorker)
		r.With(middleware.JSONResponse).Patch("/", handlers.PatchWorker)
	})
}
