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
	response := models.WorkerResponse{Worker: &user.Worker, RegisteredAt: user.Worker.CreatedAt.String()}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error(getWorkerTag, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	util.Info(getWorkerTag, "id%d returned", user.Worker.ID)
}
