package api

import "net/http"

// Sends back code 200 on request.
func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
