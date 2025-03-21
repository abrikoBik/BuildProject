package handlers

import (
	"errors"
	"mp1/db"
	"mp1/models"
	"mp1/util"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

// Creates a new database entry for session.
func Login(w http.ResponseWriter, r *http.Request) {
	var request models.SessionRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	if strings.Compare(request.Name, "") == 0 {
		http.Error(w, "empty password", http.StatusBadRequest)
		return
	}

	var session models.Session
	if err := session.Create(db.Database, request); err != nil {
		util.Error("api:login", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "user not found", http.StatusNotFound)
		} else if errors.Is(err, util.WrongPasswordError) {
			http.Error(w, util.WrongPasswordError.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, "couldn't create user", http.StatusInternalServerError)
		}
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
