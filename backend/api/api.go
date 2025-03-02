package api

import (
	"fmt"
	"mp1/api/middleware"
	"mp1/util"
	"net/http"

	"gorm.io/gorm"
)

// Global database connection.
var Database *gorm.DB

func Init(port int) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", health)
	mux.HandleFunc("POST /register", register)
	mux.HandleFunc("POST /login", login)

	logoutHandler := http.HandlerFunc(logout)
	mux.Handle("DELETE /logout", middleware.ValidateToken(logoutHandler))
	userInfoHandler := http.HandlerFunc(userInfo)
	mux.Handle("GET /userInfo", middleware.ValidateToken(userInfoHandler))

	util.Log("api", "starting server")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	return err
}
