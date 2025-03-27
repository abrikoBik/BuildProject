package handlers

import (
	"errors"
	"mp1/pkg/db"
	"mp1/pkg/l10n"
	"mp1/pkg/models"
	"mp1/pkg/util"
	"net/http"
	"net/mail"
	"strings"

	"gorm.io/gorm"
)

// Creates a new database entry for user.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var request models.UserRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	if _, err := mail.ParseAddress(request.Email); err != nil {
		http.Error(w, l10n.InvalidEmail, http.StatusBadRequest)
		return
	}
	if strings.Compare(request.Password, "") == 0 {
		http.Error(w, l10n.EmptyPassword, http.StatusBadRequest)
		return
	}

	user := models.User{
		Login: request.Login,
		Name:  request.Name,
		Email: request.Email,
	}
	if err := user.Create(db.Database, request); err != nil {
		util.Error("api:RegisterUser", err)
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			http.Error(w, l10n.NameOrEmailTaken, http.StatusBadRequest)
		} else {
			http.Error(w, l10n.UserRegisterError, http.StatusInternalServerError)
		}
		return
	}

	response := models.UserResponse{
		ID: user.ID,
	}
	if err := jsonEncode(w, response, "register"); err != nil {
		util.Error("api:RegisterUser", err)
		return
	}
	util.Info("api:RegisterUser", "%s (id%d) registered", user.Name, user.ID)
	w.WriteHeader(http.StatusCreated)
}

func RegisterWorker(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	var request models.WorkerRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	if strings.Compare(request.Role, "") == 0 {
		http.Error(w, l10n.InvalidRole, http.StatusBadRequest)
		return
	}
	if strings.Compare(request.Password, "") == 0 {
		http.Error(w, l10n.EmptyPassword, http.StatusBadRequest)
		return
	}

	worker := models.Worker{
		Role:   request.Role,
		UserID: userID,
	}
	if err := worker.Create(db.Database, request); err != nil {
		util.Error("api:RegisterWorker", err)
		http.Error(w, l10n.UserRegisterError, http.StatusInternalServerError)
		return
	}

	response := models.WorkerResponse{
		ID: worker.ID,
	}
	if err := jsonEncode(w, response, "register"); err != nil {
		util.Error("api:RegisterWorker", err)
		return
	}
	util.Info("api:RegisterWorker", "id%d registered", worker.ID)
	w.WriteHeader(http.StatusCreated)
}
