package api

import (
	// "mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"
	"mp1/pkg/util"

	"github.com/go-chi/chi/v5"
)

func addStages(r chi.Router) {
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
	util.Debug("api:addStages", "plugged in")
}
