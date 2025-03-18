package api

import (
	"fmt"
	"mp1/api/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Init(port int) *http.Server {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.Recoverer)
	router.Use(chiMiddleware.Heartbeat("/health"))

	router.Post("/register", register)
	router.Post("/login", login)

	// router.With(middleware.ValidateToken).Get("/userInfo", userInfo)
	router.With(middleware.ValidateToken).Post("/logout", logout)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
}
