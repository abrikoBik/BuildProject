package endpoints

import (
	"mp1/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
)

func AddEstimates(r chi.Router, jwtKey []byte) {
	r.Route("/estimates", func(r chi.Router) {
		r.Route("/{estimateID}", func(idRouter chi.Router) {
			idRouter.Use(middleware.ParseID(middleware.EstimateID, "estimateID"))
		})
	})
}

func AddEstimateGroups(r chi.Router, jwtKey []byte) {
	r.Route("/estimateGroups", func(r chi.Router) {
		r.Route("/{estimateGroupID}", func(idRouter chi.Router) {
			idRouter.Use(middleware.ParseID(middleware.EstimateGroupID, "estimateGroupID"))
			AddEstimates(r, jwtKey)
		})
	})
}
