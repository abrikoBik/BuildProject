package handlers

import (
	"mp1/pkg/api/middleware"
	"mp1/pkg/db"
	"mp1/pkg/l10n"
	"mp1/pkg/models"
	"mp1/pkg/util"
	"net/http"

	"gorm.io/gorm"
)

const patchUserTag = "api:PatchUser"

// Updates the user database entry.
func PatchUser(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetID(middleware.UserID, r.Context())
	var user models.User
	if err := db.Database.Where("id = ?", id).Take(&user).Error; err != nil {
		http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		return
	}
	var request models.UserRequest
	request.User = new(models.User)
	if err != nil || !jsonDecode(w, r, &request) {
		return
	}
	util.Debug(patchUserTag, "request", request, "user", user)
	token, err := middleware.GetToken(r.Context())
	if err != nil {
		util.Info(patchUserTag, err.Error())
		http.Error(w, l10n.InvalidToken, http.StatusInternalServerError)
		return
	}
	verified, err := user.VerifyToken(db.Database, token)
	if err != nil {
		util.Info(patchUserTag, err.Error())
		http.Error(w, l10n.InvalidToken, http.StatusInternalServerError)
		return
	}
	if !verified {
		http.Error(w, l10n.InvalidToken, http.StatusForbidden)
		return
	}
	if err := user.Patch(db.Database, request); err != nil {
		util.Error(patchUserTag, err)
		switch err {
		case util.InsecurePasswordError:
			http.Error(w, l10n.InsecurePassword, http.StatusBadRequest)
		case util.InvalidEmailError:
			http.Error(w, l10n.InvalidEmail, http.StatusBadRequest)
		case gorm.ErrDuplicatedKey:
			http.Error(w, l10n.NameOrEmailTaken, http.StatusBadRequest)
		default:
			http.Error(w, l10n.UserRegisterError, http.StatusInternalServerError)
		}
		return
	}
	util.Info(patchUserTag, "%s (id%d) patched", user.Name, user.ID)
	response := models.UserResponse{
		User:         &user,
		RegisteredAt: user.CreatedAt.String(),
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error(patchUserTag, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
