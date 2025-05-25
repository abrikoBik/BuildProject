package api

import (
	"fmt"
	"mp1/pkg/api/middleware"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var jwtKey []byte

func init() {
	var (
		keyString string
		exists    bool
	)
	if keyString, exists = os.LookupEnv("JWT_KEY"); !exists {
		panic("JWT_KEY not found in environment")
	}
	jwtKey = []byte(keyString)
}

func Init(hostname string, port uint) *http.Server {
	r := chi.NewRouter()
	r.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
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

	addSessions(r)
	addUsers(r)
	addProjects(r)

	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", hostname, port),
		Handler: r,
	}
}
