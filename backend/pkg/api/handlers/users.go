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

const getUserTag = "api:GetUser"

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	result := db.Database.Where("id = ?", r.Context().Value("userID")).Take(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		return
	}
	response := models.UserResponse{
		User:         &user,
		RegisteredAt: user.CreatedAt.String(),
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error(getUserTag, err)
		return
	}
	util.Info(getUserTag, "id%d returned", user.ID)
	w.WriteHeader(http.StatusOK)
}
