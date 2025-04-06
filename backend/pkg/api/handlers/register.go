package handlers

import (
	"mp1/pkg/db"
	"mp1/pkg/l10n"
	"mp1/pkg/models"
	"mp1/pkg/util"
	"net/http"

	"gorm.io/gorm"
)

const (
	registerUserTag   = "api:RegisterUser"
	registerWorkerTag = "api:RegisterWorker"
)

// Creates a new database entry for user.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var request models.UserRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	var user models.User
	if err := user.Create(db.Database, request); err != nil {
		util.Error(registerUserTag, err)
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
	util.Info(registerUserTag, "%s (id%d) registered", user.Name, user.ID)
	var response models.UserResponse
	response.ID = user.ID
	if err := jsonEncode(w, response, "register"); err != nil {
		util.Error(registerUserTag, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func RegisterWorker(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	var request models.WorkerRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	if len(request.Role) == 0 {
		http.Error(w, l10n.InvalidRole, http.StatusBadRequest)
		return
	}
	if len(request.Password) == 0 {
		http.Error(w, l10n.EmptyPassword, http.StatusBadRequest)
		return
	}
	worker := models.Worker{
		Role:   request.Role,
		UserID: userID,
	}
	if err := worker.Create(db.Database, request); err != nil {
		util.Error(registerWorkerTag, err)
		http.Error(w, l10n.UserRegisterError, http.StatusInternalServerError)
		return
	}
	util.Info(registerWorkerTag, "id%d registered", worker.ID)
	var response models.WorkerResponse
	response.ID = worker.ID
	if err := jsonEncode(w, response, "register"); err != nil {
		util.Error(registerWorkerTag, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
