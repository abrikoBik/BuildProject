package api

import (
	"fmt"
	"mp1/pkg/api/middleware"
	"mp1/pkg/util"
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
	util.Check(os.Mkdir(util.StoragePath, 0644))
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
		chiMiddleware.AllowContentType("application/json", "multipart/form-data"),
		middleware.AddHeaders,
	)

	addSessions(r)
	addUsers(r)
	addProjects(r)
	addAttachments(r)

	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", hostname, port),
		Handler: r,
	}
}
