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
	lGetEstimate.Info("not implemented")
	w.WriteHeader(http.StatusNotImplemented)
}

func CreateEstimate(w http.ResponseWriter, r *http.Request) {
	lCreateEstimate.Info("not implemented")
	w.WriteHeader(http.StatusNotImplemented)
}

func PatchEstimate(w http.ResponseWriter, r *http.Request) {
	lPatchEstimate.Info("not implemented")
	w.WriteHeader(http.StatusNotImplemented)
}

func DeleteEstimate(w http.ResponseWriter, r *http.Request) {
	lDeleteEstimate.Info("not implemented")
	w.WriteHeader(http.StatusNotImplemented)
}
