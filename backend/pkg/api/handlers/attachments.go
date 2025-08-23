package handlers

import (
	"errors"
	"mime/multipart"
	"mp1/pkg/api/middleware"
	"mp1/pkg/db"
	"mp1/pkg/l10n"
	"mp1/pkg/models"
	"mp1/pkg/util"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	getAttachmentTag    = "api:GetAttachment"
	createAttachmentTag = "api:CreateAttachment"
	deleteAttachmentTag = "api:DeleteAttachment"
)

func GetAttachment(w http.ResponseWriter, r *http.Request) {
	contextUUID := r.Context().Value(middleware.AttachmentUUID).(uuid.UUID)
	util.Debug("api.GetAttachment", "contextUUID", contextUUID)
	var attachment models.Attachment
	result := attachment.Get(db.Database, contextUUID)
	if errors.Is(result, gorm.ErrRecordNotFound) {
		http.Error(w, l10n.AttachmentNotFound, http.StatusNotFound)
		return
	}

	attachmentPath := filepath.Join(util.StoragePath, attachment.UUID.String())
	contents, err := os.ReadFile(attachmentPath)
	if err != nil {
		util.Error(getAttachmentTag, err)
		http.Error(w, l10n.AttachmentError, http.StatusInternalServerError)
		return
	}
	mw := multipart.NewWriter(w)
	fw, err := mw.CreateFormFile(util.FormKey, attachment.Name)
	if err != nil {
		util.Error(getAttachmentTag, err)
		http.Error(w, l10n.AttachmentError, http.StatusInternalServerError)
		return
	}
	if _, err := fw.Write(contents); err != nil {
		util.Error(getAttachmentTag, err)
		http.Error(w, l10n.AttachmentError, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	util.Info(getAttachmentTag, "%s returned", attachment.UUID.String())
}

func CreateAttachment(w http.ResponseWriter, r *http.Request) {
	util.Check(r.ParseMultipartForm(util.MaxAttachmentSize))
	file, header, err := r.FormFile(util.FormKey)
	if err != nil {
		util.Error(createAttachmentTag, err)
		http.Error(w, l10n.AttachmentError, http.StatusInternalServerError)
		return
	}
	defer util.Check(file.Close())
	var attachment models.Attachment
	if err := attachment.Create(db.Database, file, header); err != nil {
		util.Error(createAttachmentTag, err)
		http.Error(w, l10n.AttachmentError, http.StatusInternalServerError)
		return
	}
	attachmentUUID := attachment.UUID.String()
	util.Info(createAttachmentTag, "%s created", attachmentUUID)
	response := models.AttachmentResponse{UUID: attachmentUUID}
	if err := jsonEncode(w, response, "register"); err != nil {
		util.Error(createAttachmentTag, err)
		http.Error(w, l10n.AttachmentError, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DeleteAttachment(w http.ResponseWriter, r *http.Request) {
	// jank
	contextUUID := r.Context().Value(middleware.AttachmentUUID).(uuid.UUID)
	var attachment models.Attachment
	result := attachment.Get(db.Database, contextUUID)
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
