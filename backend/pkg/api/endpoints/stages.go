package endpoints

import (
	// "mp1/pkg/api/handlers"
	"mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
)

func AddStages(r chi.Router, jwtKey []byte) {
	r.Route("/stages", func(r chi.Router) {
		r.With(middleware.JSONResponse).Post("/", handlers.CreateStage)
		r.Route("/{stageID}", func(idRouter chi.Router) {
			idRouter.Use(middleware.ParseID(middleware.StageID, "stageID"))
			idRouter.With(middleware.JSONResponse).Get("/", handlers.GetStage)
			idRouter.With(middleware.JSONResponse).Patch("/", handlers.PatchStage)
			idRouter.Delete("/", handlers.DeleteStage)
		})
	})
}

func AddStageGroups(r chi.Router, jwtKey []byte) {
	r.Route("/stageGroups", func(r chi.Router) {
		r.Post("/", handlers.CreateStageGroup)
		r.Route("/{stageGroupID}", func(idRouter chi.Router) {
			idRouter.Use(middleware.ParseID(middleware.StageGroupID, "stageGroupID"))
			idRouter.With(middleware.JSONResponse).Get("/", handlers.GetStageGroup)
			idRouter.With(middleware.JSONResponse).Patch("/", handlers.PatchStageGroup)
			idRouter.Delete("/", handlers.DeleteStageGroup)
			AddStages(r, jwtKey)
		})
	})
}
