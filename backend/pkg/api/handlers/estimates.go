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
	lGetEstimate    = util.CreateLogger("api:GetEstimate")
	lCreateEstimate = util.CreateLogger("api:CreateEstimate")
	lPatchEstimate  = util.CreateLogger("api:PatchEstimate")
	lDeleteEstimate = util.CreateLogger("api:DeleteEstimate")
)

func GetEstimate(w http.ResponseWriter, r *http.Request) {
	estimateGroupID, err1 := middleware.GetID(middleware.EstimateGroupID, r.Context())
	estimateID, err2 := middleware.GetID(middleware.EstimateID, r.Context())
	for _, e := range []error{err1, err2} {
		if e != nil {
			lGetEstimate.Error(e)
			http.Error(w, l10n.ContextID, http.StatusInternalServerError)
			return
		}
	}
	var estimateGroup models.EstimateGroup
	if errors.Is(estimateGroup.Get(db.Database, estimateGroupID), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.EstimateGroupNotFound, http.StatusNotFound)
	}
	var estimate models.Estimate
	if errors.Is(estimate.Get(db.Database, estimateID), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.EstimateNotFound, http.StatusNotFound)
	}
	response := models.EstimateResponse{
		Estimate: &estimate,
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		lGetEstimateGroup.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	lGetEstimate.Info("id%d returned", estimate.ID)
}

func CreateEstimate(w http.ResponseWriter, r *http.Request) {
	estimateGroupID, err := middleware.GetID(middleware.EstimateGroupID, r.Context())
	if err != nil {
		lCreateEstimate.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var estimateGroup models.EstimateGroup
	if errors.Is(estimateGroup.Get(db.Database, estimateGroupID), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.EstimateGroupNotFound, http.StatusNotFound)
	}
	var request models.EstimateRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	request.EstimateGroupID = estimateGroupID
	lCreateEstimate.Debug("estimateGroupID", estimateGroupID, "estimateGroup", estimateGroup, "request", request)
	var estimate models.Estimate
	if err := estimate.Create(db.Database, request); err != nil {
		lCreateEstimate.Error(err)
		switch err {
		default:
			http.Error(w, l10n.GenericEstimate, http.StatusInternalServerError)
		}
	}
	lCreateEstimate.Info("%s (id%d) created", estimate.Name, estimate.ID)
	response := models.EstimateResponse{ID: estimate.ID}
	if err := jsonEncode(w, response, "register"); err != nil {
		lCreateEstimate.Error(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PatchEstimate(w http.ResponseWriter, r *http.Request) {
	estimateGroupID, err1 := middleware.GetID(middleware.EstimateGroupID, r.Context())
	estimateID, err2 := middleware.GetID(middleware.EstimateID, r.Context())
	for _, e := range []error{err1, err2} {
		if e != nil {
			lGetEstimate.Error(e)
			http.Error(w, l10n.ContextID, http.StatusInternalServerError)
			return
		}
	}
	var estimateGroup models.EstimateGroup
	if errors.Is(estimateGroup.Get(db.Database, estimateGroupID, "Estimates"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.EstimateGroupNotFound, http.StatusNotFound)
	}
	var estimate models.Estimate
	if estimate.Get(db.Database, estimateID) != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	var request models.EstimateRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	if err := estimate.Patch(db.Database, request); err != nil {
		lPatchEstimate.Error(err)
		switch err {
		default:
			http.Error(w, l10n.GenericEstimate, http.StatusInternalServerError)
		}
		return
	}
	lPatchEstimate.Info("id%d patched", estimate.ID)
	// TODO: fix recalculation (only recalculates after 2 PATCH requests???)
	if err := estimateGroup.CalculateCost(db.Database); err != nil {
		lPatchEstimate.Error(err)
		http.Error(w, l10n.GenericEstimateGroup, http.StatusInternalServerError)
		return
	}
	lDeleteAttachment.Info("id%d cost recalculated", estimateGroup.ID)
	response := models.EstimateResponse{
		Estimate: &estimate,
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		lPatchEstimate.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteEstimate(w http.ResponseWriter, r *http.Request) {
	estimateGroupID, err1 := middleware.GetID(middleware.EstimateGroupID, r.Context())
	estimateID, err2 := middleware.GetID(middleware.EstimateID, r.Context())
	for _, e := range []error{err1, err2} {
		if e != nil {
			lDeleteEstimate.Error(e)
			http.Error(w, l10n.ContextID, http.StatusInternalServerError)
			return
		}
	}
	var estimateGroup models.EstimateGroup
	if errors.Is(estimateGroup.Get(db.Database, estimateGroupID, "Estimates"), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.EstimateGroupNotFound, http.StatusNotFound)
	}
	var estimate models.Estimate
	if errors.Is(estimate.Get(db.Database, estimateID), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.EstimateNotFound, http.StatusNotFound)
		return
	}
	if err := estimate.Delete(db.Database); err != nil {
		lDeleteAttachment.Error(err)
		return
	}
	lDeleteAttachment.Info("id%d deleted", estimate.ID)
	if err := estimateGroup.CalculateCost(db.Database); err != nil {
		lPatchEstimate.Error(err)
		http.Error(w, l10n.GenericEstimateGroup, http.StatusInternalServerError)
		return
	}
	lDeleteAttachment.Info("id%d cost recalculated", estimateGroup.ID)
	w.WriteHeader(http.StatusOK)
}
