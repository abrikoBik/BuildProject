package endpoints

import (
	"mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func AddDebug(r chi.Router) {
	r.Route("/debug", func(r chi.Router) {
		r.Mount("/", chiMiddleware.Profiler())
		r.With(middleware.JSONResponse).Get("/info", handlers.DebugInfo)
		r.With(middleware.JSONResponse).Get("/gc", handlers.DebugGC)
	})
}
