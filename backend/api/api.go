package api

import (
	"fmt"
	// "mp1/api/middleware"
	"mp1/util"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func Init(port int) error {
	router := chi.NewRouter()
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.Recoverer)
	router.Use(chiMiddleware.Heartbeat("/health"))

	router.Post("/register", register)
	router.Post("/login", login)

	// router.With(middleware.ValidateToken).Get("/userInfo", userInfo)
	// router.With(middleware.ValidateToken).Post("/logout", logout)

	util.Info("api", "starting server")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	return err
}
