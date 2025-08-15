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

const getWorkerTag = "api:GetWorker"

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
	response := models.WorkerResponse{Worker: &worker, RegisteredAt: worker.CreatedAt.String()}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error(getWorkerTag, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	util.Info(getWorkerTag, "id%d returned", worker.ID)
}
