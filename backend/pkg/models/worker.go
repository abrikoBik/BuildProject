package models

import (
	"mp1/pkg/l10n"
	"mp1/pkg/util"

	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	Role     string     `json:"role" groups:"public,list" gorm:"not null"`
	User     User       `gorm:"foreignKey:UserID;not null"`
	UserID   uint       `json:"userID" groups:"list" gorm:"not null"`
	Projects []*Project `gorm:"many2many:project_workers"`
}

type WorkerRequest struct {
	*Worker
}

type WorkerResponse struct {
	*Worker
	Name string `json:"name" groups:"list"`
	ID   uint   `json:"id" groups:"register,public,list"`
}

func (worker *Worker) Get(db *gorm.DB, id uint, preloads ...string) error {
	return util.Preload(db, preloads...).Where("id = ?", id).Take(&worker).Error
}

func (worker *Worker) Create(db *gorm.DB, request WorkerRequest) error {
	if len(request.Role) == 0 {
		return l10n.InvalidRoleError
	}
	worker.Role = request.Role
	worker.UserID = request.UserID
	return db.Create(&worker).Error
}

func (worker *Worker) Patch(db *gorm.DB, request WorkerRequest) error {
	if len(request.Role) > 0 {
		worker.Role = request.Role
	}
	return db.Save(&worker).Error
}
