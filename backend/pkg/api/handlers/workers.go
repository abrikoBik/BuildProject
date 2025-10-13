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

var (
	lGetWorker    = util.CreateLogger("api:GetWorker")
	lCreateWorker = util.CreateLogger("api:CreateWorker")
	lPatchWorker  = util.CreateLogger("api:PatchWorker")
)

func GetWorker(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetID(middleware.UserID, r.Context())
	if err != nil {
		lGetWorker.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var user models.User
	if errors.Is(user.Get(db.Database, userID, "Worker"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		return
	}
	if user.Worker == nil {
		http.Error(w, l10n.WorkerNotFound, http.StatusNotFound)
		return
	}
	var worker models.Worker
	if worker.Get(db.Database, user.Worker.ID) != nil {
		lGetWorker.Error(err)
		http.Error(w, l10n.GenericWorker, http.StatusInternalServerError)
		return
	}
	response := models.WorkerResponse{Worker: &worker}
	if err := jsonEncode(w, response, "public"); err != nil {
		lGetWorker.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	lGetWorker.Info("id%d returned", worker.ID)
}

func CreateWorker(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetID(middleware.UserID, r.Context())
	if err != nil {
		lCreateWorker.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var user models.User
	if errors.Is(user.Get(db.Database, userID, "Worker"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		return
	}
	var worker models.Worker
	request := models.WorkerRequest{Worker: &worker}
	request.UserID = userID
	if !jsonDecode(w, r, &request) {
		return
	}
	if err := worker.Create(db.Database, request); err != nil {
		lCreateWorker.Error(err)
		switch err {
		case l10n.WorkerExistsError:
			http.Error(w, l10n.WorkerExists, http.StatusBadRequest)
		case l10n.InvalidRoleError:
			http.Error(w, l10n.InvalidRole, http.StatusBadRequest)
		case l10n.InsecurePasswordError:
			http.Error(w, l10n.InsecurePassword, http.StatusBadRequest)
		case l10n.WrongPasswordError:
			http.Error(w, l10n.WrongPassword, http.StatusForbidden)
		case gorm.ErrDuplicatedKey:
			http.Error(w, l10n.FieldTaken, http.StatusBadRequest)
		default:
			http.Error(w, l10n.RegisterError, http.StatusInternalServerError)
		}
		return
	}
	user.Worker = &worker
	if err := db.Database.Save(&user).Error; err != nil {
		lCreateWorker.Error(err)
		http.Error(w, l10n.GenericUser, http.StatusInternalServerError)
		return
	}
	lCreateWorker.Info("id%d registered", worker.ID)
	response := models.WorkerResponse{ID: worker.ID}
	if err := jsonEncode(w, response, "register"); err != nil {
		lCreateWorker.Error(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PatchWorker(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetID(middleware.UserID, r.Context())
	if err != nil {
		lPatchWorker.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var user models.User
	if user.Get(db.Database, userID, "Worker") != nil {
		http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		return
	}
	if !verifyUser(user, w, r) {
		return
	}
	var worker models.Worker
	if db.Database.Where("user_id = ?", userID).Take(&worker).Error != nil {
		http.Error(w, l10n.WorkerNotFound, http.StatusNotFound)
		return
	}
	var request models.WorkerRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	if err := worker.Patch(db.Database, request); err != nil {
		lPatchWorker.Error(err)
		http.Error(w, l10n.GenericWorker, http.StatusInternalServerError)
		return
	}
	lPatchWorker.Info("%s (id%d) patched", user.Name, worker.ID)
	response := models.WorkerResponse{Worker: &worker}
	if err := jsonEncode(w, response, "public"); err != nil {
		lPatchWorker.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
