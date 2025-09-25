package handlers

import (
	"mp1/pkg/util"
	"net/http"
)

var (
	lGetStage    = util.CreateLogger("api:GetStage")
	lCreateStage = util.CreateLogger("api:CreateStage")
	lPatchStage  = util.CreateLogger("api:PatchStage")
	lDeleteStage = util.CreateLogger("api:DeleteStage")
)

func GetStage(w http.ResponseWriter, r *http.Request) {
	lGetStage.Info("not implemented")
	w.WriteHeader(http.StatusNotImplemented)
}

func CreateStage(w http.ResponseWriter, r *http.Request) {
	lCreateStage.Info("not implemented")
	w.WriteHeader(http.StatusNotImplemented)
}

func PatchStage(w http.ResponseWriter, r *http.Request) {
	lPatchStage.Info("not implemented")
	w.WriteHeader(http.StatusNotImplemented)
}

func DeleteStage(w http.ResponseWriter, r *http.Request) {
	lDeleteStage.Info("not implemented")
	w.WriteHeader(http.StatusNotImplemented)
}
