package endpoints

import (
	// "mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
)

func AddStages(r chi.Router, jwtKey []byte) {
	r.Route("/stages", func(r chi.Router) {
		// r.With(middleware.JSONResponse).Get("/", handlers.GetStages)
		// r.With(middleware.JSONResponse).Post("/", handlers.CreateStage)
		r.Route("/{stageID}", func(idRouter chi.Router) {
			idRouter.Use(middleware.ParseID(middleware.StageID, "stageID"))
			// idRouter.With(middleware.JSONResponse).Get("/", handlers.GetStage)
			// idRouter.Patch("/", handlers.UpdateStage)
			// idRouter.Delete("/", handlers.DeleteStage)
		})
	})
}

func AddStageGroups(r chi.Router, jwtKey []byte) {
	r.Route("/stageGroups", func(r chi.Router) {
		r.Route("/{stageGroupID}", func(idRouter chi.Router) {
			idRouter.Use(middleware.ParseID(middleware.StageGroupID, "stageGroupID"))
			AddStages(r, jwtKey)
		})
	})
}
