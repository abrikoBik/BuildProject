package handlers

import (
	"mp1/db"
	"mp1/models"
	"mp1/util"
	"net/http"
)

// Creates a new database entry for session.
func Login(w http.ResponseWriter, r *http.Request) {
	var request models.SessionRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	var session models.Session
	// TODO: check for wrong password
	if err := session.Create(db.Database, request); err != nil {
		util.Error("api:login", err)
		http.Error(w, "couldn't log in", http.StatusInternalServerError)
		return
	}
	response := models.SessionResponse{
		Session: &session,
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error("api:login", err)
		return
	}
	util.Info("api:login", "id%d logged in", session.UserID)
	w.WriteHeader(http.StatusCreated)
}
