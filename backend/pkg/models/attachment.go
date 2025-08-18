package models

import (
	"mp1/pkg/util"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attachment struct {
	gorm.Model
	Name string `gorm:"not null"`
	// Size in bytes.
	Size     int       `gorm:"not null"`
	UUID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Checksum []byte    `gorm:"not null"`
}

type AttachmentRequest struct {
	Name string `json:"name"`
	// Should be provided via a follow-up request.
	contents []byte
}

type AttachmentResponse struct {
	Name string `json:"name" groups:"public"`
	Size int    `json:"size" groups:"public"`
	UUID string `json:"uuid" groups:"register"`
}

func (attachment *Attachment) Get(db *gorm.DB, id uuid.UUID) error {
	return db.Where("uuid = ?", id).Take(&attachment).Error
}

func (attachment *Attachment) File(db *gorm.DB, contents []byte) error {
	attachmentPath, err := util.MakeStoragePath(attachment.UUID.String())
	if err != nil {
		return err
	}
	if err := os.WriteFile(attachmentPath, contents, 0644); err != nil {
		return err
	}
	sum, err := util.GenerateChecksum(contents)
	if err != nil {
		return err
	}
	attachment.Checksum = sum
	attachment.Size = len(contents)
	if (attachment.Size) > util.MaxAttachmentSize {
		return util.AttachmentSizeError
	}
	return db.Save(&attachment).Error
}

func (attachment *Attachment) Create(db *gorm.DB, request AttachmentRequest) error {
	attachment.Name = request.Name
	attachment.UUID = uuid.New()
	return db.Create(&attachment).Error
}

func (attachment *Attachment) Delete(db *gorm.DB) error {
	// might get very messy
	attachmentPath, err := util.MakeStoragePath(attachment.UUID.String())
	if err != nil {
		return err
	}
	if err := os.Remove(attachmentPath); err != nil {
		return err
	}
	return db.Delete(&attachment).Error
}
