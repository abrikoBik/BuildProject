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

const getStageTag = "api:GetStage"

func GetStage(w http.ResponseWriter, r *http.Request) {
	var stage models.Stage
	result := db.Database.Where("id = ?", r.Context().Value(middleware.StageID)).Take(&stage)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, l10n.StageNotFound, http.StatusNotFound)
		return
	}
	response := models.StageResponse{Stage: &stage, CreatedAt: stage.CreatedAt.String()}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error(getStageTag, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	util.Info(getStageTag, "id%d returned", stage.ID)
}
