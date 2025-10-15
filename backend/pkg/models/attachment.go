package models

import (
	"crypto/sha256"
	"io"
	"mime/multipart"
	"mp1/pkg/consts"
	"mp1/pkg/l10n"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attachment struct {
	gorm.Model
	Name string `json:"name" groups:"public,list" gorm:"not null"`
	// Size in bytes.
	Size      int64     `json:"size" groups:"public" gorm:"not null"`
	UUID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();not null"`
	ProjectID uint      `gorm:"not null"`
	Checksum  []byte    `gorm:"not null"`
}

type AttachmentResponse struct {
	*Attachment
	UUID string `json:"uuid" groups:"register,list"`
}

func (attachment *Attachment) Get(db *gorm.DB, id uuid.UUID) error {
	return db.Where("uuid = ?", id).Take(&attachment).Error
}

func (attachment *Attachment) Create(db *gorm.DB, projectID uint, multipartFile multipart.File, header *multipart.FileHeader) error {
	var project Project
	if err := project.Get(db, projectID); err != nil {
		return l10n.ProjectNotFoundError
	}
	attachment.ProjectID = projectID
	attachment.Size = header.Size
	if (attachment.Size) > consts.MaxAttachmentSize {
		return l10n.AttachmentSizeError
	}
	attachment.Name = header.Filename
	attachment.UUID = uuid.New()
	path := filepath.Join(consts.AttachmentPath, attachment.UUID.String())
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	hash := sha256.New()
	w := io.MultiWriter(file, hash)
	if _, err = io.Copy(w, multipartFile); err != nil {
		return err
	}
	attachment.Checksum = hash.Sum(nil)
	return db.Create(&attachment).Error
}

func (attachment *Attachment) Delete(db *gorm.DB) error {
	attachmentPath := filepath.Join(consts.AttachmentPath, attachment.UUID.String())
	if err := os.Remove(attachmentPath); err != nil {
		return err
	}
	return db.Delete(&attachment).Error
}
