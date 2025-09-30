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
	lGetProject             = util.CreateLogger("api:GetProject")
	lCreateProject          = util.CreateLogger("api:CreateProject")
	lPatchProject           = util.CreateLogger("api:PatchProject")
	lDeleteProject          = util.CreateLogger("api:DeleteProject")
	lListProjectAttachments = util.CreateLogger("api:ListProjectAttachments")
	lListEstimateGroups     = util.CreateLogger("api:ListEstimateGroups")
	lListStageGroups        = util.CreateLogger("api:ListStageGroups")
)

func GetProject(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		lGetProject.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if errors.Is(project.Get(db.Database, id), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	response := models.ProjectResponse{
		Project:   &project,
		StartTime: util.MarshalTime(project.StartTime),
		EndTime:   util.MarshalTime(project.EndTime),
		CreatedAt: project.CreatedAt.String(),
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		lGetProject.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	lGetProject.Info("id%d returned", project.ID)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var request models.ProjectRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	var user models.User
	if user.Get(db.Database, request.UserID, "Projects") != nil {
		http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		return
	}
	var project models.Project
	if err := project.Create(db.Database, request); err != nil {
		lCreateProject.Error(err)
		switch err {
		default:
			http.Error(w, l10n.GenericProject, http.StatusInternalServerError)
		}
	}
	// TODO: user.AddProject
	user.Projects = append(user.Projects, project)
	if err := db.Database.Save(user).Error; err != nil {
		lCreateProject.Error(err)
		http.Error(w, l10n.GenericUser, http.StatusInternalServerError)
	}
	lCreateProject.Info("%s (id%d) created", project.Name, project.ID)
	response := models.ProjectResponse{ID: project.ID}
	if err := jsonEncode(w, response, "register"); err != nil {
		lCreateProject.Error(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PatchProject(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		lPatchProject.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if project.Get(db.Database, id) != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	var request models.ProjectRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	if err := project.Patch(db.Database, request); err != nil {
		lPatchProject.Error(err)
		switch err {
		default:
			http.Error(w, l10n.GenericProject, http.StatusInternalServerError)
		}
		return
	}
	lPatchProject.Info("%s (id%d) patched", project.Name, project.ID)
	response := models.ProjectResponse{
		Project:   &project,
		StartTime: util.MarshalTime(project.StartTime),
		EndTime:   util.MarshalTime(project.EndTime),
		CreatedAt: project.CreatedAt.String(),
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		lPatchProject.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	projectID, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		lDeleteProject.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if errors.Is(project.Get(db.Database, projectID), gorm.ErrRecordNotFound) {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
	}
	if err := project.Delete(db.Database); err != nil {
		lDeleteEstimate.Error(err)
		return
	}
	lDeleteProject.Info("id%d deleted", project.ID)
	w.WriteHeader(http.StatusOK)
}

func ListProjectAttachments(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		lListProjectAttachments.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if project.Get(db.Database, id, "Attachments") != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	sort.Slice(project.Attachments, func(i, j int) bool {
		return project.Attachments[i].UpdatedAt.UnixMilli() > project.Attachments[j].UpdatedAt.UnixMilli()
	})
	var response []models.AttachmentResponse
	for _, a := range project.Attachments {
		toAppend := models.AttachmentResponse{
			Attachment: &models.Attachment{
				Name: a.Name,
				Size: a.Size,
			},
			UUID: a.UUID.String(),
		}
		response = append(response, toAppend)
	}
	if err := jsonEncode(w, response, "list"); err != nil {
		lListProjectAttachments.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ListEstimateGroups(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		lListEstimateGroups.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if project.Get(db.Database, id, "EstimateGroups") != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	sort.Slice(project.EstimateGroups, func(i, j int) bool {
		return project.EstimateGroups[i].UpdatedAt.UnixMilli() > project.EstimateGroups[j].UpdatedAt.UnixMilli()
	})
	var response []models.EstimateGroupResponse
	for _, g := range project.EstimateGroups {
		toAppend := models.EstimateGroupResponse{
			EstimateGroup: &g,
		}
		response = append(response, toAppend)
	}
	if err := jsonEncode(w, response, "list"); err != nil {
		lListEstimateGroups.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ListStageGroups(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		lListStageGroups.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if project.Get(db.Database, id, "StageGroups") != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	sort.Slice(project.StageGroups, func(i, j int) bool {
		return project.StageGroups[i].UpdatedAt.UnixMilli() > project.StageGroups[j].UpdatedAt.UnixMilli()
	})
	var response []models.StageGroupResponse
	for _, g := range project.StageGroups {
		toAppend := models.StageGroupResponse{
			StageGroup: &g,
		}
		response = append(response, toAppend)
	}
	if err := jsonEncode(w, response, "list"); err != nil {
		lListStageGroups.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
