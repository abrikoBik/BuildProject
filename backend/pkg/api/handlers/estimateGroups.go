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
	lGetEstimateGroup    = util.CreateLogger("api:GetEstimateGroup")
	lCreateEstimateGroup = util.CreateLogger("api:CreateEstimateGroup")
	lPatchEstimateGroup  = util.CreateLogger("api:PatchEstimateGroup")
	lDeleteEstimateGroup = util.CreateLogger("api:DeleteEstimateGroup")
)

func GetEstimateGroup(w http.ResponseWriter, r *http.Request) {
	projectID, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		lGetEstimateGroup.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if errors.Is(project.Get(db.Database, projectID, "EstimateGroups"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	if project.Get(db.Database, projectID, "EstimateGroups.Estimates") != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
}

func CreateEstimateGroup(w http.ResponseWriter, r *http.Request) {
	projectID, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		lCreateEstimateGroup.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var request models.EstimateGroupRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	var project models.Project
	if project.Get(db.Database, projectID, "EstimateGroups.Estimates") != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	request.ProjectID = project.ID
	var group models.EstimateGroup
	if group.Create(db.Database, request) != nil {
		lCreateEstimateGroup.Error(err)
		http.Error(w, l10n.GenericEstimateGroup, http.StatusInternalServerError)
		return
	}
	lCreateEstimateGroup.Info("%s (id%d) created", group.Name, group.ID)
	response := models.EstimateGroupResponse{ID: group.ID}
	if err := jsonEncode(w, response, "register"); err != nil {
		lCreateEstimateGroup.Error(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PatchEstimateGroup(w http.ResponseWriter, r *http.Request) {
	lPatchEstimateGroup.Info("not implemented")
	w.WriteHeader(http.StatusNotImplemented)
}

func DeleteEstimateGroup(w http.ResponseWriter, r *http.Request) {
	lDeleteEstimateGroup.Info("not implemented")
	w.WriteHeader(http.StatusNotImplemented)
}
