package handlers

import (
	// "errors"
	// "mp1/pkg/api/middleware"
	// "mp1/pkg/db"
	// "mp1/pkg/l10n"
	// "mp1/pkg/models"
	// "mp1/pkg/util"
	"net/http"
	// "gorm.io/gorm"
)

// const logoutTag = "api:Logout"

// Removes the database entry for a session.
// TODO: rewrite
func Logout(w http.ResponseWriter, r *http.Request) {
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
