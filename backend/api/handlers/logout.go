package handlers

import (
	"errors"
	"mp1/api/middleware"
	"mp1/db"
	"mp1/models"
	"mp1/util"
	"net/http"

	"gorm.io/gorm"
)

// Removes the database entry for a session.
// TODO: rewrite
func Logout(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value(middleware.JWT_CONTEXT_KEY)
	// sanity checks
	var dbSession models.Session
	result := db.Database.Where("token = ?", token).Take(&dbSession)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "token not found", http.StatusNotFound)
		return
	}

	result = db.Database.Where("token = ?", token).Delete(&dbSession)
	if result.Error != nil {
		util.Error("api:logout", result.Error)
		http.Error(w, "error invalidating session", http.StatusInternalServerError)
		return
	}
	util.Info("api:logout", "token %s removed", token)
	w.WriteHeader(http.StatusOK)
}
