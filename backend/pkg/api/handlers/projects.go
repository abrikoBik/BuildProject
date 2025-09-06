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

const (
	getProjectTag    = "api:GetProject"
	createProjectTag = "api:CreateProject"
	patchProjectTag  = "api:PatchProject"
)

func GetProject(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		util.Error(getProjectTag, err)
		http.Error(w, l10n.GenericProject, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if errors.Is(project.Get(db.Database, id), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
	}
	response := models.ProjectResponse{
		Project:   &project,
		StartTime: util.MarshalTime(project.StartTime),
		EndTime:   util.MarshalTime(project.EndTime),
		CreatedAt: project.CreatedAt.String(),
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error(getProjectTag, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	util.Info(getProjectTag, "id%d returned", project.ID)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	request := models.ProjectRequest{Project: &project}
	if !jsonDecode(w, r, &request) {
		return
	}
	if err := project.Create(db.Database, request); err != nil {
		util.Error(createProjectTag, err)
		switch err {
		default:
			http.Error(w, l10n.GenericProject, http.StatusInternalServerError)
		}
	}
	util.Info(createProjectTag, "%s (id%d) created", project.Name, project.ID)
	response := models.ProjectResponse{ID: project.ID}
	if err := jsonEncode(w, response, "register"); err != nil {
		util.Error(createProjectTag, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PatchProject(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		util.Error(getProjectTag, err)
		http.Error(w, l10n.GenericProject, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if project.Get(db.Database, id) != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	request := models.ProjectRequest{Project: &models.Project{}}
	if !jsonDecode(w, r, &request) {
		return
	}
	if err := project.Patch(db.Database, request); err != nil {
		util.Error(patchProjectTag, err)
		switch err {
		default:
			http.Error(w, l10n.GenericProject, http.StatusInternalServerError)
		}
		return
	}
	util.Info(patchProjectTag, "%s (id%d) patched", project.Name, project.ID)
	response := models.ProjectResponse{
		Project:   &project,
		StartTime: util.MarshalTime(project.StartTime),
		EndTime:   util.MarshalTime(project.EndTime),
		CreatedAt: project.CreatedAt.String(),
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error(getProjectTag, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
