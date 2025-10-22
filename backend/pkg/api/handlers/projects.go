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
	lGetProject    = util.CreateLogger("api:GetProject")
	lCreateProject = util.CreateLogger("api:CreateProject")
	lPatchProject  = util.CreateLogger("api:PatchProject")
	lDeleteProject = util.CreateLogger("api:DeleteProject")

	lListAttachments    = util.CreateLogger("api:ListAttachments")
	lListEstimateGroups = util.CreateLogger("api:ListEstimateGroups")
	lListStageGroups    = util.CreateLogger("api:ListStageGroups")
	lListWorkers        = util.CreateLogger("api:ListWorkers")

	lAddWorker    = util.CreateLogger("api:AddWorker")
	lRemoveWorker = util.CreateLogger("api:RemoveWorker")
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
		StartTime: project.StartTime.Time.Unix(),
		EndTime:   project.EndTime.Time.Unix(),
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
	var project models.Project
	if err := project.Create(db.Database, request); err != nil {
		lCreateProject.Error(err)
		switch err {
		default:
			http.Error(w, l10n.GenericProject, http.StatusInternalServerError)
		}
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
		StartTime: project.StartTime.Time.Unix(),
		EndTime:   project.EndTime.Time.Unix(),
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

func ListAttachments(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		lListAttachments.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if project.Get(db.Database, id, "Attachments") != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	sort.Slice(project.Attachments, func(i, j int) bool {
		return project.Attachments[i].CreatedAt.UnixMilli() > project.Attachments[j].CreatedAt.UnixMilli()
	})
	var response []models.AttachmentResponse
	for _, a := range project.Attachments {
		toAppend := models.AttachmentResponse{
			Attachment: &a,
			UUID:       a.UUID.String(),
		}
		response = append(response, toAppend)
	}
	if err := jsonEncode(w, response, "list"); err != nil {
		lListAttachments.Error(err)
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
		return project.EstimateGroups[i].CreatedAt.UnixMilli() > project.EstimateGroups[j].CreatedAt.UnixMilli()
	})
	var response []models.EstimateGroupResponse
	for _, g := range project.EstimateGroups {
		toAppend := models.EstimateGroupResponse{
			EstimateGroup: &g,
			ID:            g.ID,
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
	lListStageGroups.Debug("project.StageGroups", project.StageGroups)
	sort.Slice(project.StageGroups, func(i, j int) bool {
		return project.StageGroups[i].CreatedAt.UnixMilli() > project.StageGroups[j].CreatedAt.UnixMilli()
	})
	var response []models.StageGroupResponse
	for _, g := range project.StageGroups {
		toAppend := models.StageGroupResponse{
			StageGroup: &g,
			ID:         g.ID,
		}
		response = append(response, toAppend)
	}
	if err := jsonEncode(w, response, "list"); err != nil {
		lListStageGroups.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ListWorkers(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		lListWorkers.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if project.Get(db.Database, id, "Workers.User") != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	lListWorkers.Debug("project.Workers", project.Workers)
	sort.Slice(project.Workers, func(i, j int) bool {
		return project.Workers[i].CreatedAt.UnixMilli() > project.Workers[j].CreatedAt.UnixMilli()
	})
	var response []models.WorkerResponse
	for _, w := range project.Workers {
		toAppend := models.WorkerResponse{
			Worker: w,
			Name:   w.User.Name,
			ID:     w.ID,
		}
		response = append(response, toAppend)
	}
	if err := jsonEncode(w, response, "list"); err != nil {
		lListWorkers.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func AddWorker(w http.ResponseWriter, r *http.Request) {
	projectID, err1 := middleware.GetID(middleware.ProjectID, r.Context())
	workerID, err2 := middleware.GetID(middleware.WorkerID, r.Context())
	if err := errors.Join(err1, err2); err != nil {
		lAddWorker.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if project.Get(db.Database, projectID, "Workers.User.Projects") != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	var worker models.Worker
	if worker.Get(db.Database, workerID, "User.Projects") != nil {
		http.Error(w, l10n.WorkerNotFound, http.StatusNotFound)
		return
	}
	project.Workers = append(project.Workers, &worker)
	worker.User.Projects = append(worker.User.Projects, project)
	if err := errors.Join(db.Database.Save(project).Error, db.Database.Save(&worker).Error); err != nil {
		lAddWorker.Error(err)
		http.Error(w, l10n.GenericWorker, http.StatusInternalServerError)
	}
	lAddWorker.Info("added id%d to %d", worker.ID, project.ID)
	w.WriteHeader(http.StatusOK)
}

func RemoveWorker(w http.ResponseWriter, r *http.Request) {
	projectID, err1 := middleware.GetID(middleware.ProjectID, r.Context())
	workerID, err2 := middleware.GetID(middleware.WorkerID, r.Context())
	if err := errors.Join(err1, err2); err != nil {
		lRemoveWorker.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if project.Get(db.Database, projectID, "Workers.User") != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	var worker models.Worker
	if worker.Get(db.Database, workerID, "User.Projects") != nil {
		http.Error(w, l10n.WorkerNotFound, http.StatusNotFound)
		return
	}
	lRemoveWorker.Debug("project.Workers", project.Workers)
	if err := errors.Join(
		db.Database.Model(&project).Association("Workers").Delete(&worker),
		db.Database.Model(&worker.User).Association("Projects").Delete(&project),
	); err != nil {
		lRemoveWorker.Error(err)
		http.Error(w, l10n.GenericWorker, http.StatusInternalServerError)
		return
	}
	lRemoveWorker.Debug("project.Workers", project.Workers)
	lRemoveWorker.Info("removed id%d from %d", worker.ID, project.ID)
	w.WriteHeader(http.StatusOK)
}
