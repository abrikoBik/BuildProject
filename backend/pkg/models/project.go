package models

import (
	"mp1/pkg/l10n"
	"mp1/pkg/util"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name           string      `json:"name" groups:"public,list" gorm:"not null"`
	Description    string      `json:"description" groups:"public,list"`
	Latitude       float64     `json:"latitude" groups:"public" gorm:"not null"`
	Longitude      float64     `json:"longitude" groups:"public" gorm:"not null"`
	StartTime      pq.NullTime `gorm:"not null"`
	EndTime        pq.NullTime `gorm:"not null"`
	UserID         uint        `json:"userID" groups:"public" gorm:"not null"`
	Workers        []*Worker   `gorm:"many2many:project_workers"`
	Attachments    []Attachment
	StageGroups    []StageGroup
	EstimateGroups []EstimateGroup
}

type ProjectRequest struct {
	*Project
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
}

type ProjectResponse struct {
	*Project
	StartTime int64 `json:"startTime" groups:"public,list"`
	EndTime   int64 `json:"endTime" groups:"public,list"`
	ID        uint  `json:"id" groups:"register,list"`
}

func (project *Project) Get(db *gorm.DB, id uint, preloads ...string) error {
	return util.Preload(db, preloads...).Where("id = ?", id).Take(&project).Error
}

func (project *Project) Create(db *gorm.DB, request ProjectRequest) error {
	var user User
	if err := user.Get(db, request.UserID, "Projects"); err != nil {
		return l10n.UserNotFoundError
	}
	project.UserID = user.ID
	project.Name = request.Name
	project.Description = request.Description
	project.Latitude = request.Latitude
	project.Longitude = request.Longitude
	project.StartTime = util.UnmarshalTime(request.StartTime)
	project.EndTime = util.UnmarshalTime(request.EndTime)
	project.StageGroups = make([]StageGroup, 0)
	project.Attachments = make([]Attachment, 0)
	project.EstimateGroups = make([]EstimateGroup, 0)
	return db.Create(&project).Error
}

func (project *Project) Delete(db *gorm.DB) error {
	for _, group := range project.StageGroups {
		if err := group.Delete(db); err != nil {
			return err
		}
	}
	for _, attachment := range project.Attachments {
		if err := attachment.Delete(db); err != nil {
			return err
		}
	}
	for _, sg := range project.StageGroups {
		if err := sg.Delete(db); err != nil {
			return err
		}
	}
	for _, eg := range project.EstimateGroups {
		if err := eg.Delete(db); err != nil {
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
	if request.Latitude != 0 {
		project.Latitude = request.Latitude
	}
	if request.Longitude != 0 {
		project.Longitude = request.Longitude
	}
	if request.StartTime != 0 {
		project.StartTime = util.UnmarshalTime(request.StartTime)
	}
	if request.EndTime != 0 {
		project.EndTime = util.UnmarshalTime(request.EndTime)
	}
	return db.Save(&project).Error
}
