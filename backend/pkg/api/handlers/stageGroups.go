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
	lGetStageGroup    = util.CreateLogger("api:GetStageGroup")
	lCreateStageGroup = util.CreateLogger("api:CreateStageGroup")
	lPatchStageGroup  = util.CreateLogger("api:PatchStageGroup")
	lDeleteStageGroup = util.CreateLogger("api:DeleteStageGroup")
	lListStages       = util.CreateLogger("api:ListStages")
)

func GetStageGroup(w http.ResponseWriter, r *http.Request) {

	projectID, err1 := middleware.GetID(middleware.ProjectID, r.Context())
	estimateGroupID, err2 := middleware.GetID(middleware.EstimateGroupID, r.Context())
	for _, e := range []error{err1, err2} {
		if e != nil {
			lGetStageGroup.Error(e)
			http.Error(w, l10n.ContextID, http.StatusInternalServerError)
			return
		}
	}
	var project models.Project
	if errors.Is(project.Get(db.Database, projectID, "StageGroups"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	var estimateGroup models.StageGroup
	if errors.Is(estimateGroup.Get(db.Database, estimateGroupID, "Stages"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.StageGroupNotFound, http.StatusNotFound)
		return
	}
	response := models.StageGroupResponse{
		StageGroup: &estimateGroup,
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		lGetStageGroup.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	lGetStageGroup.Info("id%d returned", project.ID)
}

func CreateStageGroup(w http.ResponseWriter, r *http.Request) {
	lCreateStageGroup.Info("not implemented")
	w.WriteHeader(http.StatusNotImplemented)
}

func PatchStageGroup(w http.ResponseWriter, r *http.Request) {
	lPatchStageGroup.Info("not implemented")
	w.WriteHeader(http.StatusNotImplemented)
}

func DeleteStageGroup(w http.ResponseWriter, r *http.Request) {
	lDeleteStageGroup.Info("not implemented")
	w.WriteHeader(http.StatusNotImplemented)
}

func ListStages(w http.ResponseWriter, r *http.Request) {

}
