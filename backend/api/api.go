package api

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

// Global database connection.
var Database *gorm.DB

func Init(port int) *http.Server {
	server := &http.Server{Addr: fmt.Sprintf(":%d", port)}
	http.HandleFunc("/health", getHealth)
	go server.ListenAndServe()
	return server
}
