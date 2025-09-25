package handlers

import (
	"errors"
	"mp1/pkg/api/middleware"
	"mp1/pkg/db"
	"mp1/pkg/l10n"
	"mp1/pkg/models"
	"mp1/pkg/util"
	"net/http"
	"sort"

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
	stageGroupID, err2 := middleware.GetID(middleware.StageGroupID, r.Context())
	// TODO: errors.Join?
	for _, e := range []error{err1, err2} {
		if e != nil {
			lGetStageGroup.Error(e)
			http.Error(w, l10n.ContextID, http.StatusInternalServerError)
			return
		}
	}
	var project models.Project
	// TODO: is this preload really needed?
	if errors.Is(project.Get(db.Database, projectID, "StageGroups"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	var stageGroup models.StageGroup
	if errors.Is(stageGroup.Get(db.Database, stageGroupID, "Stages"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.StageGroupNotFound, http.StatusNotFound)
		return
	}
	response := models.StageGroupResponse{
		StageGroup: &stageGroup,
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		lGetStageGroup.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	lGetStageGroup.Info("id%d returned", project.ID)
}

func CreateStageGroup(w http.ResponseWriter, r *http.Request) {
	projectID, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		lCreateStageGroup.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var request models.StageGroupRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	var project models.Project
	if project.Get(db.Database, projectID, "StageGroups.Stages") != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	request.ProjectID = project.ID
	var group models.StageGroup
	if group.Create(db.Database, request) != nil {
		lCreateStageGroup.Error(err)
		http.Error(w, l10n.GenericStageGroup, http.StatusInternalServerError)
		return
	}
	lCreateStageGroup.Info("%s (id%d) created", group.Name, group.ID)
	response := models.StageGroupResponse{ID: group.ID}
	if err := jsonEncode(w, response, "register"); err != nil {
		lCreateStageGroup.Error(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PatchStageGroup(w http.ResponseWriter, r *http.Request) {
	stageGroupID, err := middleware.GetID(middleware.StageGroupID, r.Context())
	if err != nil {
		lPatchStageGroup.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var stageGroup models.StageGroup
	if errors.Is(stageGroup.Get(db.Database, stageGroupID, "Stages"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.StageGroupNotFound, http.StatusNotFound)
	}
	var request models.StageGroupRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	if err := stageGroup.Patch(db.Database, request); err != nil {
		lPatchStageGroup.Error(err)
		switch err {
		default:
			http.Error(w, l10n.GenericStageGroup, http.StatusInternalServerError)
		}
		return
	}
	lPatchStageGroup.Info("id%d patched", stageGroup.ID)
	response := models.StageGroupResponse{
		StageGroup: &stageGroup,
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		lPatchStageGroup.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteStageGroup(w http.ResponseWriter, r *http.Request) {
	stageGroupID, err := middleware.GetID(middleware.StageGroupID, r.Context())
	if err != nil {
		lDeleteStageGroup.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var stageGroup models.StageGroup
	if errors.Is(stageGroup.Get(db.Database, stageGroupID, "Stages"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.StageGroupNotFound, http.StatusNotFound)
	}
	if err := stageGroup.Delete(db.Database); err != nil {
		lDeleteStageGroup.Error(err)
		return
	}
	lDeleteStageGroup.Info("id%d deleted", stageGroup.ID)
	w.WriteHeader(http.StatusOK)
}

func ListStages(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetID(middleware.StageGroupID, r.Context())
	if err != nil {
		lListStages.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var stageGroup models.StageGroup
	if stageGroup.Get(db.Database, id, "Stages") != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	sort.Slice(stageGroup.Stages, func(i, j int) bool {
		return stageGroup.Stages[i].UpdatedAt.UnixMilli() > stageGroup.Stages[j].UpdatedAt.UnixMilli()
	})
	var response []models.StageResponse
	for _, e := range stageGroup.Stages {
		toAppend := models.StageResponse{
			Stage: &e,
			ID:    e.ID,
		}
		response = append(response, toAppend)
	}
	if err := jsonEncode(w, response, "list"); err != nil {
		lListStages.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
