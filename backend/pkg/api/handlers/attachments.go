package handlers

import (
	"mp1/pkg/db"
	"mp1/pkg/models"
	"mp1/pkg/util"
	"net/http"
	"path"
)

const (
	getAttachmentTag    = "api:GetAttachment"
	createAttachmentTag = "api:CreateAttachment"
	uploadAttachmentTag = "api:UploadAttachment"
	deleteAttachmentTag = "api:DeleteAttachment"
)

// TODO
func GetAttachment(w http.ResponseWriter, r *http.Request) {

}

func CreateAttachment(w http.ResponseWriter, r *http.Request) {
	var request models.AttachmentRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	var attachment models.Attachment
	if err := attachment.Create(db.Database, request); err != nil {
		util.Error(createAttachmentTag, err)
		return
	}
	uuid := attachment.UUID.String()
	util.Info(createAttachmentTag, "%s created", uuid)
	response := models.AttachmentResponse{UUID: uuid}
	if err := jsonEncode(w, response, "register"); err != nil {
		util.Error(createAttachmentTag, err)
		return
	}
	// not sure if that's the correct way to do it
	http.Redirect(w, r, path.Join("/attachments", uuid), http.StatusPermanentRedirect)
}

// TODO
func UploadAttachment(w http.ResponseWriter, r *http.Request) {

}

// TODO
func DeleteAttachment(w http.ResponseWriter, r *http.Request) {

}
