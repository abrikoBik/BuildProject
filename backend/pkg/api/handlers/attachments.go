package handlers

import (
	"errors"
	"mp1/pkg/api/middleware"
	"mp1/pkg/db"
	"mp1/pkg/l10n"
	"mp1/pkg/models"
	"mp1/pkg/util"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	getAttachmentTag    = "api:GetAttachment"
	createAttachmentTag = "api:CreateAttachment"
	deleteAttachmentTag = "api:DeleteAttachment"
)

func GetAttachment(w http.ResponseWriter, r *http.Request) {
	// jank
	contextUUID := r.Context().Value(middleware.AttachmentUUID).(uuid.UUID)
	var attachment models.Attachment
	result := attachment.Get(db.Database, contextUUID)
	if errors.Is(result, gorm.ErrRecordNotFound) {
		http.Error(w, l10n.AttachmentNotFound, http.StatusNotFound)
		return
	}
	response := models.AttachmentResponse{
		Name: attachment.Name,
		Size: attachment.Size,
	}
	if err := jsonEncode(w, response, "public"); err != nil {
		util.Error(getAttachmentTag, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	util.Info(getAttachmentTag, "%s returned", attachment.UUID.String())
}

func CreateAttachment(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(util.MaxAttachmentSize)
	file, header, err := r.FormFile(util.FormKey)
	if err != nil {
		util.Error(createAttachmentTag, err)
		return
	}
	defer file.Close()
	var attachment models.Attachment
	if err := attachment.Create(db.Database, file, header); err != nil {
		util.Error(createAttachmentTag, err)
		return
	}
	attachmentUUID := attachment.UUID.String()
	util.Info(createAttachmentTag, "%s created", attachmentUUID)
	response := models.AttachmentResponse{UUID: attachmentUUID}
	if err := jsonEncode(w, response, "register"); err != nil {
		util.Error(createAttachmentTag, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// TODO
func DeleteAttachment(w http.ResponseWriter, r *http.Request) {

}
