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
	getWorkerTag    = "api:GetWorker"
	createWorkerTag = "api:CreateWorker"
	patchWorkerTag  = "api:PatchWorker"
)

func GetWorker(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if errors.Is(user.Get(db.Database, r.Context().Value(middleware.UserID).(uint)), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		return
	}
	util.Debug(getWorkerTag, "user", user)
	if user.Worker == nil {
		http.Error(w, l10n.WorkerNotFound, http.StatusNotFound)
		return
	}
	var worker models.Worker
	// using util.Check as a temporary solution
	util.Check(worker.Get(db.Database, user.Worker.ID))
	util.Debug(getWorkerTag, "worker", worker)
	response := models.WorkerResponse{Worker: &worker, CreatedAt: worker.CreatedAt.String()}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error(getWorkerTag, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	util.Info(getWorkerTag, "id%d returned", worker.ID)
}

func CreateWorker(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetID(middleware.UserID, r.Context())
	if err != nil {
		return
	}
	var worker models.Worker
	request := models.WorkerRequest{
		Worker: &worker,
		UserID: userID,
	}
	if !jsonDecode(w, r, &request) {
		return
	}
	if err := worker.Create(db.Database, request); err != nil {
		util.Error(createWorkerTag, err)
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
			http.Error(w, l10n.NameOrEmailTaken, http.StatusBadRequest)
		default:
			http.Error(w, l10n.RegisterError, http.StatusInternalServerError)
		}
		return
	}
	util.Info(createWorkerTag, "id%d registered", worker.ID)
	response := models.WorkerResponse{ID: worker.ID}
	if err := jsonEncode(w, response, "register"); err != nil {
		util.Error(createWorkerTag, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PatchWorker(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetID(middleware.UserID, r.Context())
	if err != nil {
		return
	}
	var user models.User
	if user.Get(db.Database, userID) != nil {
		http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		return
	}
	util.Debug(patchWorkerTag, "user", user)
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
		util.Error(patchWorkerTag, err)
		http.Error(w, l10n.GenericWorker, http.StatusInternalServerError)
		return
	}
	util.Info(patchWorkerTag, "%s (id%d) patched", user.Name, worker.ID)
	response := models.WorkerResponse{Worker: &worker, CreatedAt: worker.CreatedAt.String()}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error(patchWorkerTag, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
