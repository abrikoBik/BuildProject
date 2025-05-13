package handlers

import (
	"errors"
	"mp1/pkg/db"
	"mp1/pkg/l10n"
	"mp1/pkg/models"
	"mp1/pkg/util"
	"net/http"

	"gorm.io/gorm"
)

const loginTag = "api:Login"

// Creates a new database entry for session.
func Login(w http.ResponseWriter, r *http.Request) {
	var request models.SessionRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	if len(request.Password) == 0 {
		http.Error(w, l10n.EmptyPassword, http.StatusBadRequest)
		return
	}
	var session models.Session
	if err := session.Create(db.Database, request, jwtKey); err != nil {
		util.Error(loginTag, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		} else if errors.Is(err, util.WrongPasswordError) {
			http.Error(w, l10n.WrongPassword, http.StatusForbidden)
		} else {
			http.Error(w, l10n.LoginError, http.StatusInternalServerError)
		}
		return
	}
	util.Info(loginTag, "id%d logged in", session.UserID)
	response := models.SessionResponse{Session: &session}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error(loginTag, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
