package models

import (
	"errors"
	"mp1/pkg/l10n"
	"mp1/pkg/util"

	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	Role   string  `json:"role" groups:"public,list" gorm:"not null"`
	User   User    `gorm:"foreignKey:UserID;not null"`
	UserID uint    `json:"userID" groups:"list" gorm:"not null"`
	Stages []Stage `gorm:"many2many:stage_workers"`
}

type WorkerRequest struct {
	*Worker
}

type WorkerResponse struct {
	*Worker
	ID uint `json:"id" groups:"register,list"`
}

func (worker *Worker) Get(db *gorm.DB, id uint, preloads ...string) error {
	return util.Preload(db, preloads...).Where("id = ?", id).Take(&worker).Error
}

func (worker *Worker) Create(db *gorm.DB, request WorkerRequest) error {
	if len(request.Role) == 0 {
		return l10n.InvalidRoleError
	}
	var user User
	if err := user.Get(db, request.UserID, "Worker"); err != nil {
		return err
	}
	if user.Worker != nil {
		return l10n.WorkerExistsError
	}
	worker.Role = request.Role
	worker.User = user
	worker.UserID = user.ID
	return errors.Join(db.Create(&worker).Error, db.Save(&user).Error)
}

func (worker *Worker) Patch(db *gorm.DB, request WorkerRequest) error {
	if len(request.Role) > 0 {
		worker.Role = request.Role
	}
	return db.Save(&worker).Error
}
