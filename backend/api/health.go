package api

import (
	"mp1/util"
	"net/http"
)

// Sends back code 200 on request.
func getHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	util.Log("api", "/health 200")
}
