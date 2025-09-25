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
	lGetEstimateGroup    = util.CreateLogger("api:GetEstimateGroup")
	lCreateEstimateGroup = util.CreateLogger("api:CreateEstimateGroup")
	lPatchEstimateGroup  = util.CreateLogger("api:PatchEstimateGroup")
	lDeleteEstimateGroup = util.CreateLogger("api:DeleteEstimateGroup")
	lListEstimates       = util.CreateLogger("api:ListEstimates")
)

func GetEstimateGroup(w http.ResponseWriter, r *http.Request) {
	projectID, err1 := middleware.GetID(middleware.ProjectID, r.Context())
	estimateGroupID, err2 := middleware.GetID(middleware.EstimateGroupID, r.Context())
	for _, e := range []error{err1, err2} {
		if e != nil {
			lGetEstimateGroup.Error(e)
			http.Error(w, l10n.ContextID, http.StatusInternalServerError)
			return
		}
	}
	var project models.Project
	if errors.Is(project.Get(db.Database, projectID, "EstimateGroups"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	var estimateGroup models.EstimateGroup
	if errors.Is(estimateGroup.Get(db.Database, estimateGroupID, "Estimates"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.EstimateGroupNotFound, http.StatusNotFound)
		return
	}
	response := models.EstimateGroupResponse{
		EstimateGroup: &estimateGroup,
		TotalCost:     estimateGroup.CalculateCost(),
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		lGetEstimateGroup.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	lGetEstimateGroup.Info("id%d returned", project.ID)
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
	estimateGroupID, err := middleware.GetID(middleware.EstimateGroupID, r.Context())
	if err != nil {
		lPatchEstimateGroup.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var estimateGroup models.EstimateGroup
	if errors.Is(estimateGroup.Get(db.Database, estimateGroupID, "Estimates"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.EstimateGroupNotFound, http.StatusNotFound)
	}
	var request models.EstimateGroupRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	if err := estimateGroup.Patch(db.Database, request); err != nil {
		lPatchEstimateGroup.Error(err)
		switch err {
		default:
			http.Error(w, l10n.GenericEstimateGroup, http.StatusInternalServerError)
		}
		return
	}
	lPatchEstimateGroup.Info("id%d patched", estimateGroup.ID)
	response := models.EstimateGroupResponse{
		EstimateGroup: &estimateGroup,
		TotalCost:     estimateGroup.CalculateCost(),
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		lPatchEstimateGroup.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteEstimateGroup(w http.ResponseWriter, r *http.Request) {
	estimateGroupID, err := middleware.GetID(middleware.EstimateGroupID, r.Context())
	if err != nil {
		lDeleteEstimateGroup.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var estimateGroup models.EstimateGroup
	if errors.Is(estimateGroup.Get(db.Database, estimateGroupID, "Estimates"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.EstimateGroupNotFound, http.StatusNotFound)
	}
	if err := estimateGroup.Delete(db.Database); err != nil {
		lDeleteEstimateGroup.Error(err)
		return
	}
	lDeleteEstimateGroup.Info("id%d deleted", estimateGroup.ID)
	w.WriteHeader(http.StatusOK)
}

func ListEstimates(w http.ResponseWriter, r *http.Request) {
	// TODO: also check if project exists? (might not be needed)
	id, err := middleware.GetID(middleware.EstimateGroupID, r.Context())
	if err != nil {
		lListEstimates.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var estimateGroup models.EstimateGroup
	if estimateGroup.Get(db.Database, id, "Estimates") != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	sort.Slice(estimateGroup.Estimates, func(i, j int) bool {
		return estimateGroup.Estimates[i].UpdatedAt.UnixMilli() > estimateGroup.Estimates[j].UpdatedAt.UnixMilli()
	})
	var response []models.EstimateResponse
	for _, e := range estimateGroup.Estimates {
		toAppend := models.EstimateResponse{
			Estimate: &e,
			ID:       e.ID,
		}
		response = append(response, toAppend)
	}
	if err := jsonEncode(w, response, "list"); err != nil {
		lListEstimates.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
