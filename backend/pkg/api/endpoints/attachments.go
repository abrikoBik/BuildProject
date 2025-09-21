package endpoints

import (
	"mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"

	"github.com/go-chi/chi/v5"
)

func AddAttachments(r chi.Router, jwtKey []byte) {
	r.Route("/attachments", func(r chi.Router) {
		r.With(middleware.JSONResponse).Get("/", handlers.ListProjectAttachments)
		r.With(middleware.JSONResponse).Post("/", handlers.CreateAttachment)
		r.Route("/{attachmentUUID}", func(r chi.Router) {
			r.Use(middleware.ParseUUID(middleware.AttachmentUUID, "attachmentUUID"))
			r.Get("/", handlers.GetAttachment)
			r.Delete("/", handlers.DeleteAttachment)
		})
	})
}
