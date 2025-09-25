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
	lGetStage    = util.CreateLogger("api:GetStage")
	lCreateStage = util.CreateLogger("api:CreateStage")
	lPatchStage  = util.CreateLogger("api:PatchStage")
	lDeleteStage = util.CreateLogger("api:DeleteStage")
)

func GetStage(w http.ResponseWriter, r *http.Request) {
	stageGroupID, err1 := middleware.GetID(middleware.StageGroupID, r.Context())
	stageID, err2 := middleware.GetID(middleware.StageID, r.Context())
	for _, e := range []error{err1, err2} {
		if e != nil {
			lGetStage.Error(e)
			http.Error(w, l10n.ContextID, http.StatusInternalServerError)
			return
		}
	}
	var stageGroup models.StageGroup
	if errors.Is(stageGroup.Get(db.Database, stageGroupID), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.StageGroupNotFound, http.StatusNotFound)
	}
	var stage models.Stage
	if errors.Is(stage.Get(db.Database, stageID), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.StageNotFound, http.StatusNotFound)
	}
	response := models.StageResponse{
		Stage: &stage,
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		lGetStageGroup.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	lGetStage.Info("id%d returned", stage.ID)
}

func CreateStage(w http.ResponseWriter, r *http.Request) {
	stageGroupID, err := middleware.GetID(middleware.StageGroupID, r.Context())
	if err != nil {
		lCreateStage.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var stageGroup models.StageGroup
	if errors.Is(stageGroup.Get(db.Database, stageGroupID), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.StageGroupNotFound, http.StatusNotFound)
	}
	var request models.StageRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	request.StageGroupID = stageGroupID
	lCreateStage.Debug("stageGroupID", stageGroupID, "stageGroup", stageGroup, "request", request)
	var stage models.Stage
	if err := stage.Create(db.Database, request); err != nil {
		lCreateStage.Error(err)
		switch err {
		default:
			http.Error(w, l10n.GenericStage, http.StatusInternalServerError)
		}
	}
	lCreateStage.Info("%s (id%d) created", stage.Name, stage.ID)
	response := models.StageResponse{ID: stage.ID}
	if err := jsonEncode(w, response, "register"); err != nil {
		lCreateStage.Error(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PatchStage(w http.ResponseWriter, r *http.Request) {
	stageGroupID, err1 := middleware.GetID(middleware.StageGroupID, r.Context())
	stageID, err2 := middleware.GetID(middleware.StageID, r.Context())
	for _, e := range []error{err1, err2} {
		if e != nil {
			lGetStage.Error(e)
			http.Error(w, l10n.ContextID, http.StatusInternalServerError)
			return
		}
	}
	var stageGroup models.StageGroup
	if errors.Is(stageGroup.Get(db.Database, stageGroupID, "Stages"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.StageGroupNotFound, http.StatusNotFound)
	}
	var stage models.Stage
	if stage.Get(db.Database, stageID) != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	var request models.StageRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	if err := stage.Patch(db.Database, request); err != nil {
		lPatchStage.Error(err)
		switch err {
		default:
			http.Error(w, l10n.GenericStage, http.StatusInternalServerError)
		}
		return
	}
	lPatchStage.Info("id%d patched", stage.ID)
	response := models.StageResponse{
		Stage: &stage,
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		lPatchStage.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteStage(w http.ResponseWriter, r *http.Request) {
	stageGroupID, err1 := middleware.GetID(middleware.StageGroupID, r.Context())
	stageID, err2 := middleware.GetID(middleware.StageID, r.Context())
	for _, e := range []error{err1, err2} {
		if e != nil {
			lDeleteStage.Error(e)
			http.Error(w, l10n.ContextID, http.StatusInternalServerError)
			return
		}
	}
	var stageGroup models.StageGroup
	if errors.Is(stageGroup.Get(db.Database, stageGroupID, "Stages"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.StageGroupNotFound, http.StatusNotFound)
	}
	var stage models.Stage
	if errors.Is(stage.Get(db.Database, stageID), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.StageNotFound, http.StatusNotFound)
		return
	}
	if err := stage.Delete(db.Database); err != nil {
		lDeleteStage.Error(err)
		return
	}
	lDeleteStage.Info("id%d deleted", stage.ID)
	w.WriteHeader(http.StatusOK)
}
