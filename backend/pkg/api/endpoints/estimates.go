package endpoints

import (
	"mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
)

func AddEstimates(r chi.Router, jwtKey []byte) {
	r.Route("/estimates", func(r chi.Router) {
		r.Post("/", handlers.CreateEstimate)
		r.With(middleware.JSONResponse).Get("/", handlers.ListEstimates)
		r.Route("/{estimateID}", func(idRouter chi.Router) {
			idRouter.Use(middleware.ParseID(middleware.EstimateID, "estimateID"))
			idRouter.With(middleware.JSONResponse).Get("/", handlers.GetEstimate)
			idRouter.With(middleware.JSONResponse).Patch("/", handlers.PatchEstimate)
			idRouter.Delete("/", handlers.DeleteEstimate)
		})
	})
}

func AddEstimateGroups(r chi.Router, jwtKey []byte) {
	r.Route("/estimateGroups", func(r chi.Router) {
		r.Post("/", handlers.CreateEstimateGroup)
		r.With(middleware.JSONResponse).Get("/", handlers.ListEstimateGroups)
		r.Route("/{estimateGroupID}", func(idRouter chi.Router) {
			idRouter.Use(middleware.ParseID(middleware.EstimateGroupID, "estimateGroupID"))
			idRouter.With(middleware.JSONResponse).Get("/", handlers.GetEstimateGroup)
			idRouter.With(middleware.JSONResponse).Patch("/", handlers.PatchEstimateGroup)
			idRouter.Delete("/", handlers.DeleteEstimateGroup)
			AddEstimates(idRouter, jwtKey)
		})
	})
}
