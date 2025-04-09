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
	userID, err := middleware.GetID(middleware.UserID, r.Context())
	if err != nil {
		return
	}
	var request models.WorkerRequest
	request.Worker = new(models.Worker)
	println("trying to print userID")
	request.Worker.UserID = userID
	println(request.UserID)
	if !jsonDecode(w, r, &request) {
		return
	}
	var worker models.Worker
	if err := worker.Create(db.Database, request); err != nil {
		util.Error(registerWorkerTag, err)
		switch err {
		case util.InvalidRoleError:
			http.Error(w, l10n.InvalidRole, http.StatusBadRequest)
		case util.InsecurePasswordError:
			http.Error(w, l10n.InsecurePassword, http.StatusBadRequest)
		case gorm.ErrDuplicatedKey:
			http.Error(w, l10n.NameOrEmailTaken, http.StatusBadRequest)
		default:
			http.Error(w, l10n.UserRegisterError, http.StatusInternalServerError)
		}
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
