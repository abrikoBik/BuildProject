package api

import (
	"mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
)

func addAttachments(r chi.Router) {
	r.Route("/attachments", func(r chi.Router) {
		r.With(middleware.JSONResponse).Post("/", handlers.CreateAttachment)
		r.Route("/{attachmentUUID}", func(r chi.Router) {
			r.Use(middleware.ParseID(middleware.AttachmentUUID, "attachmentUUID"))
			r.With(middleware.JSONResponse, middleware.ValidateToken(jwtKey)).Get("/", handlers.GetAttachment)
			r.With(middleware.ValidateToken(jwtKey)).Delete("/", handlers.DeleteAttachment)
		})
	})
}
