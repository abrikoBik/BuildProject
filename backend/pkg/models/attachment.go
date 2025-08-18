package models

import (
	"crypto/sha256"
	"io"
	"mime/multipart"
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

type AttachmentResponse struct {
	Name string `json:"name" groups:"public"`
	Size int    `json:"size" groups:"public"`
	UUID string `json:"uuid" groups:"register"`
}

func (attachment *Attachment) Get(db *gorm.DB, id uuid.UUID) error {
	return db.Where("uuid = ?", id).Take(&attachment).Error
}

func (attachment *Attachment) Create(db *gorm.DB, file multipart.File, header *multipart.FileHeader) error {
	attachment.Name = header.Filename
	attachment.UUID = uuid.New()
	attachmentPath, err := util.MakeStoragePath(attachment.UUID.String())
	if err != nil {
		return err
	}
	attachmentFile, err := os.OpenFile(attachmentPath, os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer attachmentFile.Close()
	// should be ok?
	_, err = io.Copy(attachmentFile, file)
	if err != nil {
		return err
	}
	hash := sha256.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return err
	}
	attachment.Checksum = hash.Sum(nil)
	attachment.Size = int(header.Size)
	if (attachment.Size) > util.MaxAttachmentSize {
		return util.AttachmentSizeError
	}
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
