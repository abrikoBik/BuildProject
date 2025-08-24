package handlers

import (
	"errors"
	"mp1/pkg/api/middleware"
	"mp1/pkg/db"
	"mp1/pkg/l10n"
	"mp1/pkg/models"
	"mp1/pkg/util"
	"net/http"

	"gorm.io/gorm"
)

const (
	getUserTag    = "api:GetUser"
	createUserTag = "api:CreateUser"
	patchUserTag  = "api:PatchUser"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	result := db.Database.Where("id = ?", r.Context().Value(middleware.UserID)).Take(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		return
	}
	response := models.UserResponse{User: &user, RegisteredAt: user.CreatedAt.String()}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error(getUserTag, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	util.Info(getUserTag, "id%d returned", user.ID)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var request models.UserRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	var user models.User
	if err := user.Create(db.Database, request); err != nil {
		util.Error(createUserTag, err)
		switch err {
		case l10n.InsecurePasswordError:
			http.Error(w, l10n.InsecurePassword, http.StatusBadRequest)
		case l10n.InvalidEmailError:
			http.Error(w, l10n.InvalidEmail, http.StatusBadRequest)
		case gorm.ErrDuplicatedKey:
			http.Error(w, l10n.NameOrEmailTaken, http.StatusBadRequest)
		default:
			http.Error(w, l10n.RegisterError, http.StatusInternalServerError)
		}
		return
	}
	util.Info(createUserTag, "%s (id%d) registered", user.Name, user.ID)
	response := models.UserResponse{ID: user.ID}
	if err := jsonEncode(w, response, "register"); err != nil {
		util.Error(createUserTag, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetID(middleware.UserID, r.Context())
	var user models.User
	if user.Get(db.Database, id) != nil {
		http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		return
	}
	if err != nil {
		return
	}
	if !verifyUser(user, w, r) {
		return
	}
	request := models.UserRequest{User: &models.User{}}
	if !jsonDecode(w, r, &request) {
		return
	}
	if err := user.Patch(db.Database, request); err != nil {
		util.Error(patchUserTag, err)
		switch err {
		case l10n.InsecurePasswordError:
			http.Error(w, l10n.InsecurePassword, http.StatusBadRequest)
		case l10n.InvalidEmailError:
			http.Error(w, l10n.InvalidEmail, http.StatusBadRequest)
		case gorm.ErrDuplicatedKey:
			http.Error(w, l10n.NameOrEmailTaken, http.StatusBadRequest)
		default:
			http.Error(w, l10n.GenericUser, http.StatusInternalServerError)
		}
		return
	}
	util.Info(patchUserTag, "%s (id%d) patched", user.Name, user.ID)
	response := models.UserResponse{User: &user, RegisteredAt: user.CreatedAt.String()}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error(patchUserTag, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
