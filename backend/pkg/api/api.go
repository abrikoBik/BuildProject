package api

import (
	"fmt"
	"mp1/pkg/api/handlers"
	"mp1/pkg/api/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Init(hostname string, port uint) *http.Server {
	r := chi.NewRouter()
	r.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
			MaxAge:           86400,
		}),
		chiMiddleware.Heartbeat("/health"),
		chiMiddleware.RequestID,
		chiMiddleware.Logger,
		chiMiddleware.AllowContentEncoding("gzip", "deflate"),
		chiMiddleware.AllowContentType("application/json"),
		middleware.AddHeaders,
	)

	r.Post("/register", handlers.Register)
	r.With(middleware.JSONResponse).Post("/login", handlers.Login)

	// r.With(middleware.ValidateToken).Get("/users", handlers.Users)
	// r.With(middleware.ValidateToken).Post("/logout", handlers.Logout)

	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", hostname, port),
		Handler: r,
	}
}
