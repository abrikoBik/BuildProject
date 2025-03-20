package api

import (
	"fmt"
	"mp1/api/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Init(hostname string, port uint) *http.Server {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           86400,
	}))
	r.Use(chiMiddleware.Heartbeat("/health"))
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.AllowContentEncoding("gzip", "deflate"))
	r.Use(chiMiddleware.AllowContentType("application/json"))

	r.Post("/register", register)
	r.Post("/login", login)

	r.With(middleware.ValidateToken).Get("/userInfo", userInfo)
	r.With(middleware.ValidateToken).Post("/logout", logout)

	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", hostname, port),
		Handler: r,
	}
}
