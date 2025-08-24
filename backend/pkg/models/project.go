package models

import (
	"mp1/pkg/l10n"
	"mp1/pkg/util"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string      `gorm:"not null"`
	StartTime   pq.NullTime `gorm:"not null"`
	EndTime     pq.NullTime `gorm:"not null"`
	UserID      uint        `gorm:"not null"`
	Description string
	Groups      []Group
	Attachments []Attachment
}

type ProjectRequest struct {
	Name        string `json:"name"`
	StartTime   int64  `json:"startTime"`
	EndTime     int64  `json:"endTime"`
	UserID      uint   `json:"userID"`
	Description string `json:"description"`
}

type ProjectResponse struct {
	Name        string `json:"name" groups:"public"`
	StartTime   pq.NullTime
	EndTime     pq.NullTime
	UserID      uint   `json:"user" groups:"public"`
	Description string `json:"description" groups:"public"`
	CreatedAt   string `json:"createdAt" groups:"public"`
	ID          uint   `json:"id" groups:"register"`
}

func (project *Project) Get(db *gorm.DB, id uint) error {
	return db.Preload("Groups").Where("id = ?", id).Take(&project).Error
}

func (project *Project) Create(db *gorm.DB, request ProjectRequest) error {
	var user User
	if err := user.Get(db, request.UserID); err != nil {
		return l10n.UserNotFoundError
	}
	project.UserID = user.ID
	project.Name = request.Name
	project.Description = request.Description
	project.Groups = make([]Group, 0)
	return db.Create(&project).Error
}

func (project *Project) Delete(db *gorm.DB) error {
	for _, group := range project.Groups {
		if err := group.Delete(db); err != nil {
			return err
		}
	}
	return db.Delete(&project).Error
}

func (project *Project) Patch(db *gorm.DB, request ProjectRequest) error {
	if len(request.Name) > 0 {
		project.Name = request.Name
	}
	if len(request.Description) > 0 {
		project.Description = request.Description
	}
	if request.StartTime != 0 {
		project.StartTime = util.UnmarshalTime(request.StartTime)
	}
	if request.EndTime != 0 {
		project.EndTime = util.UnmarshalTime(request.EndTime)
	}
	return db.Save(&project).Error
}

func (project *Project) AddGroup(db *gorm.DB, group Group) error {
	project.Groups = append(project.Groups, group)
	return db.Save(&project).Error
}
