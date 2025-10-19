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

var (
	lGetAttachment    = util.CreateLogger("api:GetAttachment")
	lCreateAttachment = util.CreateLogger("api:CreateAttachment")
	lDeleteAttachment = util.CreateLogger("api:DeleteAttachment")
)

func GetAttachment(w http.ResponseWriter, r *http.Request) {
	attachmentUUID, err := middleware.GetUUID(middleware.AttachmentUUID, r.Context())
	lGetAttachment.Debug("attachmentUUID", attachmentUUID, "err", err)
	if err != nil {
		lGetAttachment.Error(err)
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
		lGetAttachment.Error(err)
		http.Error(w, l10n.GenericAttachment, http.StatusInternalServerError)
		return
	}
	defer f.Close()
	encoder := base64.NewEncoder(base64.StdEncoding, w)
	defer encoder.Close()
	if _, err = io.Copy(encoder, f); err != nil {
		lGetAttachment.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	lGetAttachment.Info("%s returned", attachment.UUID.String())
}

func CreateAttachment(w http.ResponseWriter, r *http.Request) {
	projectID, err := middleware.GetID(middleware.ProjectID, r.Context())
	if err != nil {
		lCreateAttachment.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var project models.Project
	if project.Get(db.Database, projectID) != nil {
		http.Error(w, l10n.ProjectNotFound, http.StatusNotFound)
		return
	}
	if r.ParseMultipartForm(consts.MaxAttachmentSize) != nil {
		lCreateAttachment.Error(err)
		http.Error(w, l10n.GenericAttachment, http.StatusInternalServerError)
		return
	}
	file, header, err := r.FormFile(consts.AttachmentMultipartKey)
	if err != nil {
		lCreateAttachment.Error(err)
		http.Error(w, l10n.GenericAttachment, http.StatusInternalServerError)
		return
	}
	defer file.Close()
	var attachment models.Attachment
	if err := attachment.Create(db.Database, projectID, file, header); err != nil {
		lCreateAttachment.Error(err)
		http.Error(w, l10n.GenericAttachment, http.StatusInternalServerError)
		return
	}
	lCreateAttachment.Info("id%d saved", project.ID)
	attachmentUUID := attachment.UUID.String()
	lCreateAttachment.Info("%s created", attachmentUUID)
	response := models.AttachmentResponse{UUID: attachmentUUID}
	if err := jsonEncode(w, response, "register"); err != nil {
		lCreateAttachment.Error(err)
		http.Error(w, l10n.GenericAttachment, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DeleteAttachment(w http.ResponseWriter, r *http.Request) {
	attachmentUUID, err := middleware.GetUUID(middleware.AttachmentUUID, r.Context())
	if err != nil {
		lDeleteAttachment.Error(err)
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
		lDeleteAttachment.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	lDeleteAttachment.Info("%s deleted", attachment.UUID.String())
}
