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
	getStageTag    = "api:GetStage"
	createStageTag = "api:CreateStage"
	patchStageTag  = "api:PatchStage"
	deleteStageTag = "api:DeleteStage"
)

func GetStage(w http.ResponseWriter, r *http.Request) {
	// projectID, err := middleware.GetID(middleware.ProjectID, r.Context())
	// if err != nil {
	// 	util.Error(getStageTag, err)
	// 	http.Error(w, l10n.ContextID, http.StatusInternalServerError)
	// 	return
	// }
	// var project models.Project
	// if errors.Is(project.Get(db.Database, projectID), gorm.ErrRecordNotFound) {
	// 	http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
	// 	return
	// }

	// stageID, err := middleware.GetID(middleware.StageID, r.Context())
	// if err != nil {
	// 	util.Error(getStageTag, err)
	// 	http.Error(w, l10n.ContextID, http.StatusInternalServerError)
	// 	return
	// }
	// response := models.StageResponse{
	// 	StartTime: ,
	// }
	// if err := jsonEncode(w, response, "public"); err != nil {
	// 	util.Error(getStageTag, err)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// util.Info(getStageTag, "id%d returned", stage.ID)

	// var stage models.Stage
	// result := db.Database.Where("id = ?", r.Context().Value(middleware.StageID)).Take(&stage)
	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	http.Error(w, l10n.StageNotFound, http.StatusNotFound)
	// 	return
	// }
	// // response := models.StageResponse{Stage: &stage, CreatedAt: stage.CreatedAt.String()}
	// // if err := jsonEncode(w, response, "public"); err != nil {
	// // 	util.Error(getStageTag, err)
	// // 	return
	// // }
	// w.WriteHeader(http.StatusOK)
	// util.Info(getStageTag, "id%d returned", stage.ID)
	w.WriteHeader(http.StatusNotImplemented)
}

func CreateStage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
