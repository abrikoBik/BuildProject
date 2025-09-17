package handlers

import (
	"encoding/base64"
	"errors"
	"io"
	"mp1/pkg/api/middleware"
	"mp1/pkg/consts"
	"mp1/pkg/db"
	"mp1/pkg/l10n"
	"mp1/pkg/models"
	"mp1/pkg/util"
	"net/http"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

const (
	getAttachmentTag    = "api:GetAttachment"
	createAttachmentTag = "api:CreateAttachment"
	deleteAttachmentTag = "api:DeleteAttachment"
)

func GetAttachment(w http.ResponseWriter, r *http.Request) {
	attachmentUUID, err := middleware.GetUUID(middleware.AttachmentUUID, r.Context())
	util.Debug(getAttachmentTag, "attachmentUUID", attachmentUUID, "err", err)
	if err != nil {
		util.Error(getAttachmentTag, err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var attachment models.Attachment
	result := attachment.Get(db.Database, attachmentUUID)
	if errors.Is(result, gorm.ErrRecordNotFound) {
		http.Error(w, l10n.AttachmentNotFound, http.StatusNotFound)
		return
	}

	attachmentPath := filepath.Join(consts.AttachmentPath, attachment.UUID.String())
	f, err := os.Open(attachmentPath)
	if err != nil {
		util.Error(getAttachmentTag, err)
		http.Error(w, l10n.GenericAttachment, http.StatusInternalServerError)
		return
	}
	defer f.Close()
	encoder := base64.NewEncoder(base64.StdEncoding, w)
	if _, err = io.Copy(encoder, f); err != nil {
		util.Error(getAttachmentTag, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	util.Info(getAttachmentTag, "%s returned", attachment.UUID.String())
}

func CreateAttachment(w http.ResponseWriter, r *http.Request) {
	projectID, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		util.Error(createAttachmentTag, err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if project.Get(db.Database, projectID) != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}

	util.Check(r.ParseMultipartForm(consts.MaxAttachmentSize))
	file, header, err := r.FormFile(consts.AttachmentMultipartKey)
	if err != nil {
		util.Error(createAttachmentTag, err)
		http.Error(w, l10n.GenericAttachment, http.StatusInternalServerError)
		return
	}
	defer util.Check(file.Close())
	var attachment models.Attachment
	if err := attachment.Create(db.Database, projectID, file, header); err != nil {
		util.Error(createAttachmentTag, err)
		http.Error(w, l10n.GenericAttachment, http.StatusInternalServerError)
		return
	}

	project.Attachments = append(project.Attachments, attachment)
	if err := db.Database.Save(&project).Error; err != nil {
		util.Error(createAttachmentTag, err)
		http.Error(w, l10n.GenericAttachment, http.StatusInternalServerError)
		return
	}
	util.Info(createAttachmentTag, "id%d saved", project.ID)

	attachmentUUID := attachment.UUID.String()
	util.Info(createAttachmentTag, "%s created", attachmentUUID)
	response := models.AttachmentResponse{UUID: attachmentUUID}
	if err := jsonEncode(w, response, "register"); err != nil {
		util.Error(createAttachmentTag, err)
		http.Error(w, l10n.GenericAttachment, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DeleteAttachment(w http.ResponseWriter, r *http.Request) {
	attachmentUUID, err := middleware.GetUUID(middleware.AttachmentUUID, r.Context())
	if err != nil {
		util.Error(deleteAttachmentTag, err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var attachment models.Attachment
	result := attachment.Get(db.Database, attachmentUUID)
	if errors.Is(result, gorm.ErrRecordNotFound) {
		http.Error(w, l10n.AttachmentNotFound, http.StatusNotFound)
		return
	}
	if err := attachment.Delete(db.Database); err != nil {
		util.Error(deleteAttachmentTag, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	util.Info(deleteAttachmentTag, "%s deleted", attachment.UUID.String())
}
