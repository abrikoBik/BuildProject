package handlers

import (
	"mp1/pkg/util"
	"net/http"
)

var (
	lGetStageGroup    = util.CreateLogger("api:GetStageGroup")
	lCreateStageGroup = util.CreateLogger("api:CreateStageGroup")
	lPatchStageGroup  = util.CreateLogger("api:PatchStageGroup")
	lDeleteStageGroup = util.CreateLogger("api:DeleteStageGroup")
)

func GetStageGroup(w http.ResponseWriter, r *http.Request) {
	lGetStageGroup.Info("not implemented")
	w.WriteHeader(http.StatusNotImplemented)
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
