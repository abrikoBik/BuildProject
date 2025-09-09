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

const (
	createSessionTag = "api:CreateSession"
	// deleteSessionTag = "api:DeleteSession"
)

func CreateSession(w http.ResponseWriter, r *http.Request) {
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
		util.Error(createSessionTag, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		} else if errors.Is(err, l10n.WrongPasswordError) {
			http.Error(w, l10n.WrongPassword, http.StatusForbidden)
		} else {
			http.Error(w, l10n.Login, http.StatusInternalServerError)
		}
		return
	}
	util.Info(createSessionTag, "id%d logged in", session.UserID)
	response := models.SessionResponse{Session: &session}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error(createSessionTag, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// TODO: rewrite
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	// token := r.Context().Value(middleware.JWT_CONTEXT_KEY)
	// // sanity checks
	// var dbSession models.Session
	// result := db.Database.Where("token = ?", token).Take(&dbSession)
	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	http.Error(w, l10n.TokenNotFound, http.StatusNotFound)
	// 	return
	// }

	// result = db.Database.Where("token = ?", token).Delete(&dbSession)
	// if result.Error != nil {
	// 	util.Error(logoutTag, result.Error)
	// 	http.Error(w, l10n.LogoutError, http.StatusInternalServerError)
	// 	return
	// }
	// util.Info(logoutTag, "token %s removed", token)
	// w.WriteHeader(http.StatusOK)
}
