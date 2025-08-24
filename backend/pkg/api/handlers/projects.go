package handlers

import "net/http"

const (
	getProjectTag    = "api:GetProject"
	createProjectTag = "api:CreateProject"
	patchProjectTag  = "api:PatchProject"
)

func GetProject(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func PatchProject(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
