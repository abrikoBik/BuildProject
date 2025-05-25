package models

import (
	"mp1/pkg/util"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string  `json:"name" groups:"public" gorm:"not null"`
	Description string  `json:"description" groups:"public"`
	UserID      uint    `json:"user" groups:"public"`
	Stages      []Stage `json:"stages" groups:"public" gorm:"foreignKey:ProjectID"`
}

type ProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"userID"`
}

type ProjectResponse struct {
	*Project
	CreatedAt string `json:"createdAt" groups:"public"`
	ID        uint   `json:"id" groups:"register"`
}

func (project *Project) Get(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).Take(&project).Error
}

func (project *Project) Create(db *gorm.DB, request ProjectRequest) error {
	var user User
	if err := user.Get(db, request.UserID); err != nil {
		return util.UserNotFoundError
	}
	project.UserID = user.ID
	project.Name = request.Name
	project.Description = request.Description
	project.Stages = make([]Stage, 0)
	return db.Create(&project).Error
}
