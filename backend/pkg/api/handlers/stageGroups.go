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
	if err := errors.Join(err1, err2); err != nil {
		lGetStageGroup.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if errors.Is(project.Get(db.Database, projectID), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	var stageGroup models.StageGroup
	if errors.Is(stageGroup.Get(db.Database, stageGroupID, "Stages"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.StageGroupNotFound, http.StatusNotFound)
		return
	}
	start, end := stageGroup.CalculateTimes(db.Database)
	lGetStageGroup.Debug("stageGroup", stageGroup, "start", start, "end", end)
	response := models.StageGroupResponse{
		StageGroup: &stageGroup,
		StartTime:  start.Unix(),
		EndTime:    end.Unix(),
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
	request.ProjectID = projectID
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
	if errors.Is(stageGroup.Get(db.Database, stageGroupID), gorm.ErrRecordNotFound) {
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
	projectID, err1 := middleware.GetID(middleware.ProjectID, r.Context())
	stageGroupID, err2 := middleware.GetID(middleware.StageGroupID, r.Context())
	if err := errors.Join(err1, err2); err != nil {
		lListStages.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if errors.Is(project.Get(db.Database, projectID), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	var stageGroup models.StageGroup
	if stageGroup.Get(db.Database, stageGroupID, "Stages") != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	sort.Slice(stageGroup.Stages, func(i, j int) bool {
		return stageGroup.Stages[i].UpdatedAt.UnixMilli() > stageGroup.Stages[j].UpdatedAt.UnixMilli()
	})
	var response []models.StageResponse
	for _, s := range stageGroup.Stages {
		toAppend := models.StageResponse{
			Stage:     &s,
			ID:        s.ID,
			StartTime: s.StartTime.Time.Unix(),
			EndTime:   s.EndTime.Time.Unix(),
		}
		response = append(response, toAppend)
	}
	if err := jsonEncode(w, response, "list"); err != nil {
		lListStages.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	lListStages.Info("%d stage(s) returned", len(stageGroup.Stages))
}
